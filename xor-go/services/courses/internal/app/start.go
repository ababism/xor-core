package app

import (
	"context"
	"fmt"
	ginzap "github.com/gin-contrib/zap"
	requestid "github.com/sumit-tembe/gin-requestid"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"time"
	//"xor-go/pkg/metrics"
	httpServer "xor-go/pkg/xhttp"
	"xor-go/pkg/xshutdown"
	"xor-go/services/courses/internal/handler/generated"
	myHttp "xor-go/services/courses/internal/handler/http"
)

// Start - Единая точка запуска приложения
func (a *App) Start(ctx context.Context) {

	go a.startHTTPServer(ctx)

	if err := xshutdown.Wait(a.cfg.GracefulShutdown); err != nil {
		a.logger.Error(fmt.Sprintf("Failed to gracefully shutdown %s app: %s", a.cfg.App.Name, err.Error()))
	} else {
		a.logger.Info("App gracefully stopped")
	}
}

func (a *App) startHTTPServer(ctx context.Context) {
	// Создаем общий роутинг http сервера
	router := httpServer.NewRouter()
	//
	//// Добавляем системные роуты
	//router.WithHandleGET("/metrics", metrics.HandleFunc())

	tracerMw := generated.MiddlewareFunc(otelgin.Middleware(a.cfg.App.Name, otelgin.WithTracerProvider(a.tracerProvider)))
	GinZapMw := generated.MiddlewareFunc(ginzap.Ginzap(a.logger, time.RFC3339, true))
	requestIdMw := generated.MiddlewareFunc(requestid.RequestID(nil))
	middlewares := []generated.MiddlewareFunc{
		tracerMw,
		GinZapMw,
		requestIdMw,
	}

	// Добавляем роуты api
	myHttp.InitHandler(router.Router(), a.logger, middlewares, a.service)

	// Создаем сервер
	srv := httpServer.NewServer(a.cfg.Http, router)
	//srv.RegisterRoutes(&router)

	// Стартуем
	a.logger.Info(fmt.Sprintf("Starting %s HTTP server at %s:%d", a.cfg.App.Name, a.cfg.Http.Host, a.cfg.Http.Port))
	if err := srv.Start(); err != nil {
		a.logger.Error(fmt.Sprintf("Fail with %s HTTP server: %s", a.cfg.App.Name, err.Error()))
		xshutdown.Now()
	}
}
