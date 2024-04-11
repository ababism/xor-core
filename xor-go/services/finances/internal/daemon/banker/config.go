package banker

import "time"

type Config struct {
	interval string `mapstructure:"interval"`
}

func (c Config) GetInterval() (time.Duration, error) {
	return time.ParseDuration(c.interval)
}
