package payment_api

import (
	"context"
	"github.com/gin-gonic/gin"
	openapitypes "github.com/oapi-codegen/runtime/types"
	global "go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"net/http"
	"xor-go/services/finances/internal/handler/generated/payment"
	http2 "xor-go/services/finances/internal/handler/handler"
	"xor-go/services/finances/internal/service/adapters"
)

const (
	spanDefaultPayment = "payment/handler."
)

var _ payment.ServerInterface = &PaymentHandler{}

type PaymentHandler struct {
	paymentService adapters.PaymentService
}

func NewPaymentHandler(paymentService adapters.PaymentService) *PaymentHandler {
	return &PaymentHandler{paymentService: paymentService}
}

func getPaymentTracerSpan(ctx *gin.Context, name string) (trace.Tracer, context.Context, trace.Span) {
	tr := global.Tracer(adapters.ServiceNamePayment)
	newCtx, span := tr.Start(ctx, spanDefaultPayment+name)

	return tr, newCtx, span
}

func (h *PaymentHandler) Get(c *gin.Context, uuid openapitypes.UUID) {
	_, newCtx, span := getPaymentTracerSpan(c, ".Get")
	defer span.End()

	domain, err := h.paymentService.Get(newCtx, uuid)
	if err != nil {
		http2.AbortWithBadResponse(c, http2.MapErrorToCode(err), err)
		return
	}

	response := DomainToGet(*domain)

	c.JSON(http.StatusOK, response)
}

func (h *PaymentHandler) GetList(c *gin.Context, params payment.GetListParams) {
	_, newCtx, span := getPaymentTracerSpan(c, ".GetList")
	defer span.End()

	domains, err := h.paymentService.List(newCtx, FilterToDomain(params.Filter))
	if err != nil {
		http2.AbortWithBadResponse(c, http2.MapErrorToCode(err), err)
		return
	}

	list := make([]payment.PaymentGet, len(domains))
	for i, item := range domains {
		list[i] = DomainToGet(item)
	}

	c.JSON(http.StatusOK, list)
}

func (h *PaymentHandler) Create(c *gin.Context, params payment.CreateParams) {
	_, newCtx, span := getPaymentTracerSpan(c, ".Create")
	defer span.End()

	domain := CreateToDomain(params.Model)
	err := h.paymentService.Create(newCtx, &domain)
	if err != nil {
		http2.AbortWithBadResponse(c, http2.MapErrorToCode(err), err)
		return
	}

	c.JSON(http.StatusOK, http.NoBody)
}
