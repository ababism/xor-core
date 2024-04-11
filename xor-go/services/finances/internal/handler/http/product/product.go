package product

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
	spanDefaultProduct = "product/handler."
)

var _ ServerInterface = &Handler{}

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

func (h *Handler) GetProductsId(c *gin.Context, uuid openapitypes.UUID) {
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

func (h *Handler) GetProductsPriceUuids(c *gin.Context, uuids []openapitypes.UUID) {
	_, newCtx, span := getProductTracerSpan(c, ".GetPrice")
	defer span.End()

	price, err := h.productService.GetPrice(newCtx, uuids)
	if err != nil {
		http2.AbortWithBadResponse(c, http2.MapErrorToCode(err), err)
		return
	}

	c.JSON(http.StatusOK, price)
}

func (h *Handler) GetProducts(ctx *gin.Context) {
	_, newCtx, span := getProductTracerSpan(ctx, ".GetList")
	defer span.End()

	var body *ProductFilter
	if err := ctx.BindJSON(&body); err != nil && err != io.EOF {
		http2.AbortWithBadResponse(ctx, http2.MapErrorToCode(err), err)
		return
	}

	domains, err := h.productService.List(newCtx, FilterToDomain(body))
	if err != nil {
		http2.AbortWithBadResponse(ctx, http2.MapErrorToCode(err), err)
		return
	}

	list := make([]ProductGet, len(domains))
	for i, item := range domains {
		list[i] = DomainToGet(item)
	}

	ctx.JSON(http.StatusOK, list)
}

func (h *Handler) PostProductsList(ctx *gin.Context) {
	_, newCtx, span := getProductTracerSpan(ctx, ".Create")
	defer span.End()

	var body ProductCreate
	if err := ctx.BindJSON(&body); err != nil {
		http2.AbortWithBadResponse(ctx, http2.MapErrorToCode(err), err)
		return
	}

	domain := CreateToDomain(body)
	err := h.productService.Create(newCtx, &domain)
	if err != nil {
		http2.AbortWithBadResponse(ctx, http2.MapErrorToCode(err), err)
		return
	}

	ctx.JSON(http.StatusOK, http.NoBody)
}

func (h *Handler) PostProducts(ctx *gin.Context) {
	_, newCtx, span := getProductTracerSpan(ctx, ".Create")
	defer span.End()

	var body []ProductCreate
	if err := ctx.BindJSON(&body); err != nil {
		http2.AbortWithBadResponse(ctx, http2.MapErrorToCode(err), err)
		return
	}

	for _, product := range body {
		domain := CreateToDomain(product)
		err := h.productService.Create(newCtx, &domain)
		if err != nil {
			http2.AbortWithBadResponse(ctx, http2.MapErrorToCode(err), err)
			return
		}
	}

	ctx.JSON(http.StatusOK, http.NoBody)
}

func (h *Handler) PutProducts(ctx *gin.Context) {
	_, newCtx, span := getProductTracerSpan(ctx, ".Update")
	defer span.End()

	var body ProductUpdate
	if err := ctx.BindJSON(&body); err != nil {
		http2.AbortWithBadResponse(ctx, http2.MapErrorToCode(err), err)
		return
	}

	domain := UpdateToDomain(body)
	err := h.productService.Update(newCtx, &domain)
	if err != nil {
		http2.AbortWithBadResponse(ctx, http2.MapErrorToCode(err), err)
		return
	}

	ctx.JSON(http.StatusOK, http.NoBody)
}

func (h *Handler) PutProductsIdDisable(c *gin.Context, id openapitypes.UUID) {
	_, newCtx, span := getProductTracerSpan(c, ".Disable")
	defer span.End()

	err := h.productService.SetAvailability(newCtx, id, false)
	if err != nil {
		http2.AbortWithBadResponse(c, http2.MapErrorToCode(err), err)
		return
	}

	c.JSON(http.StatusOK, http.NoBody)
}
