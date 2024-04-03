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
	coursesService adapters.CourseService
}

func NewDriverHandler(logger *zap.Logger, coursesService adapters.CourseService) *CoursesHandler {
	return &CoursesHandler{logger: logger, coursesService: coursesService}
}

func (h *CoursesHandler) PostBuyCourse(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (h *CoursesHandler) GetCourses(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (h *CoursesHandler) GetCoursesCourseId(c *gin.Context, courseId openapitypes.UUID) {
	//TODO implement me
	panic("implement me")
}

func (h *CoursesHandler) GetCoursesCourseIdLessons(c *gin.Context, courseId openapitypes.UUID) {
	//TODO implement me
	panic("implement me")
}

func (h *CoursesHandler) GetCoursesCourseIdLessonsLessonId(c *gin.Context, courseId openapitypes.UUID, lessonId openapitypes.UUID) {
	//TODO implement me
	panic("implement me")
}

func (h *CoursesHandler) GetTeachers(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (h *CoursesHandler) PostTeachers(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (h *CoursesHandler) DeleteTeachersTeacherID(c *gin.Context, teacherID openapitypes.UUID) {
	//TODO implement me
	panic("implement me")
}

func (h *CoursesHandler) GetTeachersTeacherID(c *gin.Context, teacherID openapitypes.UUID) {
	//TODO implement me
	panic("implement me")
}

func (h *CoursesHandler) PutTeachersTeacherID(c *gin.Context, teacherID openapitypes.UUID) {
	//TODO implement me
	panic("implement me")
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
