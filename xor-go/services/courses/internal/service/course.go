package service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/juju/zaputil/zapctx"
	global "go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"net/http"
	"xor-go/pkg/xapperror"
	"xor-go/services/courses/internal/domain"
	"xor-go/services/courses/internal/domain/keys"
)

func (c CoursesService) CreateCourse(initialCtx context.Context, actor domain.Actor, course *domain.Course) (*domain.Course, error) {
	_ = zapctx.Logger(initialCtx)

	tr := global.Tracer(domain.ServiceName)
	ctx, span := tr.Start(initialCtx, "courses/service.CreateCourse")
	defer span.End()

	ToSpan(&span, actor)

	if !actor.HasOneOfRoles(domain.TeacherRole, domain.AdminRole) {
		return nil, xapperror.New(http.StatusForbidden, "user does not have teacher rights to create course",
			fmt.Sprintf("user do not have %s or %s roles", domain.TeacherRole, domain.AdminRole), nil)
	}

	course.TeacherID = actor.ID

	//if course.ID == uuid.Nil || (course.ID == uuid.UUID{}) {
	//	course.ID = uuid.New()
	//}
	course.FillEmptyUUIDs()

	err := course.Validate()
	if err != nil {
		return nil, err
	}

	newCourse, err := c.courseEdit.Create(ctx, course)
	if err != nil {
		return nil, err
	}

	span.AddEvent("course created", trace.WithAttributes(attribute.String(keys.CourseIDAttributeKey, newCourse.ID.String())))
	return newCourse, nil
}

func (c CoursesService) GetCourse(initialCtx context.Context, actor domain.Actor, courseID uuid.UUID) (*domain.Course, error) {
	_ = zapctx.Logger(initialCtx)

	tr := global.Tracer(domain.ServiceName)
	ctx, span := tr.Start(initialCtx, "courses/service.GetCourse")
	defer span.End()

	ToSpan(&span, actor)

	if !actor.HasOneOfRoles(domain.TeacherRole, domain.AdminRole) {
		return nil, xapperror.New(http.StatusForbidden, "user does not have rights to read unpublished course",
			fmt.Sprintf("user do not have %s or %s roles", domain.TeacherRole, domain.AdminRole), nil)
	}
	course, err := c.courseEdit.Get(ctx, courseID)
	if err != nil {
		return nil, err
	}

	return course, nil
}

func (c CoursesService) UpdateCourse(initialCtx context.Context, actor domain.Actor, courseID uuid.UUID, course *domain.Course) (*domain.Course, error) {
	log := zapctx.Logger(initialCtx)

	tr := global.Tracer(domain.ServiceName)
	ctx, span := tr.Start(initialCtx, "courses/service.UpdateCourse")
	defer span.End()

	ToSpan(&span, actor)

	if !actor.HasOneOfRoles(domain.TeacherRole, domain.AdminRole) {
		return nil, xapperror.New(http.StatusForbidden, "user does not have teacher rights to update course",
			fmt.Sprintf("user do not have %s or %s roles", domain.TeacherRole, domain.AdminRole), nil)
	}

	course.ID = courseID
	err := course.Validate()
	if err != nil {
		return nil, err
	}

	ok, errAccess := c.teacher.IsCourseAccessible(ctx, actor.ID, courseID)
	if errAccess != nil {
		return nil, xapperror.New(http.StatusForbidden, "user does not own this course",
			fmt.Sprintf("user does %s not own this course %s", actor.ID, courseID), errAccess)
	}

	if !actor.HasRole(domain.AdminRole) && !ok {
		log.Info("Bool !(actor.HasRole(domain.AdminRole) || ok) and res", zap.Bools("", []bool{actor.HasRole(domain.AdminRole), ok, !actor.HasRole(domain.AdminRole) || !ok}))
		return nil, xapperror.New(http.StatusForbidden, "user does not own course",
			fmt.Sprintf("user do not have %s roles or do not own course", domain.AdminRole), nil)
	}

	err = c.courseEdit.Update(ctx, courseID, course)
	if err != nil {
		return nil, err
	}

	span.AddEvent("course updated", trace.WithAttributes(attribute.String(keys.CourseIDAttributeKey, course.ID.String())))
	return course, nil
}

func (c CoursesService) DeleteCourse(initialCtx context.Context, actor domain.Actor, courseID uuid.UUID) error {
	_ = zapctx.Logger(initialCtx)

	tr := global.Tracer(domain.ServiceName)
	ctx, span := tr.Start(initialCtx, "courses/service.DeleteCourse")
	defer span.End()

	ToSpan(&span, actor)

	if !actor.HasOneOfRoles(domain.TeacherRole, domain.AdminRole) {
		return xapperror.New(http.StatusForbidden, "user does not have teacher rights to delete course",
			fmt.Sprintf("user do not have %s or %s roles", domain.TeacherRole, domain.AdminRole), nil)
	}

	err := c.courseEdit.Delete(ctx, courseID)
	if err != nil {
		return err
	}
	span.AddEvent("course deleted", trace.WithAttributes(attribute.String(keys.CourseIDAttributeKey, courseID.String())))

	return nil
}

func (c CoursesService) ReadCourse(initialCtx context.Context, actor domain.Actor, courseID uuid.UUID) (*domain.Course, error) {
	_ = zapctx.Logger(initialCtx)

	tr := global.Tracer(domain.ServiceName)
	ctx, span := tr.Start(initialCtx, "courses/service.ReadCourse")
	defer span.End()

	ToSpan(&span, actor)

	course, err := c.course.Get(ctx, courseID)
	if err != nil {
		return nil, err
	}

	if !actor.HasOneOfRoles(domain.TeacherRole, domain.AdminRole) {
		course.ApplyVisibility()
	}

	return course, nil
}

func (c CoursesService) GetTeachersPublishedCourses(initialCtx context.Context, actor domain.Actor, teacherID uuid.UUID, offset, limit int) ([]*domain.Course, error) {
	_ = zapctx.Logger(initialCtx)

	tr := global.Tracer(domain.ServiceName)
	ctx, span := tr.Start(initialCtx, "courses/service.GetTeachersPublishedCourses")
	defer span.End()

	ToSpan(&span, actor)

	courses, err := c.course.GetAllByTeacher(ctx, teacherID, offset, limit)
	if err != nil {
		return nil, err
	}

	if !((actor.HasRole(domain.TeacherRole) && actor.ID == teacherID) || actor.HasRole(domain.AdminRole)) {
		for _, course := range courses {
			course.ApplyVisibility()
		}
	}

	return courses, nil
}

func (c CoursesService) GetPublishedCoursesList(initialCtx context.Context, actor domain.Actor, offset, limit int) ([]*domain.Course, error) {
	_ = zapctx.Logger(initialCtx)

	tr := global.Tracer(domain.ServiceName)
	ctx, span := tr.Start(initialCtx, "courses/service.GetPublishedCoursesList")
	defer span.End()

	ToSpan(&span, actor)

	courses, err := c.course.GetAll(ctx, offset, limit)
	if err != nil {
		return nil, err
	}

	if !actor.HasRole(domain.AdminRole) {
		for _, course := range courses {
			course.ApplyVisibility()
		}
	}

	return courses, nil
}

func (c CoursesService) GetTeacherCoursesTemplates(initialCtx context.Context, actor domain.Actor, teacherID uuid.UUID, offset, limit int) ([]*domain.Course, error) {
	_ = zapctx.Logger(initialCtx)

	tr := global.Tracer(domain.ServiceName)
	ctx, span := tr.Start(initialCtx, "courses/service.GetTeacherCoursesTemplates")
	defer span.End()

	ToSpan(&span, actor)

	if !((actor.HasRole(domain.TeacherRole) && actor.ID == teacherID) || actor.HasRole(domain.AdminRole)) {
		return nil, xapperror.New(http.StatusForbidden, "user does not have rights to read this teacher's courses",
			fmt.Sprintf("no rights to read this teacher's (%s) courses with actor.ID = %s (or no %s roles)", teacherID, actor.ID, domain.AdminRole), nil)
	}

	courses, err := c.courseEdit.GetAllByTeacher(ctx, teacherID, offset, limit)
	if err != nil {
		return nil, err
	}

	return courses, nil
}

func (c CoursesService) GetAllCoursesTemplates(initialCtx context.Context, actor domain.Actor, offset, limit int) ([]*domain.Course, error) {
	_ = zapctx.Logger(initialCtx)

	tr := global.Tracer(domain.ServiceName)
	ctx, span := tr.Start(initialCtx, "courses/service.GetAllCoursesTemplates")
	defer span.End()

	ToSpan(&span, actor)

	if !actor.HasRole(domain.AdminRole) {
		return nil, xapperror.New(http.StatusForbidden, "user does not have rights to read all courses templates",
			fmt.Sprintf("no rights to read all courses templates with no %s role", domain.AdminRole), nil)
	}

	courses, err := c.courseEdit.GetAll(ctx, offset, limit)
	if err != nil {
		return nil, err
	}

	return courses, nil
}
