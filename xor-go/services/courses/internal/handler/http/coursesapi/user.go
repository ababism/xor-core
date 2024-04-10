package coursesapi

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/juju/zaputil/zapctx"
	openapitypes "github.com/oapi-codegen/runtime/types"
	global "go.opentelemetry.io/otel"
	"net/http"
	"xor-go/services/courses/internal/domain"
	"xor-go/services/courses/internal/handler/generated"
	"xor-go/services/courses/internal/handler/http/models"
)

func (h CoursesHandler) PostStudentsRegister(ginCtx *gin.Context, params generated.PostStudentsRegisterParams) {
	tr := global.Tracer(domain.ServiceName)
	ctxTrace, span := tr.Start(ginCtx, "courses/handler.PostStudentsRegister")
	defer span.End()

	ctx := zapctx.WithLogger(ctxTrace, h.logger)

	var payload generated.Student
	h.bindRequestBody(ginCtx, &payload)

	err := h.coursesService.RegisterStudentProfile(ctx, params.Actor.ToDomain(), payload.ToDomain())
	if err != nil {
		h.abortWithAutoResponse(ginCtx, err)
		return
	}

	ginCtx.JSON(http.StatusOK, http.NoBody)
}

func (h CoursesHandler) PostTeachersRegister(ginCtx *gin.Context, params generated.PostTeachersRegisterParams) {
	tr := global.Tracer(domain.ServiceName)
	ctxTrace, span := tr.Start(ginCtx, "courses/handler.PostTeachersRegister")
	defer span.End()

	ctx := zapctx.WithLogger(ctxTrace, h.logger)

	var payload generated.Teacher
	h.bindRequestBody(ginCtx, &payload)

	err := h.coursesService.RegisterTeacherProfile(ctx, params.Actor.ToDomain(), payload.ToDomain())
	if err != nil {
		h.abortWithAutoResponse(ginCtx, err)
		return
	}

	ginCtx.JSON(http.StatusOK, http.NoBody)
}

func (h CoursesHandler) PostUserAccessConfirmBuyerID(ginCtx *gin.Context, buyerID openapitypes.UUID) {
	tr := global.Tracer(domain.ServiceName)
	ctxTrace, span := tr.Start(ginCtx, "courses/handler.PostTeachersRegister")
	defer span.End()

	ctx := zapctx.WithLogger(ctxTrace, h.logger)

	var payload []generated.Product
	h.bindRequestBody(ginCtx, &payload)

	err := h.coursesService.ConfirmAccess(ctx, buyerID, generated.ToProductSliceDomain(payload))
	if err != nil {
		h.abortWithAutoResponse(ginCtx, err)
		return
	}

	ginCtx.JSON(http.StatusOK, http.NoBody)
}

func (h CoursesHandler) PutUserAccessLessons(ginCtx *gin.Context, params generated.PutUserAccessLessonsParams) {
	tr := global.Tracer(domain.ServiceName)
	ctxTrace, span := tr.Start(ginCtx, "courses/handler.PutUserAccessLessons")
	defer span.End()

	ctx := zapctx.WithLogger(ctxTrace, h.logger)

	var payload generated.LessonAccess
	h.bindRequestBody(ginCtx, &payload)

	lessonAccess, err := h.coursesService.CreateOrChangeLessonAccess(ctx, params.Actor.ToDomain(), payload.ToDomain())
	if err != nil {
		h.abortWithAutoResponse(ginCtx, err)
		return
	}

	resp := models.ToLessonAccessResponse(lessonAccess)

	ginCtx.JSON(http.StatusOK, resp)
}

func (h CoursesHandler) GetUserAccessLessonsLessonID(ginCtx *gin.Context, lessonID uuid.UUID, params generated.GetUserAccessLessonsLessonIDParams) {
	tr := global.Tracer(domain.ServiceName)
	ctxTrace, span := tr.Start(ginCtx, "courses/handler.GetUserAccessLessonsLessonID")
	defer span.End()

	ctx := zapctx.WithLogger(ctxTrace, h.logger)

	lessonAccess, err := h.coursesService.GetLessonAccess(ctx, params.Actor.ToDomain(), lessonID)
	if err != nil {
		h.abortWithAutoResponse(ginCtx, err)
		return
	}

	resp := models.ToLessonAccessResponse(lessonAccess)

	ginCtx.JSON(http.StatusOK, resp)
}
