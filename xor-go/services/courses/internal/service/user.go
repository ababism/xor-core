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
	"time"
	"xor-go/pkg/xapperror"
	"xor-go/services/courses/internal/domain"
	"xor-go/services/courses/internal/domain/keys"
)

func (c CoursesService) GetActorRoles(ctx context.Context, actor domain.Actor) ([]string, error) {
	_ = zapctx.Logger(ctx)

	tr := global.Tracer(domain.ServiceName)
	ctx, span := tr.Start(ctx, "courses/service.GetActorRoles")
	defer span.End()

	ToSpan(&span, actor)

	roles := make([]string, 0)
	_, err := c.student.Get(ctx, actor.ID)
	if err == nil {
		roles = append(roles, domain.StudentRole)
	}

	_, err = c.teacher.Get(ctx, actor.ID)
	if err == nil {
		roles = append(roles, domain.TeacherRole)
	}

	return roles, nil

}
func (c CoursesService) BuyCourse(initialCtx context.Context, actor domain.Actor, courseID uuid.UUID) (domain.PaymentRedirect, error) {
	_ = zapctx.Logger(initialCtx)

	tr := global.Tracer(domain.ServiceName)
	ctx, span := tr.Start(initialCtx, "courses/service.BuyCourse")
	defer span.End()

	ToSpan(&span, actor)

	if actor.HasOneOfRoles(domain.TeacherRole, domain.AdminRole) {
		return domain.PaymentRedirect{}, xapperror.New(http.StatusForbidden, "teacher does not have rights to buy course",
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
	stringProducts := keys.ProductToStrings(productsToBuy)
	span.AddEvent("products to buy", trace.WithAttributes(attribute.StringSlice(keys.ProductSliceAttributeKey, stringProducts)))
	redirect, err := c.financesClient.CreatePurchase(ctx, productsToBuy)
	if err != nil {
		return domain.PaymentRedirect{}, xapperror.New(http.StatusForbidden, "no available lessons to buy in course",
			"no available lessons to buy in course", nil)
	}
	span.AddEvent("purchase created", trace.WithAttributes(attribute.StringSlice(keys.ProductSliceAttributeKey, stringProducts)))
	return redirect, nil
}

func (c CoursesService) BuyLesson(initialCtx context.Context, actor domain.Actor, lessonID uuid.UUID) (domain.PaymentRedirect, domain.LessonAccess, error) {
	_ = zapctx.Logger(initialCtx)

	tr := global.Tracer(domain.ServiceName)
	ctx, span := tr.Start(initialCtx, "courses/service.BuyLesson")
	defer span.End()

	ToSpan(&span, actor)

	if actor.HasOneOfRoles(domain.TeacherRole, domain.AdminRole) {
		return domain.PaymentRedirect{}, domain.LessonAccess{}, xapperror.New(http.StatusForbidden, "teacher does not have rights to buy course",
			fmt.Sprintf("%s or %s roles can't buy course", domain.TeacherRole, domain.AdminRole), nil)
	}

	access, err := c.student.GetLessonAccess(ctx, actor.ID, lessonID)
	if err != nil {
		return domain.PaymentRedirect{}, domain.LessonAccess{}, err
	}
	if access.AccessStatus != domain.Inaccessible {
		return domain.PaymentRedirect{}, access, xapperror.New(http.StatusForbidden, "lesson already purchased or pending",
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
	span.AddEvent("purchase created", trace.WithAttributes(attribute.String(keys.ProductIDAttributeKey, lesson.Product.ID.String())))
	return redirect, access, nil
}

func (c CoursesService) RegisterStudentProfile(initialCtx context.Context, actor domain.Actor, profile domain.Student) error {
	_ = zapctx.Logger(initialCtx)

	tr := global.Tracer(domain.ServiceName)
	ctx, span := tr.Start(initialCtx, "courses/service.RegisterStudentProfile")
	defer span.End()

	ToSpan(&span, actor)

	if !actor.HasRole(domain.UnregisteredRole) {
		return xapperror.New(http.StatusForbidden, "user can't be registered",
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

	span.AddEvent("student created", trace.WithAttributes(attribute.String(keys.StudentIDAttributeKey, profile.AccountID.String())))
	return nil
}

func (c CoursesService) RegisterTeacherProfile(initialCtx context.Context, actor domain.Actor, profile domain.Teacher) error {
	_ = zapctx.Logger(initialCtx)

	tr := global.Tracer(domain.ServiceName)
	ctx, span := tr.Start(initialCtx, "courses/service.RegisterTeacherProfile")
	defer span.End()

	ToSpan(&span, actor)

	if !actor.HasRole(domain.AdminRole) {
		return xapperror.New(http.StatusForbidden, "user can't register teachers",
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
	span.AddEvent("teacher created", trace.WithAttributes(attribute.String(keys.TeacherIDAttributeKey, profile.AccountID.String())))
	return nil
}

func (c CoursesService) CreateOrChangeLessonAccess(initialCtx context.Context, actor domain.Actor, access domain.LessonAccess) (domain.LessonAccess, error) {
	_ = zapctx.Logger(initialCtx)

	tr := global.Tracer(domain.ServiceName)
	ctx, span := tr.Start(initialCtx, "courses/service.CreateOrChangeLessonAccess")
	defer span.End()

	ToSpan(&span, actor)

	if !actor.HasRole(domain.AdminRole) {
		return domain.LessonAccess{}, xapperror.New(http.StatusForbidden, "user can't give access to lessons",
			fmt.Sprintf("user can't give acces to lessons no %s role", domain.AdminRole), nil)
	}

	curAccess, err := c.student.GetLessonAccess(ctx, access.StudentID, access.LessonID)
	// if access exists
	if err == nil {
		access.ID = curAccess.ID
		updatedAccess, err := c.student.UpdateAccessToLesson(ctx, access)
		if err != nil {
			return domain.LessonAccess{}, err
		}
		return updatedAccess, nil
	}
	// if access does not exist
	if access.ID == uuid.Nil || (access.ID == uuid.UUID{}) {
		access.ID = uuid.New()
	}

	newAccess, err := c.student.CreateAccessToLesson(ctx, access)
	if err != nil {
		return domain.LessonAccess{}, err
	}

	span.AddEvent("lesson access created", trace.WithAttributes(attribute.String(keys.LessonAccessIDAttributeKey, newAccess.ID.String())))
	return newAccess, nil
}

func (c CoursesService) GetLessonAccess(initialCtx context.Context, actor domain.Actor, lessonID uuid.UUID) (domain.LessonAccess, error) {
	_ = zapctx.Logger(initialCtx)

	tr := global.Tracer(domain.ServiceName)
	ctx, span := tr.Start(initialCtx, "courses/service.GetLessonAccess")
	defer span.End()

	ToSpan(&span, actor)

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

	span.AddEvent("products to confirm", trace.WithAttributes(attribute.StringSlice(keys.ProductSliceAttributeKey, keys.ProductToStrings(products))))

	for _, pr := range products {
		lessonAccess := domain.LessonAccess{
			ID:           uuid.Nil,
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
