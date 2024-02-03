package config

import (
	"xor-go/pkg/app"
	"xor-go/pkg/db/mongo"
	"xor-go/pkg/http"
	"xor-go/pkg/logger"
)

type Config struct {
	SystemConfig *app.Config    `yaml:"system"`
	LoggerConfig *logger.Config `yaml:"logger"`
	HttpConfig   *http.Config   `yaml:"http"`
	MongoConfig  *mongo.Config  `yaml:"mongo"`
}
