package driver_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gitlab.com/ArtemFed/mts-final-taxi/pkg/apperror"
	"gitlab.com/ArtemFed/mts-final-taxi/projects/template/internal/handler/http/models"
	"go.uber.org/zap"
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
