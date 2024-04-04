package service

import (
	"context"
	"github.com/google/uuid"
	"xor-go/services/courses/internal/domain"
)

//
//func (s *CoursesService) GetTripByID(ctx context.Context, driverId uuid.UUID, tripId uuid.UUID) (*domain.Trip, error) {
//	logger := zapctx.Logger(ctx)
//
//	tr := global.Tracer(domain.ServiceName)
//	newCtx, span := tr.Start(ctx, "driver/service.GetTripByID")
//	defer span.End()
//
//	// err if trip driver != nil and driver != driverId
//	trip, err := s.r.GetTripByID(newCtx, tripId)
//	if err != nil {
//		logger.Error("driver-service: get trip from repository")
//		return nil, err
//	}
//	if trip.DriverId != nil && *trip.DriverId != driverId.String() {
//		return nil, errors.Wrap(domain.ErrAccessDenied, "trip driver id does not match passed id")
//	}
//	return trip, err
//}

func (c CoursesService) CreateCourse(ctx context.Context, actor domain.Actor, course *domain.Course) (uuid.UUID, error) {
	//TODO implement me
	panic("implement me")
}

func (c CoursesService) GetCourse(ctx context.Context, actor domain.Actor, courseID uuid.UUID) (*domain.Course, error) {
	//TODO implement me
	panic("implement me")
}

func (c CoursesService) UpdateCourse(ctx context.Context, actor domain.Actor, courseID uuid.UUID, course *domain.Course) error {
	//TODO implement me
	panic("implement me")
}

func (c CoursesService) DeleteCourse(ctx context.Context, actor domain.Actor, courseID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (c CoursesService) ReadCourse(ctx context.Context, actor domain.Actor, courseID uuid.UUID) (*domain.Course, error) {
	//TODO implement me
	panic("implement me")
}

func (c CoursesService) ConfirmAccess(ctx context.Context, buyerID uuid.UUID, productIDs []uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}
