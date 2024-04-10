package config

type ResourceRoute struct {
	Pattern       string   `mapstructure:"pattern"`
	RequiredRoles []string `mapstructure:"required-roles"`
}

type ResourceConfig struct {
	Name   string          `mapstructure:"name"`
	Host   string          `mapstructure:"host"`
	Routes []ResourceRoute `mapstructure:"routes"`
}

type ResourcesConfig struct {
	Resources []ResourceConfig `mapstructure:"resources"`
}
