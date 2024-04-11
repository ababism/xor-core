package purchase

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
	spanDefaultPurchaseRequest = "purchase-request/handler."
)

var _ ServerInterface = &Handler{}

type Handler struct {
	purchaseRequestService adapters.PurchaseRequestService
}

func NewPurchaseRequestHandler(purchaseRequestService adapters.PurchaseRequestService) *Handler {
	return &Handler{purchaseRequestService: purchaseRequestService}
}

func getAccountTracerSpan(ctx *gin.Context, name string) (trace.Tracer, context.Context, trace.Span) {
	tr := global.Tracer(adapters.ServiceNamePurchaseRequest)
	newCtx, span := tr.Start(ctx, spanDefaultPurchaseRequest+name)

	return tr, newCtx, span
}

func (h *Handler) GetPurchaseRequestsId(ctx *gin.Context, uuid openapitypes.UUID) {
	_, newCtx, span := getAccountTracerSpan(ctx, ".Get")
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
	_, newCtx, span := getAccountTracerSpan(ctx, ".GetList")
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
	_, newCtx, span := getAccountTracerSpan(ctx, ".Create")
	defer span.End()

	var body PurchaseRequestCreate
	if err := ctx.BindJSON(&body); err != nil {
		http2.AbortWithBadResponse(ctx, http2.MapErrorToCode(err), err)
		return
	}

	domain := CreateToDomain(body)
	err := h.purchaseRequestService.Create(newCtx, &domain)
	if err != nil {
		http2.AbortWithBadResponse(ctx, http2.MapErrorToCode(err), err)
		return
	}

	ctx.JSON(http.StatusOK, http.NoBody)
}

func (h *Handler) PutPurchaseRequestsIdArchive(ctx *gin.Context, id openapitypes.UUID) {
	_, newCtx, span := getAccountTracerSpan(ctx, ".Archive")
	defer span.End()

	err := h.purchaseRequestService.Archive(newCtx, id)
	if err != nil {
		http2.AbortWithBadResponse(ctx, http2.MapErrorToCode(err), err)
		return
	}

	ctx.JSON(http.StatusOK, http.NoBody)
}
