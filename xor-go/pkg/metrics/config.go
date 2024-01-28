package metrics

// FIXME remove this package. should use "code.cloudfoundry.org/go-metric-registry"

type Config struct {
	MetricsEnable  bool   `mapstructure:"enable"`
	MetricsAddress string `mapstructure:"address"`
}
