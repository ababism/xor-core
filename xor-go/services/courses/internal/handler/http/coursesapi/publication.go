package coursesapi

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"xor-go/services/courses/internal/handler/generated"
)

func (h CoursesHandler) PostPublicationRequests(ginCtx *gin.Context, params generated.PostPublicationRequestsParams) {
	//TODO implement me
	panic("implement me")
}

func (h CoursesHandler) PutPublicationRequestsRequestID(ginCtx *gin.Context, requestID uuid.UUID, params generated.PutPublicationRequestsRequestIDParams) {
	//TODO implement me
	panic("implement me")
}
