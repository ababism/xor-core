package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"xor-go/pkg/xerror"
	"xor-go/services/finances/internal/handler/http/dto"
	"xor-go/services/finances/internal/log"
)

func AbortWithBadResponse(c *gin.Context, statusCode int, err error) {
	log.Logger.Debug(fmt.Sprintf("%s: %d %s", c.Request.URL, statusCode, xerror.GetLastMessage(err)))
	c.AbortWithStatusJSON(statusCode, dto.Error{Message: xerror.GetLastMessage(err)})
}

func AbortWithErrorResponse(c *gin.Context, statusCode int, message string) {
	log.Logger.Error(fmt.Sprintf("%s: %d %s", c.Request.URL, statusCode, message))
	c.AbortWithStatusJSON(statusCode, dto.Error{Message: message})
}

func MapErrorToCode(err error) int {
	return xerror.GetCode(err)
}
