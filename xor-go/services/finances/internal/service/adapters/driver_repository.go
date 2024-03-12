package adapters

import (
	"context"
	"github.com/google/uuid"
	"gitlab.com/ArtemFed/mts-final-taxi/projects/template/internal/domain"
)

type DriverRepository interface {
	GetTripByID(ctx context.Context, tripId uuid.UUID) (*domain.Trip, error)
	InsertTrip(ctx context.Context, status domain.Trip) error
}
