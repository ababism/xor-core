package coursesapi

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/juju/zaputil/zapctx"
	openapitypes "github.com/oapi-codegen/runtime/types"
	global "go.opentelemetry.io/otel"
	"go.uber.org/zap"
	"net/http"
	"xor-go/pkg/apperror"
	"xor-go/services/courses/internal/domain"
	"xor-go/services/courses/internal/handler/generated"
	"xor-go/services/courses/internal/handler/http/models"
	"xor-go/services/courses/internal/service/adapters"
)

var _ generated.ServerInterface = &CoursesHandler{}

type CoursesHandler struct {
	logger         *zap.Logger
	coursesService adapters.CoursesService
}

func NewCoursesHandler(logger *zap.Logger, coursesService adapters.CoursesService) *CoursesHandler {
	return &CoursesHandler{logger: logger, coursesService: coursesService}
}

// GetCoursesEdit LIST COURSES
func (h CoursesHandler) GetCoursesEdit(ginCtx *gin.Context, params generated.GetCoursesEditParams) {
	//	 TODO
	panic("implement me")
}

// GetCoursesCourseID READ Published
func (h CoursesHandler) GetCoursesCourseID(ginCtx *gin.Context, courseID uuid.UUID, params generated.GetCoursesCourseIDParams) {
	tr := global.Tracer(domain.ServiceName)
	ctxTrace, span := tr.Start(ginCtx, "courses/handler.GetCoursesCourseID")
	defer span.End()

	ctx := zapctx.WithLogger(ctxTrace, h.logger)

	course, err := h.coursesService.ReadCourse(ctx, params.Actor.ToDomain(), courseID)
	if err != nil {
		h.abortWithAutoResponse(ginCtx, err)
		return
	}
	if course == nil {
		err := apperror.New(http.StatusInternalServerError, "nil course without error", "GetCourse returned nil course without error", nil)
		h.logger.Error("nil course", zap.Error(err))
		AbortWithBadResponse(ginCtx, h.logger, MapErrorToCode(err), err)
	}
	resp := models.ToCourseResponse(*course)

	ginCtx.JSON(http.StatusOK, resp)
}

// PostCoursesEdit CREATE course
func (h CoursesHandler) PostCoursesEdit(ginCtx *gin.Context, params generated.PostCoursesEditParams) {
	tr := global.Tracer(domain.ServiceName)
	ctxTrace, span := tr.Start(ginCtx, "courses/handler.PostCoursesEdit")
	defer span.End()

	ctx := zapctx.WithLogger(ctxTrace, h.logger)

	var coursePayload generated.Course
	h.bindRequestBody(ginCtx, &coursePayload)

	course, err := h.coursesService.CreateCourse(ctx, params.Actor.ToDomain(), coursePayload.ToDomain())
	if err != nil {
		h.abortWithAutoResponse(ginCtx, err)
		return
	}

	if course == nil {
		err := apperror.New(http.StatusInternalServerError, "nil course without error", "GetCourse returned nil course without error", nil)
		h.logger.Error("nil course", zap.Error(err))
		AbortWithBadResponse(ginCtx, h.logger, MapErrorToCode(err), err)
	}
	resp := models.ToCourseResponse(*course)

	ginCtx.JSON(http.StatusOK, resp)
}

// GetCoursesEditCourseID UPDATE course
func (h CoursesHandler) GetCoursesEditCourseID(ginCtx *gin.Context, courseID openapitypes.UUID, params generated.GetCoursesEditCourseIDParams) {
	tr := global.Tracer(domain.ServiceName)
	ctxTrace, span := tr.Start(ginCtx, "courses/handler.GetCoursesEditCourseID")
	defer span.End()

	ctx := zapctx.WithLogger(ctxTrace, h.logger)

	course, err := h.coursesService.GetCourse(ctx, params.Actor.ToDomain(), courseID)
	if err != nil {
		h.abortWithAutoResponse(ginCtx, err)
		return
	}
	if course == nil {
		err := apperror.New(http.StatusInternalServerError, "nil course without error", "GetCourse returned nil course without error", nil)
		h.logger.Error("nil course", zap.Error(err))
		AbortWithBadResponse(ginCtx, h.logger, MapErrorToCode(err), err)
	}
	resp := models.ToCourseResponse(*course)

	ginCtx.JSON(http.StatusOK, resp)
}
func (h CoursesHandler) PutCoursesEditCourseID(ginCtx *gin.Context, courseID uuid.UUID, params generated.PutCoursesEditCourseIDParams) {
	tr := global.Tracer(domain.ServiceName)
	ctxTrace, span := tr.Start(ginCtx, "courses/handler.PutCoursesEditCourseID")
	defer span.End()

	ctx := zapctx.WithLogger(ctxTrace, h.logger)

	var coursePayload generated.Course
	h.bindRequestBody(ginCtx, &coursePayload)

	course, err := h.coursesService.UpdateCourse(ctx, params.Actor.ToDomain(), courseID, coursePayload.ToDomain())
	if err != nil {
		h.abortWithAutoResponse(ginCtx, err)
		return
	}

	if course == nil {
		err := apperror.New(http.StatusInternalServerError, "nil course without error", "GetCourse returned nil course without error", nil)
		h.logger.Error("nil course", zap.Error(err))
		AbortWithBadResponse(ginCtx, h.logger, MapErrorToCode(err), err)
	}
	resp := models.ToCourseResponse(*course)

	ginCtx.JSON(http.StatusOK, resp)
}

func (h CoursesHandler) DeleteCoursesEditCourseID(ginCtx *gin.Context, courseID uuid.UUID, params generated.DeleteCoursesEditCourseIDParams) {
	tr := global.Tracer(domain.ServiceName)
	ctxTrace, span := tr.Start(ginCtx, "courses/handler.DeleteCoursesEditCourseID")
	defer span.End()

	ctx := zapctx.WithLogger(ctxTrace, h.logger)

	err := h.coursesService.DeleteCourse(ctx, params.Actor.ToDomain(), courseID)
	if err != nil {
		h.abortWithAutoResponse(ginCtx, err)
		return
	}

	ginCtx.JSON(http.StatusNoContent, http.NoBody)
}

// PostCoursesCourseIDBuy Buy lessons from course
func (h CoursesHandler) PostCoursesCourseIDBuy(ginCtx *gin.Context, courseID openapitypes.UUID, params generated.PostCoursesCourseIDBuyParams) {
	tr := global.Tracer(domain.ServiceName)
	ctxTrace, span := tr.Start(ginCtx, "courses/handler.PostCoursesCourseIDBuy")
	defer span.End()

	ctx := zapctx.WithLogger(ctxTrace, h.logger)

	redirect, err := h.coursesService.BuyCourse(ctx, params.Actor.ToDomain(), courseID)
	if err != nil {
		h.abortWithAutoResponse(ginCtx, err)
		return
	}

	if redirect.Response == "" {
		err := apperror.New(http.StatusInternalServerError, "nil redirect without error", "GetCourse returned nil redirect without error", nil)
		h.logger.Error("nil redirect", zap.Error(err))
		AbortWithBadResponse(ginCtx, h.logger, MapErrorToCode(err), err)
	}
	resp := models.ToPaymentRedirectResponse(redirect)

	ginCtx.JSON(http.StatusOK, resp)
}

//func (h *CoursesHandler) GetTripByID(ginCtx *gin.Context, tripId uuid.UUID, params generated.GetTripByIDParams) {
//	tr := global.Tracer(domain.ServiceName)
//	ctxTrace, span := tr.Start(ginCtx, "driver/driver_api.GetTripByID")
//	defer span.End()
//
//	ctx := zapctx.WithLogger(ctxTrace, h.logger)
//
//	trip, err := h.coursesService.GetTripByID(ctx, params.UserId, tripId)
//	if err != nil {
//		AbortWithBadResponse(ginCtx, h.logger, MapErrorToCode(err), err)
//		return
//	}
//	resp := models.ToTripResponse(*trip)
//
//	ginCtx.JSON(http.StatusOK, resp)
//}
//
//func (h *CoursesHandler) AcceptTrip(ginCtx *gin.Context, tripId uuid.UUID, params generated.AcceptTripParams) {
//	tr := global.Tracer(domain.ServiceName)
//	ctxTrace, span := tr.Start(ginCtx, "driver/driver_api.AcceptTrip")
//	defer span.End()s
//
//	ctx := zapctx.WithLogger(ctxTrace, h.logger)
//
//	err := h.coursesService.AcceptTrip(ctx, params.UserId, tripId)
//	if err != nil {
//		AbortWithBadResponse(ginCtx, h.logger, MapErrorToCode(err), err)
//		return
//	}
//
//	ginCtx.JSON(http.StatusOK, http.NoBody)
//}
//
