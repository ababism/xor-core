package service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/juju/zaputil/zapctx"
	global "go.opentelemetry.io/otel"
	"net/http"
	"xor-go/pkg/apperror"
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

func (c CoursesService) CreateCourse(initialCtx context.Context, actor domain.Actor, course *domain.Course) (*domain.Course, error) {
	_ = zapctx.Logger(initialCtx)

	tr := global.Tracer(domain.ServiceName)
	ctx, span := tr.Start(initialCtx, "courses/service.CreateCourse")
	defer span.End()

	if actor.HasRole(domain.TeacherRole) || actor.HasRole(domain.AdminRole) {
		return nil, apperror.New(http.StatusForbidden, "user does not have teacher rights to create course",
			fmt.Sprintf("user do not have %s or %s roles", domain.TeacherRole, domain.AdminRole), nil)
	}

	err := course.Validate()
	if err != nil {
		return nil, err
	}

	newCourse, err := c.course.Create(ctx, course)
	if err != nil {
		return nil, err
	}

	return newCourse, nil
}

func (c CoursesService) GetCourse(initialCtx context.Context, actor domain.Actor, courseID uuid.UUID) (*domain.Course, error) {
	//TODO implement me
	panic("implement me")
}

func (c CoursesService) UpdateCourse(initialCtx context.Context, actor domain.Actor, courseID uuid.UUID, course *domain.Course) error {
	//TODO implement me
	panic("implement me")
}

func (c CoursesService) DeleteCourse(initialCtx context.Context, actor domain.Actor, courseID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (c CoursesService) ReadCourse(initialCtx context.Context, actor domain.Actor, courseID uuid.UUID) (*domain.Course, error) {
	//TODO implement me
	panic("implement me")
}

func (c CoursesService) ConfirmAccess(initialCtx context.Context, buyerID uuid.UUID, productIDs []uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}
