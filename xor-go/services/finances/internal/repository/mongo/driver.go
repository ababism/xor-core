package mongo

import (
	"context"
	"github.com/google/uuid"
	"github.com/juju/zaputil/zapctx"
	"gitlab.com/ArtemFed/mts-final-taxi/pkg/apperror"
	"gitlab.com/ArtemFed/mts-final-taxi/projects/template/internal/domain"
	"gitlab.com/ArtemFed/mts-final-taxi/projects/template/internal/repository/mongo/models"
	"go.mongodb.org/mongo-driver/bson"
	global "go.opentelemetry.io/otel"
	"go.uber.org/zap"
	"net/http"
)

// GetTripByID returns Trip from driver collection by trip_id
func (r *DriveRepository) GetTripByID(ctx context.Context, tripId uuid.UUID) (*domain.Trip, error) {
	log := zapctx.Logger(ctx)

	tr := global.Tracer(domain.ServiceName)
	newCtx, span := tr.Start(ctx, "driver/repository/mongo.GetTripByID")
	defer span.End()

	var trip models.MongoTrip
	filter := bson.M{"trip_id": tripId.String()}
	err := r.driverCollection.FindOne(newCtx, filter).Decode(&trip)
	if err != nil {
		appErr := apperror.NewAppError(http.StatusNotFound, "can't find trip", "error fetching trip from MongoDB", err)
		log.Error("error fetching trip from MongoDB", zap.Error(appErr))
		return nil, appErr
	}

	res, err := models.ToDomainTripModel(trip)
	if err != nil {
		appErr := apperror.NewAppError(http.StatusInternalServerError, "internal error",
			"error converting mongo to domain model", err)
		log.Error("error converting mongo to domain model", zap.Error(appErr))
		return nil, appErr
	}
	return res, nil
}

// InsertTrip inserts Trip
func (r *DriveRepository) InsertTrip(ctx context.Context, trip domain.Trip) error {
	logger := zapctx.Logger(ctx)

	tr := global.Tracer(domain.ServiceName)
	newCtx, span := tr.Start(ctx, "driver/repository/mongo.InsertTrip")
	defer span.End()

	mongoTrip := models.ToMongoTripModel(trip)
	_, err := r.driverCollection.InsertOne(newCtx, mongoTrip)
	if err != nil {
		appErr := apperror.NewAppError(http.StatusBadRequest, "trip already created",
			"failed to create trip in MongoDB", err)
		logger.Error("failed to create trip in MongoDB", zap.Error(appErr))
		return appErr
	}

	return nil
}
