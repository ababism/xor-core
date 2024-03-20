package app

import (
	"context"
	"fmt"
	"go.opentelemetry.io/otel/sdk/trace"
	"go.uber.org/zap"
	"xor-go/pkg/xdb/postgres"
	"xor-go/pkg/xerror"
	"xor-go/pkg/xshutdown"
	"xor-go/pkg/xtracer"
	"xor-go/services/finances/internal/config"
	"xor-go/services/finances/internal/log"
	"xor-go/services/finances/internal/service"
	"xor-go/services/finances/internal/service/adapters"
)

type App struct {
	cfg            *config.Config
	address        string
	tracerProvider *trace.TracerProvider
	service        adapters.DriverService
}

func NewApp(cfg *config.Config) (*App, error) {

	// INFRASTRUCTURE ----------------------------------------------------------------------

	err := log.Init(cfg)
	if err != nil {
		return nil, err
	}
	// Чистим кэш logger при shutdown
	xshutdown.AddCallback(
		&xshutdown.Callback{
			Name: "ZapLoggerCacheWipe",
			FnCtx: func(ctx context.Context) error {
				return log.Logger.Sync()
			},
		})
	log.Logger.Info("Init Logger – success")

	// Инициализируем обработку ошибок
	err = xerror.InitAppError(cfg.App)
	if err != nil {
		log.Logger.Fatal("while initializing App Error handling package", zap.Error(err))
	}

	// Инициализируем трассировку
	tp, err := xtracer.Init(cfg.Tracer, cfg.App)
	if err != nil {
		return nil, err
	}
	xshutdown.AddCallback(
		&xshutdown.Callback{
			Name: "OpenTelemetryShutdown",
			FnCtx: func(ctx context.Context) error {
				if err := tp.Shutdown(context.Background()); err != nil {
					log.Logger.Error("Error shutting down tracer provider: %v", zap.Error(err))
					return err
				}
				return nil
			},
		})
	log.Logger.Info("Init Tracer – success")

	// TODO Инициализируем Prometheus
	//httpResponser := httpresponse.NewHttpResponseWrapper(logger)

	//metrics.InitOnce(cfg.Metrics, logger, metrics.AppInfo{
	//	Name:        cfg.App.Name,
	//	Environment: cfg.App.Environment,
	//	Version:     cfg.App.Version,
	//})
	log.Logger.Info("Init Metrics – success")

	// REPOSITORY ----------------------------------------------------------------------

	postgresDb, err := postgres.NewDB(cfg.PostgresConfig)
	if err != nil {
		return nil, err
	}

	// SERVICE LAYER ----------------------------------------------------------------------

	// Service layer
	driverService := service.NewDriverService(driverRepo)

	log.Logger.Info(fmt.Sprintf("Init %s – success", cfg.App.Service))

	// TRANSPORT LAYER ----------------------------------------------------------------------

	// инициализируем адрес сервера
	address := fmt.Sprintf(":%d", cfg.Http.Port)

	return &App{
		cfg:            cfg,
		service:        driverService,
		address:        address,
		tracerProvider: tp,
	}, nil
}
