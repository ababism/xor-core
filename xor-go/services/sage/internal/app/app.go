package app

import (
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"xor-go/pkg/xhttp"
	"xor-go/pkg/xhttp/response"
	"xor-go/pkg/xlogger"
	idmproto "xor-go/proto/idm"
	"xor-go/services/sage/internal/api/http/handler"
	"xor-go/services/sage/internal/config"
	"xor-go/services/sage/internal/service"
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

	idmGrpcClient, err := getIdmGrpcClient(cfg.IdmClientConfig.Host)
	if err != nil {
		return nil, err
	}

	gatewayResourceService := service.NewGatewayResourceService(resourceToConfig, idmGrpcClient)

	gatewayHandler := handler.NewGatewayHandler(
		logger,
		httpResponser,
		gatewayResourceService,
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

func getIdmGrpcClient(host string) (idmproto.IdmClient, error) {
	conn, err := grpc.Dial(host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return idmproto.NewIdmClient(conn), nil
}
