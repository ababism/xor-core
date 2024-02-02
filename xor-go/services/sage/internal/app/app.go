package app

import (
	"context"
	"go.uber.org/zap"
	"xor-go/pkg/xor_db"
	"xor-go/pkg/xor_http"
	"xor-go/pkg/xor_http/response"
	"xor-go/pkg/xor_log"
	"xor-go/services/sage/internal/config"
	"xor-go/services/sage/internal/handler"
	"xor-go/services/sage/internal/repository"
	"xor-go/services/sage/internal/service"
)

type Application struct {
	config         *config.Config
	logger         *zap.Logger
	accountHandler *handler.AccountHandler
}

func NewApp(cfg *config.Config) (*Application, error) {
	logger, err := xor_log.Init(cfg.LoggerConfig, cfg.SystemConfig)
	if err != nil {
		return nil, err
	}

	mongoClient, err := xor_db.NewMongoClient(context.Background(), cfg.MongoConfig)
	if err != nil {
		return nil, err
	}

	httpResponser := response.NewHttpResponseWrapper(logger)

	db := mongoClient.Database(cfg.MongoConfig.Database)
	accountRepository := repository.NewAccountMongoRepository(logger, db)
	accountService := service.NewAccountService(accountRepository)
	accountHandler := handler.NewAccountHandler(httpResponser, accountService)

	return &Application{
		config:         cfg,
		logger:         logger,
		accountHandler: accountHandler,
	}, nil
}

func (r *Application) Start() {
	r.startHTTPServer()
}

func (r *Application) startHTTPServer() {
	router := xor_http.NewRouterWithSystemHandlers()

	api := router.Router().Group("/api")
	r.accountHandler.InitAccountRoutes(api)

	httpServer := xor_http.NewServer(r.config.HttpConfig, router)
	if err := httpServer.Start(); err != nil {
		r.logger.Error(err.Error())
	}
}
