package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"net/http"
	"xor-go/pkg/xhttp/response"
	"xor-go/services/sage/internal/api/http/model"
	"xor-go/services/sage/internal/domain"
	"xor-go/services/sage/internal/service/adapter"
)

type GatewayHandler struct {
	logger         *zap.Logger
	responser      *response.HttpResponseWrapper
	gatewayService adapter.GatewayService
}

func NewGatewayHandler(
	logger *zap.Logger,
	responser *response.HttpResponseWrapper,
	gatewayService adapter.GatewayService,
) *GatewayHandler {
	return &GatewayHandler{
		logger:         logger,
		responser:      responser,
		gatewayService: gatewayService,
	}
}

func (r *GatewayHandler) InitRoutes(g *gin.RouterGroup) {
	gateway := g.Group("/gateway")
	gateway.POST("/secure/pass", r.PassSecure)
	gateway.POST("/insecure/pass", r.PassInsecure)
}

func (r *GatewayHandler) PassSecure(ctx *gin.Context) {
	var passSecureResourceRequest model.PassSecureResourceRequest
	if err := ctx.BindJSON(&passSecureResourceRequest); err != nil {
		r.responser.HandleErrorWithMessage(ctx, http.StatusBadRequest, err)
		return
	}

	requestUuid, err := uuid.NewUUID()
	if err != nil {
		r.responser.HandleXorError(ctx, err)
		return
	}
	r.logger.Info(
		fmt.Sprintf(
			"secure request: rq_uuid=%v; access_token=%v",
			requestUuid,
			passSecureResourceRequest.AccessToken,
		),
	)

	idmVerifyResponse, err := r.gatewayService.Verify(ctx, model.ToPassSecureResourceInfo(&passSecureResourceRequest))
	if err != nil {
		r.responser.HandleXorError(ctx, err)
		return
	}

	internalResourceResponse, err := r.gatewayService.PassSecure(ctx, &domain.PassSecureResourceRequest{
		RequestUuid:  requestUuid,
		Resource:     passSecureResourceRequest.Resource,
		Route:        passSecureResourceRequest.Route,
		Method:       passSecureResourceRequest.Method,
		Body:         passSecureResourceRequest.Body,
		AccountUuid:  idmVerifyResponse.AccountUuid,
		AccountEmail: idmVerifyResponse.AccountEmail,
		Roles:        idmVerifyResponse.Roles,
	})
	if err != nil {
		r.responser.HandleXorError(ctx, err)
		return
	}

	ctx.JSON(200, model.ToPassResourceResponse(internalResourceResponse))
}

func (r *GatewayHandler) PassInsecure(ctx *gin.Context) {
	var passInsecureResourceRequest model.PassInsecureResourceRequest
	if err := ctx.BindJSON(&passInsecureResourceRequest); err != nil {
		r.responser.HandleErrorWithMessage(ctx, http.StatusBadRequest, err)
		return
	}

	requestUuid, err := uuid.NewUUID()
	if err != nil {
		r.responser.HandleXorError(ctx, err)
		return
	}
	r.logger.Info(fmt.Sprintf("insecure request: rq_uuid=%v", requestUuid))

	internalResourceResponse, err := r.gatewayService.PassInsecure(ctx, &domain.PassInsecureResourceRequest{
		RequestUuid: requestUuid,
		Resource:    passInsecureResourceRequest.Resource,
		Route:       passInsecureResourceRequest.Route,
		Method:      passInsecureResourceRequest.Method,
		Body:        passInsecureResourceRequest.Body,
	})
	if err != nil {
		r.responser.HandleXorError(ctx, err)
		return
	}

	ctx.JSON(200, model.ToPassResourceResponse(internalResourceResponse))
}
