package coursesapi

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/juju/zaputil/zapctx"
	global "go.opentelemetry.io/otel"
	"go.uber.org/zap"
	"net/http"
	"xor-go/pkg/xapperror"
	"xor-go/services/courses/internal/domain"
	"xor-go/services/courses/internal/handler/generated"
	"xor-go/services/courses/internal/handler/http/models"
)

// GetLessonsLessonID READ lesson
func (h CoursesHandler) GetLessonsLessonID(ginCtx *gin.Context, lessonID uuid.UUID, params generated.GetLessonsLessonIDParams) {
	tr := global.Tracer(domain.ServiceName)
	ctxTrace, span := tr.Start(ginCtx, "courses/handler.GetLessonsLessonID")
	defer span.End()

	ctx := zapctx.WithLogger(ctxTrace, h.logger)

	roles, err := h.coursesService.GetActorRoles(ctx, params.Actor.ToDomain())

	lesson, err := h.coursesService.ReadLesson(ctx, params.Actor.ToDomainWithRoles(roles), lessonID)
	if err != nil {
		h.abortWithAutoResponse(ginCtx, err)
		return
	}

	if lesson == nil {
		err := xapperror.New(http.StatusInternalServerError, "nil lesson without error", "get lesson returned nil lesson without error", nil)
		h.logger.Error("nil lesson", zap.Error(err))
		AbortWithBadResponse(ginCtx, h.logger, MapErrorToCode(err), err)
		return
	}
	resp := models.ToLessonResponse(*lesson)

	ginCtx.JSON(http.StatusOK, resp)
}

// PostLessonsEdit CREATE
func (h CoursesHandler) PostLessonsEdit(ginCtx *gin.Context, params generated.PostLessonsEditParams) {
	tr := global.Tracer(domain.ServiceName)
	ctxTrace, span := tr.Start(ginCtx, "courses/handler.PostLessonsEdit")
	defer span.End()

	ctx := zapctx.WithLogger(ctxTrace, h.logger)

	var payload generated.Lesson
	h.bindRequestBody(ginCtx, &payload)
	roles, err := h.coursesService.GetActorRoles(ctx, params.Actor.ToDomain())

	lesson, err := h.coursesService.CreateLesson(ctx, params.Actor.ToDomainWithRoles(roles), payload.ToDomain())
	if err != nil {
		h.abortWithAutoResponse(ginCtx, err)
		return
	}

	if lesson == nil {
		err := xapperror.New(http.StatusInternalServerError, "nil lesson without error", "create lesson returned nil lesson without error", nil)
		h.logger.Error("nil lesson", zap.Error(err))
		AbortWithBadResponse(ginCtx, h.logger, MapErrorToCode(err), err)
		return
	}
	resp := models.ToLessonResponse(*lesson)

	ginCtx.JSON(http.StatusOK, resp)
}

func (h CoursesHandler) GetLessonsEditLessonID(ginCtx *gin.Context, lessonID uuid.UUID, params generated.GetLessonsEditLessonIDParams) {
	tr := global.Tracer(domain.ServiceName)
	ctxTrace, span := tr.Start(ginCtx, "courses/handler.PostLessonsEdit")
	defer span.End()

	ctx := zapctx.WithLogger(ctxTrace, h.logger)

	roles, err := h.coursesService.GetActorRoles(ctx, params.Actor.ToDomain())

	lesson, err := h.coursesService.GetLesson(ctx, params.Actor.ToDomainWithRoles(roles), lessonID)
	if err != nil {
		h.abortWithAutoResponse(ginCtx, err)
		return
	}

	if lesson == nil {
		err := xapperror.New(http.StatusInternalServerError, "nil lesson without error", "get lesson returned nil lesson without error", nil)
		h.logger.Error("nil lesson", zap.Error(err))
		AbortWithBadResponse(ginCtx, h.logger, MapErrorToCode(err), err)
		return
	}
	resp := models.ToLessonResponse(*lesson)

	ginCtx.JSON(http.StatusOK, resp)
}

// PutLessonsEditLessonID UPDATE
func (h CoursesHandler) PutLessonsEditLessonID(ginCtx *gin.Context, lessonID uuid.UUID, params generated.PutLessonsEditLessonIDParams) {
	tr := global.Tracer(domain.ServiceName)
	ctxTrace, span := tr.Start(ginCtx, "courses/handler.PutLessonsEditLessonID")
	defer span.End()

	ctx := zapctx.WithLogger(ctxTrace, h.logger)

	var payload generated.Lesson
	h.bindRequestBody(ginCtx, &payload)

	roles := make([]string, 0)
	roles, err := h.coursesService.GetActorRoles(ctx, params.Actor.ToDomain())

	lesson, err := h.coursesService.UpdateLesson(ctx, params.Actor.ToDomainWithRoles(roles), lessonID, payload.ToDomain())
	if err != nil {
		h.abortWithAutoResponse(ginCtx, err)
		return
	}

	if lesson == nil {
		err := xapperror.New(http.StatusInternalServerError, "nil lesson without error", "update lesson returned nil lesson without error", nil)
		h.logger.Error("nil lesson", zap.Error(err))
		AbortWithBadResponse(ginCtx, h.logger, MapErrorToCode(err), err)
		return
	}
	resp := models.ToLessonResponse(*lesson)

	ginCtx.JSON(http.StatusOK, resp)
}

func (h CoursesHandler) DeleteLessonsEditLessonID(ginCtx *gin.Context, lessonID uuid.UUID, params generated.DeleteLessonsEditLessonIDParams) {
	tr := global.Tracer(domain.ServiceName)
	ctxTrace, span := tr.Start(ginCtx, "courses/handler.DeleteLessonsEditLessonID")
	defer span.End()

	ctx := zapctx.WithLogger(ctxTrace, h.logger)

	roles, err := h.coursesService.GetActorRoles(ctx, params.Actor.ToDomain())

	err = h.coursesService.DeleteLesson(ctx, params.Actor.ToDomainWithRoles(roles), lessonID)
	if err != nil {
		h.abortWithAutoResponse(ginCtx, err)
		return
	}

	ginCtx.JSON(http.StatusOK, http.NoBody)
}

// PostLessonsLessonIDBuy Buy lesson
func (h CoursesHandler) PostLessonsLessonIDBuy(ginCtx *gin.Context, lessonID uuid.UUID, params generated.PostLessonsLessonIDBuyParams) {
	tr := global.Tracer(domain.ServiceName)
	ctxTrace, span := tr.Start(ginCtx, "courses/handler.PostCoursesCourseIDBuy")
	defer span.End()

	ctx := zapctx.WithLogger(ctxTrace, h.logger)

	roles, err := h.coursesService.GetActorRoles(ctx, params.Actor.ToDomain())

	redirect, lAccess, err := h.coursesService.BuyLesson(ctx, params.Actor.ToDomainWithRoles(roles), lessonID)
	if err != nil {
		h.abortWithAutoResponse(ginCtx, err)
		return
	}

	if redirect.Response == "" {
		err := xapperror.New(http.StatusInternalServerError, "nil redirect without error", "GetCourse returned nil redirect without error", nil)
		h.logger.Error("nil redirect", zap.Error(err))
		AbortWithBadResponse(ginCtx, h.logger, MapErrorToCode(err), err)
		return
	}
	respRedirect := models.ToPaymentRedirectResponse(redirect)
	respLessonAccess := models.ToLessonAccessResponse(lAccess)

	ginCtx.JSON(http.StatusOK, struct {
		generated.PaymentRedirect
		generated.LessonAccess
	}{respRedirect, respLessonAccess})
}
