package discount

import (
	"context"
	"github.com/gin-gonic/gin"
	openapitypes "github.com/oapi-codegen/runtime/types"
	global "go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"net/http"
	http2 "xor-go/services/finances/internal/handler/http"
	"xor-go/services/finances/internal/service/adapters"
)

const (
	spanDefaultDiscount = "discount/handler."
)

var _ ServerInterface = &Handler{}

type Handler struct {
	discountService adapters.DiscountService
}

func NewDiscountHandler(discountService adapters.DiscountService) *Handler {
	return &Handler{discountService: discountService}
}

func getDiscountTracerSpan(ctx *gin.Context, name string) (trace.Tracer, context.Context, trace.Span) {
	tr := global.Tracer(adapters.ServiceNameDiscount)
	newCtx, span := tr.Start(ctx, spanDefaultDiscount+name)

	return tr, newCtx, span
}

func (h *Handler) Get(c *gin.Context, uuid openapitypes.UUID) {
	_, newCtx, span := getDiscountTracerSpan(c, ".Get")
	defer span.End()

	domain, err := h.discountService.Get(newCtx, uuid)
	if err != nil {
		http2.AbortWithBadResponse(c, http2.MapErrorToCode(err), err)
		return
	}

	response := DomainToGet(*domain)

	c.JSON(http.StatusOK, response)
}

func (h *Handler) GetList(c *gin.Context, params GetListParams) {
	_, newCtx, span := getDiscountTracerSpan(c, ".GetList")
	defer span.End()

	domains, err := h.discountService.List(newCtx, FilterToDomain(params.Filter))
	if err != nil {
		http2.AbortWithBadResponse(c, http2.MapErrorToCode(err), err)
		return
	}

	list := make([]DiscountGet, len(domains))
	for i, item := range domains {
		list[i] = DomainToGet(item)
	}

	c.JSON(http.StatusOK, list)
}

func (h *Handler) Create(c *gin.Context, params CreateParams) {
	_, newCtx, span := getDiscountTracerSpan(c, ".Create")
	defer span.End()

	domain := CreateToDomain(params.Model)
	err := h.discountService.Create(newCtx, &domain)
	if err != nil {
		http2.AbortWithBadResponse(c, http2.MapErrorToCode(err), err)
		return
	}

	c.JSON(http.StatusOK, http.NoBody)
}

func (h *Handler) Update(c *gin.Context, params UpdateParams) {
	_, newCtx, span := getDiscountTracerSpan(c, ".Update")
	defer span.End()

	domain := UpdateToDomain(params.Model)
	err := h.discountService.Update(newCtx, &domain)
	if err != nil {
		http2.AbortWithBadResponse(c, http2.MapErrorToCode(err), err)
		return
	}

	c.JSON(http.StatusOK, http.NoBody)
}

func (h *Handler) End(c *gin.Context, id openapitypes.UUID) {
	_, newCtx, span := getDiscountTracerSpan(c, ".End")
	defer span.End()

	err := h.discountService.EndDiscount(newCtx, id)
	if err != nil {
		http2.AbortWithBadResponse(c, http2.MapErrorToCode(err), err)
		return
	}

	c.JSON(http.StatusOK, http.NoBody)
}
