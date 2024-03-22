package purchase

import (
	"context"
	"github.com/gin-gonic/gin"
	openapitypes "github.com/oapi-codegen/runtime/types"
	global "go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"net/http"
	purchaserequest "xor-go/services/finances/internal/handler/generated/purchase-request"
	http2 "xor-go/services/finances/internal/handler/handler"
	"xor-go/services/finances/internal/service/adapters"
)

const (
	spanDefaultPurchaseRequest = "purchase-request/handler."
)

var _ purchaserequest.ServerInterface = &Handler{}

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

func (h *Handler) Get(c *gin.Context, uuid openapitypes.UUID) {
	_, newCtx, span := getAccountTracerSpan(c, ".Get")
	defer span.End()

	domain, err := h.purchaseRequestService.Get(newCtx, uuid)
	if err != nil {
		http2.AbortWithBadResponse(c, http2.MapErrorToCode(err), err)
		return
	}

	response := DomainToGet(*domain)

	c.JSON(http.StatusOK, response)
}

func (h *Handler) GetList(c *gin.Context, params purchaserequest.GetListParams) {
	_, newCtx, span := getAccountTracerSpan(c, ".GetList")
	defer span.End()

	models, err := h.purchaseRequestService.List(newCtx, FilterToDomain(params.Filter))
	if err != nil {
		http2.AbortWithBadResponse(c, http2.MapErrorToCode(err), err)
		return
	}

	list := make([]purchaserequest.PurchaseRequestGet, len(models))
	for i, item := range models {
		list[i] = DomainToGet(item)
	}

	c.JSON(http.StatusOK, list)
}

func (h *Handler) Create(c *gin.Context, params purchaserequest.CreateParams) {
	_, newCtx, span := getAccountTracerSpan(c, ".Create")
	defer span.End()

	domain := CreateToDomain(params.Model)
	err := h.purchaseRequestService.Create(newCtx, &domain)
	if err != nil {
		http2.AbortWithBadResponse(c, http2.MapErrorToCode(err), err)
		return
	}

	c.JSON(http.StatusOK, http.NoBody)
}

func (h *Handler) Archive(c *gin.Context, id openapitypes.UUID) {
	_, newCtx, span := getAccountTracerSpan(c, ".Archive")
	defer span.End()

	err := h.purchaseRequestService.Archive(newCtx, id)
	if err != nil {
		http2.AbortWithBadResponse(c, http2.MapErrorToCode(err), err)
		return
	}

	c.JSON(http.StatusOK, http.NoBody)
}
