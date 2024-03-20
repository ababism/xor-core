package bank_account_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"xor-go/pkg/xerror"
	"xor-go/services/finances/internal/handler/http/dto_models"
)

func AbortWithBadResponse(c *gin.Context, logger *zap.Logger, statusCode int, err error) {
	logger.Debug(fmt.Sprintf("%s: %d %s", c.Request.URL, statusCode, xerror.GetLastMessage(err)))
	c.AbortWithStatusJSON(statusCode, dto_models.Error{Message: xerror.GetLastMessage(err)})
}

func AbortWithErrorResponse(c *gin.Context, logger *zap.Logger, statusCode int, message string) {
	logger.Error(fmt.Sprintf("%s: %d %s", c.Request.URL, statusCode, message))
	c.AbortWithStatusJSON(statusCode, dto_models.Error{Message: message})
}

func MapErrorToCode(err error) int {
	return xerror.GetCode(err)
}
