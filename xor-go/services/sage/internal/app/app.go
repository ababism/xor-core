package app

import (
	"context"
	"go.uber.org/zap"
	"xor-go/pkg/db/mongo"
	"xor-go/pkg/http"
	httpresponse "xor-go/pkg/http/response"
	"xor-go/pkg/logger"
	"xor-go/services/sage/internal/config"
	"xor-go/services/sage/internal/handler"
	repomongo "xor-go/services/sage/internal/repository/mongo"
	"xor-go/services/sage/internal/service"
)

type Application struct {
	config         *config.Config
	logger         *zap.Logger
	accountHandler *handler.AccountHandler
}

func NewApp(cfg *config.Config) (*Application, error) {
	logger, err := logger.Init(cfg.LoggerConfig, cfg.SystemConfig)
	if err != nil {
		return nil, err
	}

	mongoClient, err := mongo.NewClient(context.Background(), cfg.MongoConfig)
	if err != nil {
		return nil, err
	}

	httpResponser := httpresponse.NewHttpResponseWrapper(logger)

	db := mongoClient.Database(cfg.MongoConfig.Database)
	accountRepository := repomongo.NewAccountMongoRepository(logger, db)
	accountService := service.NewAccountService(logger, accountRepository)
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
	router := http.NewRouter()

	api := router.Router().Group("/api")
	r.accountHandler.InitAccountRoutes(api)

	httpServer := http.NewServer(r.config.HttpConfig, router)
	if err := httpServer.Start(); err != nil {
		r.logger.Error(err.Error())
	}
}
