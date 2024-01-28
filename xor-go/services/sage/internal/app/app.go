package app

import (
	"fmt"
	"go.uber.org/zap"
	XorHttpServer "xor-go/pkg/xor_http_server"
	XorLogger "xor-go/pkg/xor_logger"
	"xor-go/services/sage/internal/config"
)

type Application struct {
	config *config.Config
	logger *zap.Logger
}

func NewApp(config *config.Config) (*Application, error) {
	logger, err := XorLogger.InitLogger(config.LoggerConfig)
	if err != nil {
		return nil, fmt.Errorf("failed initializing logger with error: {%v}", err)
	}

	return &Application{
		config: config,
		logger: logger,
	}, nil
}

func (a *Application) Start() {
	a.startHTTPServer()
}

func (a *Application) startHTTPServer() {
	router := XorHttpServer.NewRouter()
	httpServer := XorHttpServer.NewServer(a.config.HttpConfig, router)

	if err := httpServer.Start(); err != nil {
		a.logger.Error(err.Error())
	}
}
