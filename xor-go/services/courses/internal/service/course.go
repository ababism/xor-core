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

func (c CoursesService) CreateCourse(initialCtx context.Context, actor domain.Actor, course *domain.Course) (*domain.Course, error) {
	_ = zapctx.Logger(initialCtx)

	tr := global.Tracer(domain.ServiceName)
	ctx, span := tr.Start(initialCtx, "courses/service.CreateCourse")
	defer span.End()

	if !actor.HasOneOfRoles(domain.TeacherRole, domain.AdminRole) {
		return nil, apperror.New(http.StatusForbidden, "user does not have teacher rights to create course",
			fmt.Sprintf("user do not have %s or %s roles", domain.TeacherRole, domain.AdminRole), nil)
	}

	err := course.Validate()
	if err != nil {
		return nil, err
	}

	newCourse, err := c.courseEdit.Create(ctx, course)
	if err != nil {
		return nil, err
	}

	return newCourse, nil
}

func (c CoursesService) GetCourse(initialCtx context.Context, actor domain.Actor, courseID uuid.UUID) (*domain.Course, error) {
	_ = zapctx.Logger(initialCtx)

	tr := global.Tracer(domain.ServiceName)
	ctx, span := tr.Start(initialCtx, "courses/service.GetCourse")
	defer span.End()

	if !actor.HasOneOfRoles(domain.TeacherRole, domain.AdminRole) {
		return nil, apperror.New(http.StatusForbidden, "user does not have rights to read unpublished course",
			fmt.Sprintf("user do not have %s or %s roles", domain.TeacherRole, domain.AdminRole), nil)
	}
	course, err := c.courseEdit.Get(ctx, courseID)
	if err != nil {
		return nil, err
	}

	return course, nil
}

func (c CoursesService) UpdateCourse(initialCtx context.Context, actor domain.Actor, courseID uuid.UUID, course *domain.Course) (*domain.Course, error) {
	_ = zapctx.Logger(initialCtx)

	tr := global.Tracer(domain.ServiceName)
	ctx, span := tr.Start(initialCtx, "courses/service.UpdateCourse")
	defer span.End()

	if !actor.HasOneOfRoles(domain.TeacherRole, domain.AdminRole) {
		return nil, apperror.New(http.StatusForbidden, "user does not have teacher rights to update course",
			fmt.Sprintf("user do not have %s or %s roles", domain.TeacherRole, domain.AdminRole), nil)
	}

	err := course.Validate()
	if err != nil {
		return nil, err
	}

	err = c.courseEdit.Update(ctx, courseID, course)
	if err != nil {
		return nil, err
	}

	return course, nil
}

func (c CoursesService) DeleteCourse(initialCtx context.Context, actor domain.Actor, courseID uuid.UUID) error {
	_ = zapctx.Logger(initialCtx)

	tr := global.Tracer(domain.ServiceName)
	ctx, span := tr.Start(initialCtx, "courses/service.DeleteCourse")
	defer span.End()

	if !actor.HasOneOfRoles(domain.TeacherRole, domain.AdminRole) {
		return apperror.New(http.StatusForbidden, "user does not have teacher rights to delete course",
			fmt.Sprintf("user do not have %s or %s roles", domain.TeacherRole, domain.AdminRole), nil)
	}

	err := c.courseEdit.Delete(ctx, courseID)
	if err != nil {
		return err
	}

	return nil
}

func (c CoursesService) ReadCourse(initialCtx context.Context, actor domain.Actor, courseID uuid.UUID) (*domain.Course, error) {
	_ = zapctx.Logger(initialCtx)

	tr := global.Tracer(domain.ServiceName)
	ctx, span := tr.Start(initialCtx, "courses/service.ReadCourse")
	defer span.End()

	course, err := c.course.Get(ctx, courseID)
	if err != nil {
		return nil, err
	}

	if !actor.HasOneOfRoles(domain.TeacherRole, domain.AdminRole) {
		course.ApplyVisibility()
	}

	return course, nil
}

func (c CoursesService) ConfirmAccess(initialCtx context.Context, buyerID uuid.UUID, products []domain.Product) error {
	_ = zapctx.Logger(initialCtx)

	tr := global.Tracer(domain.ServiceName)
	ctx, span := tr.Start(initialCtx, "courses/service.ConfirmAccess")
	defer span.End()

	for _, pr := range products {
		err := c.student.CreateAccessToLesson(ctx, buyerID, pr.Item, domain.Accessible)
		if err != nil {
			return err
		}
	}

	return nil
}
