package product

import (
	"context"
	"github.com/gin-gonic/gin"
	openapitypes "github.com/oapi-codegen/runtime/types"
	global "go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"net/http"
	"xor-go/services/finances/internal/handler/generated/product"
	http2 "xor-go/services/finances/internal/handler/handler"
	"xor-go/services/finances/internal/service/adapters"
)

const (
	spanDefaultProduct = "product/handler."
)

var _ product.ServerInterface = &Handler{}

type Handler struct {
	productService adapters.ProductService
}

func NewProductHandler(productService adapters.ProductService) *Handler {
	return &Handler{productService: productService}
}

func getProductTracerSpan(ctx *gin.Context, name string) (trace.Tracer, context.Context, trace.Span) {
	tr := global.Tracer(adapters.ServiceNameProduct)
	newCtx, span := tr.Start(ctx, spanDefaultProduct+name)

	return tr, newCtx, span
}

func (h *Handler) Get(c *gin.Context, uuid openapitypes.UUID) {
	_, newCtx, span := getProductTracerSpan(c, ".Get")
	defer span.End()

	domain, err := h.productService.Get(newCtx, uuid)
	if err != nil {
		http2.AbortWithBadResponse(c, http2.MapErrorToCode(err), err)
		return
	}

	response := DomainToGet(*domain)

	c.JSON(http.StatusOK, response)
}

func (h *Handler) GetList(c *gin.Context, params product.GetListParams) {
	_, newCtx, span := getProductTracerSpan(c, ".GetList")
	defer span.End()

	domains, err := h.productService.List(newCtx, FilterToDomain(params.Filter))
	if err != nil {
		http2.AbortWithBadResponse(c, http2.MapErrorToCode(err), err)
		return
	}

	list := make([]product.ProductGet, len(domains))
	for i, item := range domains {
		list[i] = DomainToGet(item)
	}

	c.JSON(http.StatusOK, list)
}

func (h *Handler) Create(c *gin.Context, params product.CreateParams) {
	_, newCtx, span := getProductTracerSpan(c, ".Create")
	defer span.End()

	domain := CreateToDomain(params.Model)
	err := h.productService.Create(newCtx, &domain)
	if err != nil {
		http2.AbortWithBadResponse(c, http2.MapErrorToCode(err), err)
		return
	}

	c.JSON(http.StatusOK, http.NoBody)
}

func (h *Handler) Update(c *gin.Context, params product.UpdateParams) {
	_, newCtx, span := getProductTracerSpan(c, ".Update")
	defer span.End()

	domain := UpdateToDomain(params.Model)
	err := h.productService.Update(newCtx, &domain)
	if err != nil {
		http2.AbortWithBadResponse(c, http2.MapErrorToCode(err), err)
		return
	}

	c.JSON(http.StatusOK, http.NoBody)
}

func (h *Handler) Disable(c *gin.Context, id openapitypes.UUID) {
	_, newCtx, span := getProductTracerSpan(c, ".Disable")
	defer span.End()

	err := h.productService.SetAvailability(newCtx, id, false)
	if err != nil {
		http2.AbortWithBadResponse(c, http2.MapErrorToCode(err), err)
		return
	}

	c.JSON(http.StatusOK, http.NoBody)
}
