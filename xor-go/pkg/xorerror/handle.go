package xorerror

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	xorhttp "xor-go/pkg/xorhttp/response"
)

func HandleInternalError(ctx *gin.Context, responser *xorhttp.HttpResponseWrapper, err error) {
	handleInternalError(ctx, err, responser.HandleError)
}

func HandleInternalErrorWithMessage(ctx *gin.Context, responser *xorhttp.HttpResponseWrapper, err error) {
	handleInternalError(ctx, err, responser.HandleErrorWithMessage)
}

type errorHandler func(ctx *gin.Context, code int, err error)

var (
	illegalArgumentError *IllegalArgumentError
)

func handleInternalError(ctx *gin.Context, err error, handler errorHandler) {
	switch {
	case errors.As(err, &illegalArgumentError):
		handler(ctx, http.StatusBadRequest, err)
	default:
		handler(ctx, http.StatusInternalServerError, err)
	}
}
