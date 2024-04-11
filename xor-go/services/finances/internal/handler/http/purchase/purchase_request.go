package purchase

import (
	"context"
	"github.com/gin-gonic/gin"
	openapitypes "github.com/oapi-codegen/runtime/types"
	global "go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	"io"
	"net/http"
	"xor-go/services/finances/internal/handler/http/dto"
	http2 "xor-go/services/finances/internal/handler/http/utils"
	"xor-go/services/finances/internal/service/adapters"
)

const (
	spanDefaultPurchaseRequest = "purchase-request/handler."
)

var _ ServerInterface = &Handler{}

type Handler struct {
	purchaseRequestService adapters.PurchaseRequestService
}

func NewPurchaseRequestHandler(purchaseRequestService adapters.PurchaseRequestService) *Handler {
	return &Handler{purchaseRequestService: purchaseRequestService}
}

func getAccountTracerSpan(ctx context.Context, name string) (trace.Tracer, context.Context, trace.Span) {
	tr := global.Tracer(adapters.ServiceNamePurchaseRequest)
	newCtx, span := tr.Start(ctx, spanDefaultPurchaseRequest+name)

	return tr, newCtx, span
}

func (h *Handler) GetPurchaseRequestsId(ctx *gin.Context, uuid openapitypes.UUID) {
	ctxTrace := global.GetTextMapPropagator().Extract(ctx, propagation.HeaderCarrier(ctx.Request.Header))
	_, newCtx, span := getAccountTracerSpan(ctxTrace, ".GetByLogin")
	defer span.End()

	domain, err := h.purchaseRequestService.Get(newCtx, uuid)
	if err != nil {
		http2.AbortWithBadResponse(ctx, http2.MapErrorToCode(err), err)
		return
	}

	response := DomainToGet(*domain)

	ctx.JSON(http.StatusOK, response)
}

func (h *Handler) GetPurchaseRequests(ctx *gin.Context) {
	ctxTrace := global.GetTextMapPropagator().Extract(ctx, propagation.HeaderCarrier(ctx.Request.Header))
	_, newCtx, span := getAccountTracerSpan(ctxTrace, ".GetList")
	defer span.End()

	var body *PurchaseRequestFilter
	if err := ctx.BindJSON(&body); err != nil && err != io.EOF {
		http2.AbortWithBadResponse(ctx, http2.MapErrorToCode(err), err)
		return
	}

	models, err := h.purchaseRequestService.List(newCtx, FilterToDomain(body))
	if err != nil {
		http2.AbortWithBadResponse(ctx, http2.MapErrorToCode(err), err)
		return
	}

	list := make([]PurchaseRequestGet, len(models))
	for i, item := range models {
		list[i] = DomainToGet(item)
	}

	ctx.JSON(http.StatusOK, list)
}

func (h *Handler) PostPurchaseRequests(ctx *gin.Context) {
	ctxTrace := global.GetTextMapPropagator().Extract(ctx, propagation.HeaderCarrier(ctx.Request.Header))
	_, newCtx, span := getAccountTracerSpan(ctxTrace, ".Create")
	defer span.End()

	var body PurchaseRequestCreate
	if err := ctx.BindJSON(&body); err != nil {
		http2.AbortWithBadResponse(ctx, http2.MapErrorToCode(err), err)
		return
	}

	domain := CreateToDomain(body)
	id, err := h.purchaseRequestService.Create(newCtx, &domain)
	if err != nil {
		http2.AbortWithBadResponse(ctx, http2.MapErrorToCode(err), err)
		return
	}

	ctx.JSON(http.StatusOK, dto.ModelUUID{UUID: *id})
}

func (h *Handler) PutPurchaseRequestsIdArchive(ctx *gin.Context, id openapitypes.UUID) {
	ctxTrace := global.GetTextMapPropagator().Extract(ctx, propagation.HeaderCarrier(ctx.Request.Header))
	_, newCtx, span := getAccountTracerSpan(ctxTrace, ".Archive")
	defer span.End()

	err := h.purchaseRequestService.Archive(newCtx, id)
	if err != nil {
		http2.AbortWithBadResponse(ctx, http2.MapErrorToCode(err), err)
		return
	}

	ctx.JSON(http.StatusOK, http.NoBody)
}
