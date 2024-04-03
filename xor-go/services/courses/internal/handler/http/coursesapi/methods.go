package coursesapi

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"xor-go/pkg/apperror"
	"xor-go/services/courses/internal/handler/http/models"
)

func AbortWithBadResponse(c *gin.Context, logger *zap.Logger, statusCode int, err error) {
	logger.Debug(fmt.Sprintf("%s: %d %s", c.Request.URL, statusCode, apperror.GetLastMessage(err)))
	c.AbortWithStatusJSON(statusCode, models.Error{Message: apperror.GetLastMessage(err)})
}

func AbortWithErrorResponse(c *gin.Context, logger *zap.Logger, statusCode int, message string) {
	logger.Error(fmt.Sprintf("%s: %d %s", c.Request.URL, statusCode, message))
	c.AbortWithStatusJSON(statusCode, models.Error{Message: message})
}

func MapErrorToCode(err error) int {
	return apperror.GetCode(err)
}
