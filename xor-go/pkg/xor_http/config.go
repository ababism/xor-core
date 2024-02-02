package xor_http

import (
	"time"
)

const (
	defaultHost         = "localhost"
	defaultPort         = "8080"
	defaultReadTimeout  = time.Second
	defaultWriteTimeout = time.Second
)

type Config struct {
	Host         string        `yaml:"host" env:"HOST"`
	Port         string        `yaml:"port" env:"PORT"`
	ReadTimeout  time.Duration `yaml:"read_timeout" env:"READ_TIMEOUT"`
	WriteTimeout time.Duration `yaml:"write_timeout" env:"WRITE_TIMEOUT"`
}

func NewDefaultConfig() *Config {
	return &Config{
		Host:         defaultHost,
		Port:         defaultPort,
		ReadTimeout:  defaultReadTimeout,
		WriteTimeout: defaultWriteTimeout,
	}
}
