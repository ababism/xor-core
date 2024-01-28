package config

import (
	XorHttpServer "xor-go/pkg/xor_http_server"
	XorLogger "xor-go/pkg/xor_logger"
)

type Config struct {
	LoggerConfig *XorLogger.Config     `yaml:"logger" env-prefix:"LOGGER"`
	HttpConfig   *XorHttpServer.Config `yaml:"http"  env-prefix:"HTTP"`
}
