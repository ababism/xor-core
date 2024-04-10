package config

import (
	"github.com/spf13/viper"
	"xor-go/pkg/xapp"
	"xor-go/pkg/xhttp"
	"xor-go/pkg/xlogger"
)

type Config struct {
	App    *xapp.Config    `mapstructure:"app"`
	Http   *xhttp.Config   `mapstructure:"http"`
	Logger *xlogger.Config `mapstructure:"logger"`
	//Postgres         *postgres.Config  `mapstructure:"postgres"`
	//PaymentsClient   *payments.Config  `mapstructure:"payments_client"`
	//GracefulShutdown *xshutdown.Config `mapstructure:"graceful_shutdown"`
	//Tracer           *xtracer.Config   `mapstructure:"tracer"`
	//Metrics          *metrics.Config              `mapstructure:"metrics"`
}

func ParseConfig(configPath string, config interface{}) error {
	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(config); err != nil {
		return err
	}

	return nil
}

//func NewConfig(filePath string) (*Config, error) {
//	viper.SetConfigFile(filePath)
//	if err := viper.ReadInConfig(); err != nil {
//		log.Fatalf("Failed to read config: %v", err)
//	}
//
//	var config Config
//	if err := viper.Unmarshal(&config); err != nil {
//		log.Fatalf("Failed to unmarshal config: %v", err)
//	}
//
//	return &config, nil
//}
