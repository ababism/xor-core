package payout_request_api

import (
	"context"
	"github.com/gin-gonic/gin"
	openapitypes "github.com/oapi-codegen/runtime/types"
	global "go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"net/http"
	"xor-go/services/finances/internal/handler/generated/payout-request"
	http2 "xor-go/services/finances/internal/handler/handler"
	"xor-go/services/finances/internal/service/adapters"
)

const (
	spanDefaultPayoutRequest = "payout-request/handler."
)

var _ payout_request.ServerInterface = &PayoutRequestHandler{}

type PayoutRequestHandler struct {
	payoutRequestService adapters.PayoutRequestService
}

func NewPayoutRequestHandler(payoutRequestService adapters.PayoutRequestService) *PayoutRequestHandler {
	return &PayoutRequestHandler{payoutRequestService: payoutRequestService}
}

func getAccountTracerSpan(ctx *gin.Context, name string) (trace.Tracer, context.Context, trace.Span) {
	tr := global.Tracer(adapters.ServiceNamePayoutRequest)
	newCtx, span := tr.Start(ctx, spanDefaultPayoutRequest+name)

	return tr, newCtx, span
}

func (h *PayoutRequestHandler) Get(c *gin.Context, uuid openapitypes.UUID) {
	_, newCtx, span := getAccountTracerSpan(c, ".Get")
	defer span.End()

	domain, err := h.payoutRequestService.Get(newCtx, uuid)
	if err != nil {
		http2.AbortWithBadResponse(c, http2.MapErrorToCode(err), err)
		return
	}

	response := DomainToGet(*domain)

	c.JSON(http.StatusOK, response)
}

func (h *PayoutRequestHandler) GetList(c *gin.Context, params payout_request.GetListParams) {
	_, newCtx, span := getAccountTracerSpan(c, ".GetList")
	defer span.End()

	models, err := h.payoutRequestService.List(newCtx, FilterToDomain(params.Filter))
	if err != nil {
		http2.AbortWithBadResponse(c, http2.MapErrorToCode(err), err)
		return
	}

	list := make([]payout_request.PayoutRequestGet, len(models))
	for i, item := range models {
		list[i] = DomainToGet(item)
	}

	c.JSON(http.StatusOK, list)
}

func (h *PayoutRequestHandler) Create(c *gin.Context, params payout_request.CreateParams) {
	_, newCtx, span := getAccountTracerSpan(c, ".Create")
	defer span.End()

	domain := CreateToDomain(params.Model)
	err := h.payoutRequestService.Create(newCtx, &domain)
	if err != nil {
		http2.AbortWithBadResponse(c, http2.MapErrorToCode(err), err)
		return
	}

	c.JSON(http.StatusOK, http.NoBody)
}

func (h *PayoutRequestHandler) Archive(c *gin.Context, id openapitypes.UUID) {
	_, newCtx, span := getAccountTracerSpan(c, ".Archive")
	defer span.End()

	err := h.payoutRequestService.Archive(newCtx, id)
	if err != nil {
		http2.AbortWithBadResponse(c, http2.MapErrorToCode(err), err)
		return
	}

	c.JSON(http.StatusOK, http.NoBody)
}
