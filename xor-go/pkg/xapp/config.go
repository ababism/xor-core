package xapp

type Config struct {
	Service     string
	Environment Environment
	Dc          string
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
		Service:     r.Service,
		Environment: parsedEnv,
		Dc:          r.Dc,
	}, nil
}

func (r *Config) IsProdEnvironment() bool {
	return r.Environment == ProdEnvironment
}

func (r *Config) IsTestEnvironment() bool {
	return r.Environment == TestEnvironment
}

func (r *Config) IsDevEnvironment() bool {
	return r.Environment == DevEnvironment
}
