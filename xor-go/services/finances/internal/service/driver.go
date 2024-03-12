package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/juju/zaputil/zapctx"
	"github.com/pkg/errors"
	"gitlab.com/ArtemFed/mts-final-taxi/projects/template/internal/domain"
	"gitlab.com/ArtemFed/mts-final-taxi/projects/template/internal/service/adapters"
	global "go.opentelemetry.io/otel"
	"go.uber.org/zap"
)

var _ adapters.DriverService = &driverService{}

type driverService struct {
	r              adapters.DriverRepository
	locationClient adapters.LocationClient
	kafkaClient    adapters.KafkaClient
}

func NewDriverService(driverRepository adapters.DriverRepository, locationClient adapters.LocationClient, kafkaClient adapters.KafkaClient) adapters.DriverService {
	return &driverService{
		r:              driverRepository,
		locationClient: locationClient,
		kafkaClient:    kafkaClient,
	}
}

func (s *driverService) GetTripByID(ctx context.Context, driverId uuid.UUID, tripId uuid.UUID) (*domain.Trip, error) {
	logger := zapctx.Logger(ctx)

	tr := global.Tracer(domain.ServiceName)
	newCtx, span := tr.Start(ctx, "driver/service.GetTripByID")
	defer span.End()

	// err if trip driver != nil and driver != driverId
	trip, err := s.r.GetTripByID(newCtx, tripId)
	if err != nil {
		logger.Error("driver-service: get trip from repository")
		return nil, err
	}
	if trip.DriverId != nil && *trip.DriverId != driverId.String() {
		return nil, errors.Wrap(domain.ErrAccessDenied, "trip driver id does not match passed id")
	}
	return trip, err
}

// InsertTrip inserts a trip
func (s *driverService) InsertTrip(ctx context.Context, trip domain.Trip) error {
	log := zapctx.Logger(ctx)

	tr := global.Tracer(domain.ServiceName)
	newCtx, span := tr.Start(ctx, "driver/service.InsertTrip")
	defer span.End()

	err := s.r.InsertTrip(newCtx, trip)
	if err != nil {
		log.Error("driver-service: insert trip in repository", zap.Error(err))
		return err
	}

	return nil
}

func (s *driverService) AcceptTrip(ctx context.Context, driverId uuid.UUID, tripId uuid.UUID) error {
	logger := zapctx.Logger(ctx)

	tr := global.Tracer(domain.ServiceName)
	newCtx, span := tr.Start(ctx, "driver/service.AcceptTrip")
	defer span.End()

	trip, err := s.r.GetTripByID(newCtx, tripId)
	if err != nil || trip == nil {
		logger.Error("can't get trip from repository", zap.Error(err))
		return err
	}

	if trip.Status == nil || *trip.Status != domain.TripStatuses.GetDriverSearch() {
		return errors.Wrap(domain.ErrAccessDenied, "trip doesn't need driver")
	}
	dId := driverId.String()
	trip.DriverId = &dId
	err = s.kafkaClient.SendUpdate(newCtx, *trip, domain.TripCommandAccept, nil)
	if err != nil {
		logger.Error("can't send accept trip to kafka:", zap.Error(err))
		return err
	}

	return nil
}
