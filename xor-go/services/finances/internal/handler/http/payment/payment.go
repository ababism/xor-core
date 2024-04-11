package payment

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
	spanDefaultPayment = "payment/handler."
)

var _ ServerInterface = &Handler{}

type Handler struct {
	paymentService adapters.PaymentService
}

func NewPaymentHandler(paymentService adapters.PaymentService) *Handler {
	return &Handler{paymentService: paymentService}
}

func getPaymentTracerSpan(ctx *gin.Context, name string) (trace.Tracer, context.Context, trace.Span) {
	tr := global.Tracer(adapters.ServiceNamePayment)
	newCtx, span := tr.Start(ctx, spanDefaultPayment+name)

	return tr, newCtx, span
}

func (h *Handler) GetPaymentsUuid(ctx *gin.Context, uuid openapitypes.UUID) {
	_, newCtx, span := getPaymentTracerSpan(ctx, ".GetByLogin")
	defer span.End()

	domain, err := h.paymentService.Get(newCtx, uuid)
	if err != nil {
		http2.AbortWithBadResponse(ctx, http2.MapErrorToCode(err), err)
		return
	}

	response := DomainToGet(*domain)

	ctx.JSON(http.StatusOK, response)
}

func (h *Handler) GetPayments(ctx *gin.Context) {
	_, newCtx, span := getPaymentTracerSpan(ctx, ".GetList")
	defer span.End()

	var body *PaymentFilter
	if err := ctx.BindJSON(&body); err != nil && err != io.EOF {
		http2.AbortWithBadResponse(ctx, http2.MapErrorToCode(err), err)
		return
	}

	domains, err := h.paymentService.List(newCtx, FilterToDomain(body))
	if err != nil {
		http2.AbortWithBadResponse(ctx, http2.MapErrorToCode(err), err)
		return
	}

	list := make([]PaymentGet, len(domains))
	for i, item := range domains {
		list[i] = DomainToGet(item)
	}

	ctx.JSON(http.StatusOK, list)
}

func (h *Handler) PostPayments(ctx *gin.Context) {
	_, newCtx, span := getPaymentTracerSpan(ctx, ".Create")
	defer span.End()

	var body PaymentCreate
	if err := ctx.BindJSON(&body); err != nil {
		http2.AbortWithBadResponse(ctx, http2.MapErrorToCode(err), err)
		return
	}

	domain := CreateToDomain(body)
	id, err := h.paymentService.Create(newCtx, &domain)
	if err != nil {
		http2.AbortWithBadResponse(ctx, http2.MapErrorToCode(err), err)
		return
	}

	ctx.JSON(http.StatusOK, ModelUUID{UUID: *id})
}
