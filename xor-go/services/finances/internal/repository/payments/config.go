package payments

type Config struct {
	Base   string `mapstructure:"base" env:"base"`
	Prefix string `mapstructure:"prefix" env:"prefix"`
}
