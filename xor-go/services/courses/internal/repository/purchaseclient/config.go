package purchaseclient

type ClientConfig struct {
	Uri             string `mapstructure:"uri"`
	WebhookBasePath string `mapstructure:"webhook_base_path"`
}
