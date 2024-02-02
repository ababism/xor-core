package app

import (
	"context"
	"go.uber.org/zap"
	xordbmongo "xor-go/pkg/xordb/mongo"
	"xor-go/pkg/xorhttp"
	xorhttpresponse "xor-go/pkg/xorhttp/response"
	"xor-go/pkg/xorlogger"
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
	logger, err := xorlogger.Init(cfg.LoggerConfig, cfg.SystemConfig)
	if err != nil {
		return nil, err
	}

	mongoClient, err := xordbmongo.NewClient(context.Background(), cfg.MongoConfig)
	if err != nil {
		return nil, err
	}

	httpResponser := xorhttpresponse.NewHttpResponseWrapper(logger)

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
	router := xorhttp.NewRouter()

	api := router.Router().Group("/api")
	r.accountHandler.InitAccountRoutes(api)

	httpServer := xorhttp.NewServer(r.config.HttpConfig, router)
	if err := httpServer.Start(); err != nil {
		r.logger.Error(err.Error())
	}
}
