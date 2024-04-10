package app

import (
	"go.uber.org/zap"
	"xor-go/pkg/xhttp"
	"xor-go/pkg/xhttp/response"
	"xor-go/pkg/xlogger"
	"xor-go/services/sage/internal/api/http/handler"
	"xor-go/services/sage/internal/config"
	"xor-go/services/sage/internal/service"
	"xor-go/services/sage/pkg/idm"
)

type Application struct {
	config         *config.Config
	logger         *zap.Logger
	securedHandler *handler.GatewayHandler
}

func NewApp(cfg *config.Config, servicesCfg *config.ResourcesConfig) (*Application, error) {
	logger, err := xlogger.Init(cfg.Logger, cfg.App)
	if err != nil {
		return nil, err
	}
	httpResponser := response.NewHttpResponseWrapper(logger)

	resourceToConfig := getResourceToConfig(servicesCfg)

	idmClient := idm.NewIdmClient(cfg.IdmClientConfig.Host)
	gatewayResourceService := service.NewGatewayResourceService(resourceToConfig, idmClient)

	gatewayHandler := handler.NewGatewayHandler(
		logger,
		httpResponser,
		gatewayResourceService,
		idmClient,
	)

	return &Application{
		config:         cfg,
		logger:         logger,
		securedHandler: gatewayHandler,
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

func getResourceToConfig(cfg *config.ResourcesConfig) map[string]*config.ResourceConfig {
	resourceToConfig := make(map[string]*config.ResourceConfig)
	for i := 0; i < len(cfg.Resources); i++ {
		resource := cfg.Resources[i]
		resourceToConfig[resource.Name] = &resource
	}
	return resourceToConfig
}
