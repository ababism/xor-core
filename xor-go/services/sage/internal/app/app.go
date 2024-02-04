package app

import (
	"go.uber.org/zap"
	"xor-go/pkg/xdb/postgres"
	"xor-go/pkg/xhttp"
	httpresponse "xor-go/pkg/xhttp/response"
	"xor-go/pkg/xlogger"
	"xor-go/services/sage/internal/config"
	"xor-go/services/sage/internal/handler"
	repopostgres "xor-go/services/sage/internal/repository/postgres"
	"xor-go/services/sage/internal/service"
)

type Application struct {
	config         *config.Config
	logger         *zap.Logger
	accountHandler *handler.AccountHandler
}

func NewApp(cfg *config.Config) (*Application, error) {
	logger, err := xlogger.Init(cfg.LoggerConfig, cfg.SystemConfig)
	if err != nil {
		return nil, err
	}

	httpResponser := httpresponse.NewHttpResponseWrapper(logger)

	postgresDb, err := postgres.NewDB(cfg.PostgresConfig)
	if err != nil {
		return nil, err
	}

	accountRepository := repopostgres.NewAccountPostgresRepository(logger, postgresDb)
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
	router := xhttp.NewRouter()

	api := router.Router().Group("/api")
	r.accountHandler.InitRoutes(api)

	httpServer := xhttp.NewServer(r.config.HttpConfig, router)
	if err := httpServer.Start(); err != nil {
		r.logger.Error(err.Error())
	}
}
