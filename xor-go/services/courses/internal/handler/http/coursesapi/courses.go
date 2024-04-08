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

//func (h *CoursesHandler) Get(ginCtx *gin.Context, uuid openapitypes.UUID) {
//	_, newCtx, span := getPaymentTracerSpan(c, ".Get")
//	defer span.End()
//
//	domain, err := h.coursesService.Get(newCtx, uuid)
//	if err != nil {
//		http2.AbortWithBadResponse(c, http2.MapErrorToCode(err), err)
//		return
//	}
//
//	response := DomainToGet(*domain)
//
//	c.JSON(http.StatusOK, response)
//}

func (h CoursesHandler) GetCoursesEdit(ginCtx *gin.Context, params generated.GetCoursesEditParams) {
	//	 TODO
}

func (h CoursesHandler) PostCoursesEdit(ginCtx *gin.Context, params generated.PostCoursesEditParams) {
	//TODO implement me
	panic("implement me")
}

func (h CoursesHandler) DeleteCoursesEditCourseID(ginCtx *gin.Context, courseID uuid.UUID, params generated.DeleteCoursesEditCourseIDParams) {
	//TODO implement me
	panic("implement me")
}

func (h CoursesHandler) GetCoursesEditCourseID(ginCtx *gin.Context, courseID openapitypes.UUID, params generated.GetCoursesEditCourseIDParams) {
	tr := global.Tracer(domain.ServiceName)
	ctxTrace, span := tr.Start(ginCtx, "courses/handler.GetCoursesEdit")
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
	ctxTrace, span := tr.Start(ginCtx, "courses/handler.GetCoursesEdit")
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

func (h CoursesHandler) GetCoursesCourseID(ginCtx *gin.Context, courseID uuid.UUID, params generated.GetCoursesCourseIDParams) {
	tr := global.Tracer(domain.ServiceName)
	ctxTrace, span := tr.Start(ginCtx, "courses/handler.GetCoursesEdit")
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

func (h CoursesHandler) PostCoursesCourseIDBuy(c *gin.Context, courseID openapitypes.UUID, params generated.PostCoursesCourseIDBuyParams) {
	//TODO implement me
	panic("implement me")
}

func (h CoursesHandler) PostLessonsEdit(ginCtx *gin.Context, params generated.PostLessonsEditParams) {
	//TODO implement me
	panic("implement me")
}

func (h CoursesHandler) DeleteLessonsEditLessonID(ginCtx *gin.Context, lessonID uuid.UUID, params generated.DeleteLessonsEditLessonIDParams) {
	//TODO implement me
	panic("implement me")
}

func (h CoursesHandler) GetLessonsEditLessonID(ginCtx *gin.Context, lessonID uuid.UUID, params generated.GetLessonsEditLessonIDParams) {
	//TODO implement me
	panic("implement me")
}

func (h CoursesHandler) PutLessonsEditLessonID(ginCtx *gin.Context, lessonID uuid.UUID, params generated.PutLessonsEditLessonIDParams) {
	//TODO implement me
	panic("implement me")
}

func (h CoursesHandler) GetLessonsLessonID(ginCtx *gin.Context, lessonID uuid.UUID, params generated.GetLessonsLessonIDParams) {
	//TODO implement me
	panic("implement me")
}

func (h CoursesHandler) PostLessonsLessonIDBuy(ginCtx *gin.Context, lessonID uuid.UUID, params generated.PostLessonsLessonIDBuyParams) {
	//TODO implement me
	panic("implement me")
}

func (h CoursesHandler) PostPublicationRequests(ginCtx *gin.Context, params generated.PostPublicationRequestsParams) {
	//TODO implement me
	panic("implement me")
}

func (h CoursesHandler) PutPublicationRequestsRequestID(ginCtx *gin.Context, requestID uuid.UUID, params generated.PutPublicationRequestsRequestIDParams) {
	//TODO implement me
	panic("implement me")
}

func (h CoursesHandler) PostStudentsRegister(ginCtx *gin.Context, params generated.PostStudentsRegisterParams) {
	//TODO implement me
	panic("implement me")
}

func (h CoursesHandler) PostTeachersRegister(ginCtx *gin.Context, params generated.PostTeachersRegisterParams) {
	//TODO implement me
	panic("implement me")
}

func (h CoursesHandler) PostUserAccessConfirm(ginCtx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (h CoursesHandler) PutUserAccessLessons(ginCtx *gin.Context, params generated.PutUserAccessLessonsParams) {
	//TODO implement me
	panic("implement me")
}

func (h CoursesHandler) GetUserAccessLessonsLessonID(ginCtx *gin.Context, lessonID uuid.UUID, params generated.GetUserAccessLessonsLessonIDParams) {
	//TODO implement me
	panic("implement me")
}

func NewCoursesHandler(logger *zap.Logger, coursesService adapters.CoursesService) *CoursesHandler {
	return &CoursesHandler{logger: logger, coursesService: coursesService}
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
