package adapters

import (
	"context"

	"gitlab.com/ArtemFed/mts-final-taxi/projects/template/internal/domain"
)

type LocationClient interface {
	GetDrivers(ctx context.Context, centerLocation domain.LatLngLiteral, radius float32) ([]domain.DriverLocation, error)
}
