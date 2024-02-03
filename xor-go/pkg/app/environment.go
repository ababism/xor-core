package app

import (
	"fmt"
	"strings"
)

const (
	DevEnvironment     = "dev"
	TestEnvironment    = "test"
	ProdEnvironment    = "prod"
	UnknownEnvironment = "unknown"
)

type Environment string

func ParseEnvironment(env string) (Environment, error) {
	switch strings.ToLower(env) {
	case DevEnvironment:
		return DevEnvironment, nil
	case TestEnvironment:
		return TestEnvironment, nil
	case ProdEnvironment:
		return ProdEnvironment, nil
	default:
		return UnknownEnvironment, fmt.Errorf("got unsupported environment: %s", env)
	}
}
