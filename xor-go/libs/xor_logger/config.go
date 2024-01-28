package xor_logger

type Config struct {
	Environment      string   `mapstructure:"environment"`
	Encoding         string   `mapstructure:"encoding"`
	Level            string   `mapstructure:"level"`
	OutputPaths      []string `mapstructure:"output_paths"`
	ErrorOutputPaths []string `mapstructure:"error_output_paths"`
}
