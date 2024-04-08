package coursesapi

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"xor-go/services/courses/internal/handler/generated"
)

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
