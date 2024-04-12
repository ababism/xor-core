package payout

import (
	"context"
	"github.com/gin-gonic/gin"
	openapitypes "github.com/oapi-codegen/runtime/types"
	global "go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"io"
	"net/http"
	http2 "xor-go/services/finances/internal/handler/http/utils"
	"xor-go/services/finances/internal/service/adapters"
)

const (
	spanDefaultPayoutRequest = "payout-request/handler."
)

var _ ServerInterface = &Handler{}

type Handler struct {
	payoutRequestService adapters.PayoutRequestService
}

func NewPayoutRequestHandler(payoutRequestService adapters.PayoutRequestService) *Handler {
	return &Handler{payoutRequestService: payoutRequestService}
}

func getAccountTracerSpan(ctx *gin.Context, name string) (trace.Tracer, context.Context, trace.Span) {
	tr := global.Tracer(adapters.ServiceNamePayoutRequest)
	newCtx, span := tr.Start(ctx, spanDefaultPayoutRequest+name)

	return tr, newCtx, span
}

func (h *Handler) GetPayoutRequestsId(ctx *gin.Context, uuid openapitypes.UUID) {
	_, newCtx, span := getAccountTracerSpan(ctx, ".GetByLogin")
	defer span.End()

	domain, err := h.payoutRequestService.Get(newCtx, uuid)
	if err != nil {
		http2.AbortWithBadResponse(ctx, http2.MapErrorToCode(err), err)
		return
	}

	response := DomainToGet(*domain)

	ctx.JSON(http.StatusOK, response)
}

func (h *Handler) GetPayoutRequests(ctx *gin.Context) {
	_, newCtx, span := getAccountTracerSpan(ctx, ".GetList")
	defer span.End()

	var body *PayoutRequestFilter
	if err := ctx.BindJSON(&body); err != nil && err != io.EOF {
		http2.AbortWithBadResponse(ctx, http2.MapErrorToCode(err), err)
		return
	}

	models, err := h.payoutRequestService.List(newCtx, FilterToDomain(body))
	if err != nil {
		http2.AbortWithBadResponse(ctx, http2.MapErrorToCode(err), err)
		return
	}

	list := make([]PayoutRequestGet, len(models))
	for i, item := range models {
		list[i] = DomainToGet(item)
	}

	ctx.JSON(http.StatusOK, list)
}

func (h *Handler) PostPayoutRequests(ctx *gin.Context) {
	_, newCtx, span := getAccountTracerSpan(ctx, ".Create")
	defer span.End()

	var body PayoutRequestCreate
	if err := ctx.BindJSON(&body); err != nil {
		http2.AbortWithBadResponse(ctx, http2.MapErrorToCode(err), err)
		return
	}

	domain := CreateToDomain(body)
	id, err := h.payoutRequestService.Create(newCtx, &domain)
	if err != nil {
		http2.AbortWithBadResponse(ctx, http2.MapErrorToCode(err), err)
		return
	}

	ctx.JSON(http.StatusOK, ModelUUID{UUID: *id})
}

func (h *Handler) PutPayoutRequestsIdArchive(ctx *gin.Context, id openapitypes.UUID) {
	_, newCtx, span := getAccountTracerSpan(ctx, ".Archive")
	defer span.End()

	err := h.payoutRequestService.Archive(newCtx, id)
	if err != nil {
		http2.AbortWithBadResponse(ctx, http2.MapErrorToCode(err), err)
		return
	}

	ctx.JSON(http.StatusOK, http.NoBody)
}
