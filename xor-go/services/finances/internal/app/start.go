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
	"xor-go/services/finances/internal/handler/http"
	"xor-go/services/finances/internal/log"
)

// Start - Единая точка запуска приложения
func (a *App) Start(ctx context.Context) {

	go a.startHTTPServer(ctx)

	if err := xshutdown.Wait(a.cfg.GracefulShutdown); err != nil {
		log.Logger.Error(fmt.Sprintf("Failed to gracefully shutdown %s app: %s", a.cfg.App.Name, err.Error()))
	} else {
		log.Logger.Info("App gracefully stopped")
	}
}

func (a *App) startHTTPServer(ctx context.Context) {
	// Создаем общий роутинг http сервера
	router := xhttp.NewRouter()

	// TODO Добавляем системные роуты
	//router.WithHandleGET("/metrics", metrics.HandleFunc())

	tracerMw := http.MiddlewareFunc(otelgin.Middleware(a.cfg.App.Name, otelgin.WithTracerProvider(a.tracerProvider)))
	GinZapMw := http.MiddlewareFunc(ginzap.Ginzap(log.Logger, time.RFC3339, true))
	requestIdMw := http.MiddlewareFunc(requestid.RequestID(nil))
	middlewares := []http.MiddlewareFunc{
		tracerMw,
		GinZapMw,
		requestIdMw,
	}

	// TODO Добавляем роуты api
	http.InitHandler(a.handler, router.Router(), middlewares, "finances")

	// TODO Создаем сервер
	srv := xhttp.NewServer(a.cfg.Http, router)
	//srv.RegisterRoutes(&router)

	// Стартуем
	log.Logger.Info(fmt.Sprintf("Starting %s HTTP server at %s:%d", a.cfg.App.Name, a.cfg.Http.Host, a.cfg.Http.Port))
	if err := srv.Start(); err != nil {
		log.Logger.Error(fmt.Sprintf("Fail with %s HTTP server: %s", a.cfg.App.Name, err.Error()))
		xshutdown.Now()
	}
}
