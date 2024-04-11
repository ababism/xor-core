package product

import (
	"context"
	"github.com/gin-gonic/gin"
	openapitypes "github.com/oapi-codegen/runtime/types"
	global "go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
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

func getProductTracerSpan(ctx context.Context, name string) (trace.Tracer, context.Context, trace.Span) {
	tr := global.Tracer(adapters.ServiceNameProduct)
	newCtx, span := tr.Start(ctx, spanDefaultProduct+name)

	return tr, newCtx, span
}

func (h *Handler) GetProductId(c *gin.Context, uuid openapitypes.UUID) {
	ctxTrace := global.GetTextMapPropagator().Extract(c, propagation.HeaderCarrier(c.Request.Header))
	_, newCtx, span := getProductTracerSpan(ctxTrace, ".GetProductId")
	defer span.End()

	domain, err := h.productService.Get(newCtx, uuid)
	if err != nil {
		http2.AbortWithBadResponse(c, http2.MapErrorToCode(err), err)
		return
	}

	response := DomainToGet(*domain)

	c.JSON(http.StatusOK, response)
}

func (h *Handler) GetProductPriceUuids(c *gin.Context, uuids []openapitypes.UUID) {
	ctxTrace := global.GetTextMapPropagator().Extract(c, propagation.HeaderCarrier(c.Request.Header))
	_, newCtx, span := getProductTracerSpan(ctxTrace, ".GetProductPriceUuids")
	defer span.End()

	price, err := h.productService.GetPrice(newCtx, uuids)
	if err != nil {
		http2.AbortWithBadResponse(c, http2.MapErrorToCode(err), err)
		return
	}

	c.JSON(http.StatusOK, price)
}

func (h *Handler) GetProduct(c *gin.Context) {
	ctxTrace := global.GetTextMapPropagator().Extract(c, propagation.HeaderCarrier(c.Request.Header))
	_, newCtx, span := getProductTracerSpan(ctxTrace, ".GetProduct")
	defer span.End()

	var body *ProductFilter
	if err := c.BindJSON(&body); err != nil && err != io.EOF {
		http2.AbortWithBadResponse(c, http2.MapErrorToCode(err), err)
		return
	}

	domains, err := h.productService.List(newCtx, FilterToDomain(body))
	if err != nil {
		http2.AbortWithBadResponse(c, http2.MapErrorToCode(err), err)
		return
	}

	list := make([]ProductGet, len(domains))
	for i, item := range domains {
		list[i] = DomainToGet(item)
	}

	c.JSON(http.StatusOK, list)
}

func (h *Handler) PostProductList(c *gin.Context) {
	ctxTrace := global.GetTextMapPropagator().Extract(c, propagation.HeaderCarrier(c.Request.Header))
	_, newCtx, span := getProductTracerSpan(ctxTrace, ".PostProductList")
	defer span.End()

	var body []ProductCreate
	if err := c.BindJSON(&body); err != nil {
		http2.AbortWithBadResponse(c, http2.MapErrorToCode(err), err)
		return
	}

	productIds := make([]ModelUUID, 0)

	for _, product := range body {
		domain := CreateToDomain(product)
		productId, err := h.productService.Create(newCtx, &domain)
		if err != nil {
			http2.AbortWithBadResponse(c, http2.MapErrorToCode(err), err)
			return
		}
		productIds = append(productIds, ModelUUID{*productId})
	}

	c.JSON(http.StatusOK, productIds)
}

func (h *Handler) PostProduct(c *gin.Context) {
	ctxTrace := global.GetTextMapPropagator().Extract(c, propagation.HeaderCarrier(c.Request.Header))
	_, newCtx, span := getProductTracerSpan(ctxTrace, ".PostProduct")
	defer span.End()

	var body ProductCreate
	if err := c.BindJSON(&body); err != nil {
		http2.AbortWithBadResponse(c, http2.MapErrorToCode(err), err)
		return
	}

	domain := CreateToDomain(body)
	id, err := h.productService.Create(newCtx, &domain)
	if err != nil {
		http2.AbortWithBadResponse(c, http2.MapErrorToCode(err), err)
		return
	}

	c.JSON(http.StatusOK, ModelUUID{UUID: *id})
}

func (h *Handler) PutProduct(c *gin.Context) {
	ctxTrace := global.GetTextMapPropagator().Extract(c, propagation.HeaderCarrier(c.Request.Header))
	_, newCtx, span := getProductTracerSpan(ctxTrace, ".PutProduct")
	defer span.End()

	var body ProductUpdate
	if err := c.BindJSON(&body); err != nil {
		http2.AbortWithBadResponse(c, http2.MapErrorToCode(err), err)
		return
	}

	domain := UpdateToDomain(body)
	err := h.productService.Update(newCtx, &domain)
	if err != nil {
		http2.AbortWithBadResponse(c, http2.MapErrorToCode(err), err)
		return
	}

	c.JSON(http.StatusOK, http.NoBody)
}

func (h *Handler) PutProductIdDisable(c *gin.Context, id openapitypes.UUID) {
	ctxTrace := global.GetTextMapPropagator().Extract(c, propagation.HeaderCarrier(c.Request.Header))
	_, newCtx, span := getProductTracerSpan(ctxTrace, ".PutProductIdDisable")
	defer span.End()

	err := h.productService.SetAvailability(newCtx, id, false)
	if err != nil {
		http2.AbortWithBadResponse(c, http2.MapErrorToCode(err), err)
		return
	}

	c.JSON(http.StatusOK, http.NoBody)
}
