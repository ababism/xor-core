package service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/juju/zaputil/zapctx"
	global "go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"net/http"
	"xor-go/pkg/xapperror"
	"xor-go/services/courses/internal/domain"
	"xor-go/services/courses/internal/domain/keys"
)

// ReadLesson to get published lesson info
func (c CoursesService) ReadLesson(initialCtx context.Context, actor domain.Actor, lessonID uuid.UUID) (*domain.Lesson, error) {
	_ = zapctx.Logger(initialCtx)

	tr := global.Tracer(domain.ServiceName)
	ctx, span := tr.Start(initialCtx, "courses/service.ReadLesson")
	defer span.End()

	ToSpan(&span, actor)

	lesson, err := c.lesson.Get(ctx, lessonID)
	if err != nil {
		return nil, err
	}

	access, err := c.student.GetLessonAccess(ctx, actor.ID, lessonID)
	if err != nil {
		return nil, err
	}

	if access.AccessStatus != domain.Accessible && !actor.HasRole(domain.AdminRole) {
		lesson.ApplyPaywall()
	}

	if !actor.HasOneOfRoles(domain.TeacherRole, domain.AdminRole) {
		lesson.ApplyVisibility()
	}

	return lesson, nil
}

func (c CoursesService) CreateLesson(initialCtx context.Context, actor domain.Actor, lesson *domain.Lesson) (*domain.Lesson, error) {
	log := zapctx.Logger(initialCtx)

	tr := global.Tracer(domain.ServiceName)
	ctx, span := tr.Start(initialCtx, "courses/service.CreateLesson")
	defer span.End()

	ToSpan(&span, actor)

	if !actor.HasOneOfRoles(domain.TeacherRole, domain.AdminRole) {
		return nil, xapperror.New(http.StatusForbidden, "user does not have teacher rights to create lesson",
			fmt.Sprintf("user do not have %s or %s roles", domain.TeacherRole, domain.AdminRole), nil)
	}

	if lesson == nil {
		log.Error("lesson is nil")
		return nil, xapperror.New(http.StatusBadRequest, "lesson is nil", "lesson is nil", nil)
	}

	lesson.TeacherID = actor.ID
	lesson.Product.Owner = actor.ID

	if lesson.ID == uuid.Nil || (lesson.ID == uuid.UUID{}) {
		lesson.ID = uuid.New()
	}

	lesson.Product.Item = lesson.ID

	err := lesson.Validate()
	if err != nil {
		return nil, err
	}

	newLesson, err := c.lessonEdit.Create(ctx, lesson)
	if err != nil {
		return nil, err
	}

	span.AddEvent("lesson created", trace.WithAttributes(attribute.String(keys.LessonIDAttributeKey, newLesson.ID.String())))
	return newLesson, nil
}

func (c CoursesService) GetLesson(initialCtx context.Context, actor domain.Actor, lessonID uuid.UUID) (*domain.Lesson, error) {
	_ = zapctx.Logger(initialCtx)

	tr := global.Tracer(domain.ServiceName)
	ctx, span := tr.Start(initialCtx, "courses/service.GetLesson")
	defer span.End()

	ToSpan(&span, actor)

	if !actor.HasOneOfRoles(domain.TeacherRole, domain.AdminRole) {
		return nil, xapperror.New(http.StatusForbidden, "user does not have rights to see unpublished lesson",
			fmt.Sprintf("user do not have %s or %s roles", domain.TeacherRole, domain.AdminRole), nil)
	}

	lesson, err := c.lessonEdit.Get(ctx, lessonID)
	if err != nil {
		return nil, err
	}

	return lesson, nil
}

func (c CoursesService) UpdateLesson(initialCtx context.Context, actor domain.Actor, lessonID uuid.UUID, lesson *domain.Lesson) (*domain.Lesson, error) {
	log := zapctx.Logger(initialCtx)

	tr := global.Tracer(domain.ServiceName)
	ctx, span := tr.Start(initialCtx, "courses/service.UpdateLesson")
	defer span.End()

	ToSpan(&span, actor)

	if !actor.HasOneOfRoles(domain.TeacherRole, domain.AdminRole) {
		return nil, xapperror.New(http.StatusForbidden, "user does not have teacher rights to update lesson",
			fmt.Sprintf("user do not have %s or %s roles", domain.TeacherRole, domain.AdminRole), nil)
	}

	if lesson == nil {
		log.Error("lesson is nil")
		return nil, xapperror.New(http.StatusBadRequest, "lesson is nil", "lesson is nil", nil)
	}

	curLesson, err := c.lessonEdit.Get(ctx, lessonID)
	if err != nil {
		return nil, err
	}
	if curLesson.TeacherID != actor.ID {
		return nil, xapperror.New(http.StatusForbidden, "user does not have rights to update someone else's lesson",
			"lesson teacher id does not match actor id", nil)
	}

	lesson.TeacherID = actor.ID
	lesson.ID = lessonID

	err = lesson.Validate()
	if err != nil {
		return nil, err
	}

	err = c.lessonEdit.Update(ctx, lesson)
	if err != nil {
		return nil, err
	}

	span.AddEvent("lesson updated", trace.WithAttributes(attribute.String(keys.LessonIDAttributeKey, lesson.ID.String())))
	return lesson, nil
}

func (c CoursesService) DeleteLesson(initialCtx context.Context, actor domain.Actor, lessonID uuid.UUID) error {
	_ = zapctx.Logger(initialCtx)

	tr := global.Tracer(domain.ServiceName)
	ctx, span := tr.Start(initialCtx, "courses/service.DeleteLesson")
	defer span.End()

	ToSpan(&span, actor)

	if !actor.HasOneOfRoles(domain.TeacherRole, domain.AdminRole) {
		return xapperror.New(http.StatusForbidden, "user does not have teacher rights to delete lesson",
			fmt.Sprintf("user do not have %s or %s roles", domain.TeacherRole, domain.AdminRole), nil)
	}

	err := c.lessonEdit.Delete(ctx, lessonID)
	if err != nil {
		return err
	}

	span.AddEvent("lesson deleted", trace.WithAttributes(attribute.String(keys.LessonIDAttributeKey, lessonID.String())))
	return nil
}
