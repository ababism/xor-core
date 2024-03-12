package adapters

import (
	"context"
	"gitlab.com/ArtemFed/mts-final-taxi/projects/template/internal/domain"
)

type KafkaClient interface {
	SendUpdate(ctx context.Context, trip domain.Trip, commandType domain.CommandType, reason *string) error
}
