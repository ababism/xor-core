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

// ReadLesson to get published lesson info
func (c CoursesService) ReadLesson(initialCtx context.Context, actor domain.Actor, lessonID uuid.UUID) (*domain.Lesson, error) {
	_ = zapctx.Logger(initialCtx)

	tr := global.Tracer(domain.ServiceName)
	ctx, span := tr.Start(initialCtx, "courses/service.ReadCourse")
	defer span.End()

	lesson, err := c.lessonEdit.Get(ctx, lessonID)
	if err != nil {
		return nil, err
	}

	if !actor.HasOneOfRoles(domain.TeacherRole, domain.AdminRole) {
		return nil, apperror.New(http.StatusForbidden, "user does not have teacher rights to create course",
			fmt.Sprintf("user do not have %s or %s roles", domain.TeacherRole, domain.AdminRole), nil)
	}

	return lesson, nil
}

func (c CoursesService) CreateLesson(initialCtx context.Context, actor domain.Actor, lesson *domain.Lesson) error {
	//TODO implement me
	panic("implement me")
}

func (c CoursesService) GetLesson(initialCtx context.Context, actor domain.Actor, lessonID uuid.UUID) (*domain.Lesson, error) {
	//TODO implement me
	panic("implement me")
}

func (c CoursesService) UpdateLesson(initialCtx context.Context, actor domain.Actor, lesson *domain.Lesson) error {
	//TODO implement me
	panic("implement me")
}

func (c CoursesService) DeleteLesson(initialCtx context.Context, actor domain.Actor, lessonID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}
