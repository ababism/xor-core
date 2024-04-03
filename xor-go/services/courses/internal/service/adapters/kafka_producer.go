package adapters

import (
	"context"
)

type KafkaClient interface {
	SendUpdate(ctx context.Context, message string, topic string) error
}
