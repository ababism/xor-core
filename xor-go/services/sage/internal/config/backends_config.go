package config

type PlatformService struct {
	Name string `mapstructure:"name"`
	Host string `mapstructure:"host"`
}

type PlatformServicesConfig struct {
	PlatformServices []PlatformService `mapstructure:"services"`
}
