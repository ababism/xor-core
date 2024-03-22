package xapp

type Config struct {
	Name        string      `mapstructure:"name"`
	Environment Environment `mapstructure:"env"`
	Version     string      `mapstructure:"version"`
	Dc          string      `mapstructure:"dc"`
}

type RawConfig struct {
	Service     string `yaml:"service" env:"SERVICE"`
	Environment string `yaml:"environment" env:"ENVIRONMENT"`
	Dc          string `yaml:"dc" env:"DC"`
}

func (r *RawConfig) ToConfig() (*Config, error) {
	parsedEnv, err := ParseEnvironment(r.Environment)
	if err != nil {
		return nil, err
	}
	return &Config{
		Name:        r.Service,
		Environment: parsedEnv,
		Dc:          r.Dc,
	}, nil
}

func (r *Config) IsProduction() bool {
	return r.Environment == ProductionEnv
}

func (r *Config) IsTesting() bool {
	return r.Environment == TestingEnv
}

func (r *Config) IsDevelopment() bool {
	return r.Environment == DevelopmentEnv
}
