package adapters

import (
	"context"
	"github.com/google/uuid"
	"xor-go/services/finances/internal/domain"
)

type AccountRepository interface {
	GetTripByID(ctx context.Context, tripId uuid.UUID) (*domain.Trip, error)
	CreateAccount(ctx context.Context, account Account) error
}
