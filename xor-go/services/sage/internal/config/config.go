package config

import (
	"xor-go/pkg/xorapp"
	xordbmongo "xor-go/pkg/xordb/mongo"
	"xor-go/pkg/xorhttp"
	"xor-go/pkg/xorlogger"
)

type Config struct {
	SystemConfig *xorapp.Config     `yaml:"system"`
	LoggerConfig *xorlogger.Config  `yaml:"logger"`
	HttpConfig   *xorhttp.Config    `yaml:"http"`
	MongoConfig  *xordbmongo.Config `yaml:"mongo"`
}
