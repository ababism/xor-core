package config

import (
	"xor-go/pkg/xor_app"
	"xor-go/pkg/xor_db"
	"xor-go/pkg/xor_http"
	"xor-go/pkg/xor_log"
)

type Config struct {
	SystemConfig *xor_app.Config     `yaml:"system"`
	LoggerConfig *xor_log.Config     `yaml:"logger"`
	HttpConfig   *xor_http.Config    `yaml:"http"`
	MongoConfig  *xor_db.MongoConfig `yaml:"mongo"`
}
