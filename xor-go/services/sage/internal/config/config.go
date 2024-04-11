package config

import (
	"github.com/spf13/viper"
	"xor-go/pkg/xapp"
	"xor-go/pkg/xhttp"
	"xor-go/pkg/xlogger"
)

type IdmClientConfig struct {
	Host string `mapstructure:"host"`
}

type Config struct {
	App             *xapp.Config    `mapstructure:"app"`
	Http            *xhttp.Config   `mapstructure:"http"`
	Logger          *xlogger.Config `mapstructure:"logger"`
	IdmClientConfig IdmClientConfig `mapstructure:"idm-client"`
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
