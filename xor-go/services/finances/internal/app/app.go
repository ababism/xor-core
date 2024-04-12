package app

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel/sdk/trace"
	"go.uber.org/zap"
	"time"
	"xor-go/pkg/metrics"
	"xor-go/pkg/xdb/postgres"
	"xor-go/pkg/xerror"
	"xor-go/pkg/xshutdown"
	"xor-go/pkg/xtracer"
	"xor-go/services/finances/internal/config"
	"xor-go/services/finances/internal/daemon/banker"
	"xor-go/services/finances/internal/handler/http"
	"xor-go/services/finances/internal/log"
	"xor-go/services/finances/internal/repository/payments"
	"xor-go/services/finances/internal/repository/postgre"
	"xor-go/services/finances/internal/service"
)

type App struct {
	cfg            *config.Config
	handler        http.Handler
	address        string
	tracerProvider *trace.TracerProvider
}

func NewApp(cfg *config.Config) (*App, error) {

	// INFRASTRUCTURE ----------------------------------------------------------------------

	err := log.Init(cfg.Logger, cfg.App)
	if err != nil {
		return nil, errors.Wrap(err, "Init Logger")
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

	// Инициализируем Prometheus
	metrics.InitOnce(cfg.Metrics, log.Logger, metrics.AppInfo{
		Name:        cfg.App.Name,
		Environment: string(cfg.App.Environment),
		Version:     cfg.App.Version,
	})
	log.Logger.Info("Init Metrics – success")

	// REPOSITORY ----------------------------------------------------------------------

	// Инициализация PostgreSQL
	postgresDb, err := postgres.NewDB(cfg.Postgres)
	if err != nil {
		log.Logger.Fatal("Error init Postgres DB:", zap.Error(err))
		return nil, errors.Wrap(err, "Init Postgres DB")
	}

	// Инициализация всех репозиториев
	bankAccountRepo := postgre.NewBankAccountRepository(postgresDb)
	discountRepo := postgre.NewDiscountRepository(postgresDb)
	paymentRepo := postgre.NewPaymentRepository(postgresDb)
	productRepo := postgre.NewProductRepository(postgresDb)
	payoutRequest := postgre.NewPayoutRequestRepository(postgresDb)
	purchaseRequest := postgre.NewPurchaseRequestRepository(postgresDb)

	// Инициализация всех клиентов
	paymentsClientServer, err := payments.NewClientWithResponses(cfg.PaymentsClient.Uri)
	if err != nil {
		log.Logger.Fatal("cannot initialize generated Payments Client:", zap.Error(err))
		return nil, err
	}
	paymentsClient := payments.NewPaymentsClient(paymentsClientServer)

	// SERVICE LAYER ----------------------------------------------------------------------

	// Name layer
	bankAccountService := service.NewBankAccountService(bankAccountRepo)
	discountService := service.NewDiscountService(discountRepo)
	productService := service.NewProductService(productRepo)
	paymentService := service.NewPaymentService(paymentRepo, paymentsClient)
	payoutRequestService := service.NewPayoutRequestService(
		payoutRequest,
		paymentRepo,
		paymentsClient,
		bankAccountService,
	)
	purchaseRequestService := service.NewPurchaseRequestService(
		purchaseRequest,
		paymentsClient,
		productService,
		paymentRepo,
		discountRepo,
	)

	log.Logger.Info(fmt.Sprintf("Init %s – success", cfg.App.Name))

	// TRANSPORT LAYER ----------------------------------------------------------------------

	mainHandler := http.NewHandler(
		cfg,
		bankAccountService,
		discountService,
		paymentService,
		productService,
		purchaseRequestService,
		payoutRequestService,
	)

	// инициализируем адрес сервера
	address := fmt.Sprintf(":%s", cfg.Http.Port)

	// DAEMONS ----------------------------------------------------------------------

	bankerDaemon := banker.NewBanker(
		payoutRequestService,
		purchaseRequestService,
		productService,
		bankAccountService,
		paymentsClient,
	)

	bankerDaemon.Start(10 * time.Second)

	return &App{
		cfg:            cfg,
		handler:        mainHandler,
		address:        address,
		tracerProvider: tp,
	}, nil
}
