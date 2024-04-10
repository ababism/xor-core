package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"xor-go/pkg/xhttp/response"
	"xor-go/services/sage/internal/api/http/model"
	"xor-go/services/sage/internal/domain"
	"xor-go/services/sage/internal/service/adapter"
	"xor-go/services/sage/pkg/idm"
)

type GatewayHandler struct {
	logger         *zap.Logger
	responser      *response.HttpResponseWrapper
	gatewayService adapter.GatewayService
}

func NewGatewayHandler(logger *zap.Logger, responser *response.HttpResponseWrapper, gatewayService adapter.GatewayService, client *idm.Client) *GatewayHandler {
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
	var gatewayPassRequest model.PassSecureResourceRequest
	if err := ctx.BindJSON(&gatewayPassRequest); err != nil {
		r.responser.HandleErrorWithMessage(ctx, http.StatusBadRequest, err)
		return
	}

	idmVerifyResponse, err := r.gatewayService.Verify(ctx, model.ToPassSecureResourceInfo(&gatewayPassRequest))
	if err != nil {
		r.responser.HandleXorErrorWithMessage(ctx, err)
		return
	}

	internalResourceAccessResponse, err := r.gatewayService.PassSecure(ctx, &domain.PassSecureResourceRequest{
		Resource:     gatewayPassRequest.Resource,
		Route:        gatewayPassRequest.Route,
		Method:       gatewayPassRequest.Method,
		Body:         gatewayPassRequest.Body,
		AccountUuid:  idmVerifyResponse.AccountUuid,
		AccountEmail: idmVerifyResponse.AccountEmail,
		Roles:        idmVerifyResponse.Roles,
	})
	if err != nil {
		r.responser.HandleXorErrorWithMessage(ctx, err)
		return
	}

	ctx.JSON(200, model.ToPassResourceResponse(internalResourceAccessResponse))
}

func (r *GatewayHandler) PassInsecure(ctx *gin.Context) {
	var gatewayPassRequest model.PassInsecureResourceRequest
	if err := ctx.BindJSON(&gatewayPassRequest); err != nil {
		r.responser.HandleErrorWithMessage(ctx, http.StatusBadRequest, err)
		return
	}

	//idmVerifyResponse, err := r.gatewayService.Verify(ctx, gatewayPassRequest.ToPassSecureResourceInfo())
	//if err != nil {
	//	r.responser.HandleXorErrorWithMessage(ctx, err)
	//	return
	//}

	internalResourceAccessResponse, err := r.gatewayService.PassInsecure(ctx, &domain.PassInsecureResourceRequest{
		Resource: gatewayPassRequest.Resource,
		Route:    gatewayPassRequest.Route,
		Method:   gatewayPassRequest.Method,
		Body:     gatewayPassRequest.Body,
	})
	if err != nil {
		r.responser.HandleXorErrorWithMessage(ctx, err)
		return
	}

	ctx.JSON(200, model.ToPassResourceResponse(internalResourceAccessResponse))
}
