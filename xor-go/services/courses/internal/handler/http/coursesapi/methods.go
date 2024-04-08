package coursesapi

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"xor-go/pkg/xapperror"
	"xor-go/services/courses/internal/handler/http/models"
)

func AbortWithBadResponse(c *gin.Context, logger *zap.Logger, statusCode int, err error) {
	logger.Debug(fmt.Sprintf("%s: %d %s", c.Request.URL, statusCode, xapperror.GetLastMessage(err)))
	c.AbortWithStatusJSON(statusCode, models.Error{Message: xapperror.GetLastMessage(err)})
}

func AbortWithErrorResponse(c *gin.Context, logger *zap.Logger, statusCode int, message string) {
	logger.Error(fmt.Sprintf("%s: %d %s", c.Request.URL, statusCode, message))
	c.AbortWithStatusJSON(statusCode, models.Error{Message: message})
}

func MapErrorToCode(err error) int {
	return xapperror.GetCode(err)
}

func (h CoursesHandler) abortWithBadResponse(c *gin.Context, statusCode int, err error) {
	h.logger.Debug(fmt.Sprintf("%s: %d %s", c.Request.URL, statusCode, xapperror.GetLastMessage(err)))
	c.AbortWithStatusJSON(statusCode, models.Error{Message: xapperror.GetLastMessage(err)})
}

func (h CoursesHandler) abortWithAutoResponse(c *gin.Context, err error) {
	h.logger.Debug(fmt.Sprintf("%s: %d %s", c.Request.URL, xapperror.GetCode(err), xapperror.GetLastMessage(err)))
	c.AbortWithStatusJSON(xapperror.GetCode(err), models.Error{Message: xapperror.GetLastMessage(err)})
}

func (h CoursesHandler) bindRequestBody(c *gin.Context, obj any) bool {
	if err := c.BindJSON(obj); err != nil {
		AbortWithBadResponse(c, h.logger, http.StatusBadRequest, err)
		return false
	}
	return true
}
