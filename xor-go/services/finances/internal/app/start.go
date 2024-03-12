package app

import (
	"context"
	"fmt"
	ginzap "github.com/gin-contrib/zap"
	requestid "github.com/sumit-tembe/gin-requestid"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"time"
	"xor-go/pkg/xhttp"
	"xor-go/pkg/xshutdown"
	"xor-go/services/finances/internal/handler/generated"
)

// Start - Единая точка запуска приложения
func (a *App) Start(ctx context.Context) {

	go a.startHTTPServer(ctx)

	if err := xshutdown.Wait(a.cfg.GracefulShutdown); err != nil {
		a.logger.Error(fmt.Sprintf("Failed to gracefully shutdown %s app: %s", a.cfg.App.Service, err.Error()))
	} else {
		a.logger.Info("App gracefully stopped")
	}
}

func (a *App) startHTTPServer(ctx context.Context) {
	// Создаем общий роутинг http сервера
	router := xhttp.NewRouter()

	// TODO Добавляем системные роуты
	//router.WithHandleGET("/metrics", metrics.HandleFunc())

	tracerMw := generated.MiddlewareFunc(otelgin.Middleware(a.cfg.App.Service, otelgin.WithTracerProvider(a.tracerProvider)))
	GinZapMw := generated.MiddlewareFunc(ginzap.Ginzap(a.logger, time.RFC3339, true))
	requestIdMw := generated.MiddlewareFunc(requestid.RequestID(nil))
	middlewares := []generated.MiddlewareFunc{
		tracerMw,
		GinZapMw,
		requestIdMw,
	}

	// Добавляем роуты api
	xhttp.InitHandler(router.Router(), a.logger, middlewares, a.service, longPollTimeout)

	// Создаем сервер
	srv := xhttp.NewServer(a.cfg.Http, router)
	srv.RegisterRoutes(&router)

	// Стартуем
	a.logger.Info(fmt.Sprintf("Starting %s HTTP server at %s:%d", a.cfg.App.Service, a.cfg.Http.Host, a.cfg.Http.Port))
	if err := srv.Start(); err != nil {
		a.logger.Error(fmt.Sprintf("Fail with %s HTTP server: %s", a.cfg.App.Service, err.Error()))
		xshutdown.Now()
	}
}
