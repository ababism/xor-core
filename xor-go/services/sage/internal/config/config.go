package config

import (
	"xor-go/pkg/xapp"
	"xor-go/pkg/xdb/mongo"
	"xor-go/pkg/xdb/postgres"
	"xor-go/pkg/xhttp"
	"xor-go/pkg/xlogger"
)

type Config struct {
	SystemConfig   *xapp.Config     `yaml:"system"`
	LoggerConfig   *xlogger.Config  `yaml:"xlogger"`
	HttpConfig     *xhttp.Config    `yaml:"xhttp"`
	MongoConfig    *mongo.Config    `yaml:"postgre"`
	PostgresConfig *postgres.Config `yaml:"postgres"`
}
