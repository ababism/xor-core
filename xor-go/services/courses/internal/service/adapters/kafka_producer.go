package adapters

import (
	"context"
)

type KafkaClient interface {
	SendMessage(ctx context.Context, message string) error
}
