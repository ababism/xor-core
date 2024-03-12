package adapters

import (
	"context"
	"github.com/google/uuid"
	"gitlab.com/ArtemFed/mts-final-taxi/projects/template/internal/domain"
)

type DriverService interface {
	GetTripByID(ctx context.Context, driverId uuid.UUID, tripId uuid.UUID) (*domain.Trip, error)
	InsertTrip(ctx context.Context, trip domain.Trip) error
	AcceptTrip(ctx context.Context, driverId uuid.UUID, tripId uuid.UUID) error
}
