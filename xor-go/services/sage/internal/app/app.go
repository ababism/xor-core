package app

import (
	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
	"xor-go/pkg/xhttp"
	"xor-go/pkg/xhttp/response"
	"xor-go/pkg/xlogger"
	"xor-go/services/sage/internal/api/http/handler"
	"xor-go/services/sage/internal/config"
	"xor-go/services/sage/pkg/idm"
)

type Application struct {
	config         *config.Config
	logger         *zap.Logger
	securedHandler *handler.SecuredHandler
	//accountHandler *handler.AccountHandler
}

func NewApp(cfg *config.Config, servicesCfg *config.PlatformServicesConfig) (*Application, error) {
	logger, err := xlogger.Init(cfg.Logger, cfg.App)
	if err != nil {
		return nil, err
	}

	restyClient := resty.New()
	idmClient := idm.NewClient("http://localhost:5758", restyClient)

	//client := idm.NewClient("http://localhost:5758")
	//
	httpResponser := response.NewHttpResponseWrapper(logger)
	//
	//postgresDb, err := postgres.NewDB(cfg.PostgresConfig)
	//if err != nil {
	//	return nil, err
	//}
	//
	//accountRepository := repopostgres.NewAccountPostgresRepository(logger, postgresDb)
	//accountService := service.NewAccountService(logger, accountRepository)
	//accountHandler := handler.NewAccountHandler(httpResponser, accountService)
	//

	securedHandler := handler.NewSecuredHandler(logger, httpResponser, idmClient)

	return &Application{
		config:         cfg,
		logger:         logger,
		securedHandler: securedHandler,
		//accountHandler: accountHandler,
	}, nil
}

func (r *Application) Start() {
	r.startHTTPServer()
}

func (r *Application) startHTTPServer() {
	router := xhttp.NewRouter()

	api := router.Router().Group("/api")
	r.securedHandler.InitRoutes(api)

	httpServer := xhttp.NewServer(r.config.Http, router)
	if err := httpServer.Start(); err != nil {
		r.logger.Error(err.Error())
	}
}
