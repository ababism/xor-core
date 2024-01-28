package xor_logger

import (
	"fmt"
	"strings"
)

type Environment string

const (
	DevEnvironment     = "dev"
	ProdEnvironment    = "prod"
	UnknownEnvironment = "unknown"
)

func ParseEnvironment(env string) (Environment, error) {
	switch strings.ToLower(env) {
	case DevEnvironment:
		return DevEnvironment, nil
	case ProdEnvironment:
		return ProdEnvironment, nil
	default:
		return UnknownEnvironment, fmt.Errorf("unknown xor_logger environment: %s", env)
	}
}
