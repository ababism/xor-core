package coursesapi

import (
	"github.com/gin-gonic/gin"
	openapitypes "github.com/oapi-codegen/runtime/types"
	"go.uber.org/zap"
	"xor-go/services/courses/internal/handler/generated"
	"xor-go/services/courses/internal/service/adapters"
)

var _ generated.ServerInterface = &CoursesHandler{}

type CoursesHandler struct {
	logger         *zap.Logger
	coursesService adapters.CoursesService
}

func (h CoursesHandler) GetCoursesEdit(c *gin.Context, params generated.GetCoursesEditParams) {
	//TODO implement me
	panic("implement me")
}

func (h CoursesHandler) PostCoursesEdit(c *gin.Context, params generated.PostCoursesEditParams) {
	//TODO implement me
	panic("implement me")
}

func (h CoursesHandler) DeleteCoursesEditCourseID(c *gin.Context, courseID openapitypes.UUID, params generated.DeleteCoursesEditCourseIDParams) {
	//TODO implement me
	panic("implement me")
}

func (h CoursesHandler) GetCoursesEditCourseID(c *gin.Context, courseID openapitypes.UUID, params generated.GetCoursesEditCourseIDParams) {
	//TODO implement me
	panic("implement me")
}

func (h CoursesHandler) PutCoursesEditCourseID(c *gin.Context, courseID openapitypes.UUID, params generated.PutCoursesEditCourseIDParams) {
	//TODO implement me
	panic("implement me")
}

func (h CoursesHandler) GetCoursesCourseID(c *gin.Context, courseID openapitypes.UUID, params generated.GetCoursesCourseIDParams) {
	//TODO implement me
	panic("implement me")
}

func (h CoursesHandler) PostCoursesCourseIDBuy(c *gin.Context, courseID openapitypes.UUID) {
	//TODO implement me
	panic("implement me")
}

func (h CoursesHandler) PostLessonsEdit(c *gin.Context, params generated.PostLessonsEditParams) {
	//TODO implement me
	panic("implement me")
}

func (h CoursesHandler) DeleteLessonsEditLessonID(c *gin.Context, lessonID openapitypes.UUID, params generated.DeleteLessonsEditLessonIDParams) {
	//TODO implement me
	panic("implement me")
}

func (h CoursesHandler) GetLessonsEditLessonID(c *gin.Context, lessonID openapitypes.UUID, params generated.GetLessonsEditLessonIDParams) {
	//TODO implement me
	panic("implement me")
}

func (h CoursesHandler) PutLessonsEditLessonID(c *gin.Context, lessonID openapitypes.UUID, params generated.PutLessonsEditLessonIDParams) {
	//TODO implement me
	panic("implement me")
}

func (h CoursesHandler) GetLessonsLessonID(c *gin.Context, lessonID openapitypes.UUID, params generated.GetLessonsLessonIDParams) {
	//TODO implement me
	panic("implement me")
}

func (h CoursesHandler) PostLessonsLessonIDBuy(c *gin.Context, lessonID openapitypes.UUID, params generated.PostLessonsLessonIDBuyParams) {
	//TODO implement me
	panic("implement me")
}

func (h CoursesHandler) PostPublicationRequests(c *gin.Context, params generated.PostPublicationRequestsParams) {
	//TODO implement me
	panic("implement me")
}

func (h CoursesHandler) PutPublicationRequestsRequestID(c *gin.Context, requestID openapitypes.UUID, params generated.PutPublicationRequestsRequestIDParams) {
	//TODO implement me
	panic("implement me")
}

func (h CoursesHandler) PostStudentsRegister(c *gin.Context, params generated.PostStudentsRegisterParams) {
	//TODO implement me
	panic("implement me")
}

func (h CoursesHandler) PostTeachersRegister(c *gin.Context, params generated.PostTeachersRegisterParams) {
	//TODO implement me
	panic("implement me")
}

func (h CoursesHandler) PostUserAccessConfirm(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (h CoursesHandler) PutUserAccessLessons(c *gin.Context, params generated.PutUserAccessLessonsParams) {
	//TODO implement me
	panic("implement me")
}

func (h CoursesHandler) GetUserAccessLessonsLessonID(c *gin.Context, lessonID openapitypes.UUID, params generated.GetUserAccessLessonsLessonIDParams) {
	//TODO implement me
	panic("implement me")
}

func NewCoursesHandler(logger *zap.Logger, coursesService adapters.CoursesService) *CoursesHandler {
	return &CoursesHandler{logger: logger, coursesService: coursesService}
}

//func (h *CoursesHandler) GetTripByID(ginCtx *gin.Context, tripId openapitypes.UUID, params generated.GetTripByIDParams) {
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
//func (h *CoursesHandler) AcceptTrip(ginCtx *gin.Context, tripId openapitypes.UUID, params generated.AcceptTripParams) {
//	tr := global.Tracer(domain.ServiceName)
//	ctxTrace, span := tr.Start(ginCtx, "driver/driver_api.AcceptTrip")
//	defer span.End()
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
