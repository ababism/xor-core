package xor_error

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"xor-go/pkg/xor_http/response"
)

func HandleInternalError(ctx *gin.Context, responser *response.HttpResponseWrapper, err error) {
	handleInternalError(ctx, err, responser.HandleError)
}

func HandleInternalErrorWithMessage(ctx *gin.Context, responser *response.HttpResponseWrapper, err error) {
	handleInternalError(ctx, err, responser.HandleErrorWithMessage)
}

var (
	illegalArgumentError *IllegalArgumentError
)

type errorHandler func(ctx *gin.Context, code int, err error)

func handleInternalError(ctx *gin.Context, err error, handler errorHandler) {
	switch {
	case errors.As(err, &illegalArgumentError):
		handler(ctx, http.StatusBadRequest, err)
	default:
		handler(ctx, http.StatusInternalServerError, err)
	}
}
