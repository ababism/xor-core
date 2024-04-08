package service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/juju/zaputil/zapctx"
	global "go.opentelemetry.io/otel"
	"net/http"
	"time"
	"xor-go/pkg/apperror"
	"xor-go/services/courses/internal/domain"
)

func (c CoursesService) BuyCourse(initialCtx context.Context, actor domain.Actor, courseID uuid.UUID) (domain.PaymentRedirect, error) {
	_ = zapctx.Logger(initialCtx)

	tr := global.Tracer(domain.ServiceName)
	ctx, span := tr.Start(initialCtx, "courses/service.BuyCourse")
	defer span.End()

	if actor.HasOneOfRoles(domain.TeacherRole, domain.AdminRole) {
		return domain.PaymentRedirect{}, apperror.New(http.StatusForbidden, "teacher does not have rights to buy course",
			fmt.Sprintf("%s or %s roles can't buy course", domain.TeacherRole, domain.AdminRole), nil)
	}

	lessons, err := c.lesson.GetAllByCourse(ctx, courseID)
	if err != nil {
		return domain.PaymentRedirect{}, err
	}

	productsToBuy := make([]domain.Product, 0)
	for _, l := range lessons {
		access, err := c.student.GetLessonAccess(ctx, actor.ID, l.ID)
		if err != nil {
			return domain.PaymentRedirect{}, err
		}
		if access.AccessStatus == domain.Inaccessible {
			productsToBuy = append(productsToBuy, l.Product)
		}
	}
	if len(productsToBuy) == 0 {
		return domain.PaymentRedirect{}, err
	}
	redirect, err := c.financesClient.CreatePurchase(ctx, productsToBuy)
	if err != nil {
		return domain.PaymentRedirect{}, apperror.New(http.StatusForbidden, "no available lessons to buy in course",
			"no available lessons to buy in course", nil)
	}

	return redirect, nil
}

func (c CoursesService) BuyLesson(initialCtx context.Context, actor domain.Actor, lessonID uuid.UUID) (domain.PaymentRedirect, domain.LessonAccess, error) {
	_ = zapctx.Logger(initialCtx)

	tr := global.Tracer(domain.ServiceName)
	ctx, span := tr.Start(initialCtx, "courses/service.BuyLesson")
	defer span.End()

	if actor.HasOneOfRoles(domain.TeacherRole, domain.AdminRole) {
		return domain.PaymentRedirect{}, domain.LessonAccess{}, apperror.New(http.StatusForbidden, "teacher does not have rights to buy course",
			fmt.Sprintf("%s or %s roles can't buy course", domain.TeacherRole, domain.AdminRole), nil)
	}

	access, err := c.student.GetLessonAccess(ctx, actor.ID, lessonID)
	if err != nil {
		return domain.PaymentRedirect{}, domain.LessonAccess{}, err
	}
	if access.AccessStatus != domain.Inaccessible {
		return domain.PaymentRedirect{}, access, apperror.New(http.StatusForbidden, "lesson already purchased or pending",
			"lesson already purchased", nil)

	}

	lesson, err := c.lesson.Get(ctx, lessonID)
	if err != nil {
		return domain.PaymentRedirect{}, domain.LessonAccess{}, err
	}

	payload := []domain.Product{lesson.Product}
	redirect, err := c.financesClient.CreatePurchase(ctx, payload)
	if err != nil {
		return domain.PaymentRedirect{}, domain.LessonAccess{}, err
	}

	return redirect, access, nil
}

func (c CoursesService) RegisterStudentProfile(initialCtx context.Context, actor domain.Actor, profile domain.Student) error {
	_ = zapctx.Logger(initialCtx)

	tr := global.Tracer(domain.ServiceName)
	ctx, span := tr.Start(initialCtx, "courses/service.RegisterStudentProfile")
	defer span.End()

	if !actor.HasRole(domain.UnregisteredRole) {
		return apperror.New(http.StatusForbidden, "user can't be registered",
			fmt.Sprintf("user can't be registere user do not have %s role", domain.UnregisteredRole), nil)
	}

	err := profile.Validate()
	if err != nil {
		return err
	}

	err = c.student.Create(ctx, profile)
	if err != nil {
		return err
	}

	return nil
}

func (c CoursesService) RegisterTeacherProfile(initialCtx context.Context, actor domain.Actor, profile domain.Teacher) error {
	_ = zapctx.Logger(initialCtx)

	tr := global.Tracer(domain.ServiceName)
	ctx, span := tr.Start(initialCtx, "courses/service.RegisterTeacherProfile")
	defer span.End()

	if !actor.HasRole(domain.AdminRole) {
		return apperror.New(http.StatusForbidden, "user can't register teachers",
			fmt.Sprintf("user can't register teachers no %s role", domain.AdminRole), nil)
	}

	err := profile.Validate()
	if err != nil {
		return err
	}

	err = c.teacher.Create(ctx, profile)
	if err != nil {
		return err
	}

	return nil
}

func (c CoursesService) ChangeLessonAccess(initialCtx context.Context, actor domain.Actor, access domain.LessonAccess) (domain.LessonAccess, error) {
	_ = zapctx.Logger(initialCtx)

	tr := global.Tracer(domain.ServiceName)
	ctx, span := tr.Start(initialCtx, "courses/service.ChangeLessonAccess")
	defer span.End()

	if !actor.HasRole(domain.AdminRole) {
		return domain.LessonAccess{}, apperror.New(http.StatusForbidden, "user can't give access to lessons",
			fmt.Sprintf("user can't give acces to lessons no %s role", domain.AdminRole), nil)
	}

	newAccess, err := c.student.CreateAccessToLesson(ctx, access)
	if err != nil {
		return domain.LessonAccess{}, err
	}

	return newAccess, nil
}

func (c CoursesService) GetLessonAccess(initialCtx context.Context, actor domain.Actor, lessonID uuid.UUID) (domain.LessonAccess, error) {
	_ = zapctx.Logger(initialCtx)

	tr := global.Tracer(domain.ServiceName)
	ctx, span := tr.Start(initialCtx, "courses/service.GetLessonAccess")
	defer span.End()

	access, err := c.student.GetLessonAccess(ctx, actor.ID, lessonID)
	if err != nil {
		return domain.LessonAccess{}, err
	}

	return access, nil
}

func (c CoursesService) ConfirmAccess(initialCtx context.Context, buyerID uuid.UUID, products []domain.Product) error {
	_ = zapctx.Logger(initialCtx)

	tr := global.Tracer(domain.ServiceName)
	ctx, span := tr.Start(initialCtx, "courses/service.ConfirmAccess")
	defer span.End()

	for _, pr := range products {
		lessonAccess := domain.LessonAccess{
			ID:           uuid.UUID{},
			LessonID:     pr.Item,
			StudentID:    buyerID,
			AccessStatus: domain.Accessible,
			UpdatedAt:    time.Now(),
		}
		_, err := c.student.CreateAccessToLesson(ctx, lessonAccess)
		if err != nil {
			return err
		}
	}

	return nil
}
