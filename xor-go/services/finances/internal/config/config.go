package config

import (
	"github.com/spf13/viper"
	"log"
	"xor-go/pkg/xapp"
	"xor-go/pkg/xconfig"
	"xor-go/pkg/xdb/postgres"
	"xor-go/pkg/xhttp"
	"xor-go/pkg/xlogger"
	"xor-go/pkg/xshutdown"
	"xor-go/pkg/xtracer"
	"xor-go/services/finances/internal/repository/payments"
)

type Config struct {
	App              *xapp.Config      `mapstructure:"app"`
	Http             *xhttp.Config     `mapstructure:"http"`
	Logger           *xlogger.Config   `mapstructure:"logger"`
	Postgres         *postgres.Config  `mapstructure:"postgres"`
	PaymentsClient   *payments.Config  `mapstructure:"payments_client"`
	GracefulShutdown *xshutdown.Config `mapstructure:"graceful_shutdown"`
	Tracer           *xtracer.Config   `mapstructure:"tracer"`
	//Metrics          *metrics.Config              `mapstructure:"metrics"`
}

func NewConfig(filePath string, appName string) (*Config, error) {
	viper.SetConfigFile(filePath)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error while reading config file: %v", err)
	}

	// Загрузка конфигурации в структуру Config
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("error while unmarshalling config file: %v", err)
	}

	// Замена значений из переменных окружения, если они заданы
	xconfig.ReplaceWithEnv(&config, appName)

	return &config, nil
}
