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

func (h *Handler) GetProductsId(c *gin.Context, uuid openapitypes.UUID) {
	ctxTrace := global.GetTextMapPropagator().Extract(c, propagation.HeaderCarrier(c.Request.Header))
	_, newCtx, span := getProductTracerSpan(ctxTrace, ".Get")
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
	ctxTrace := global.GetTextMapPropagator().Extract(c, propagation.HeaderCarrier(c.Request.Header))
	_, newCtx, span := getProductTracerSpan(ctxTrace, ".GetPrice")
	defer span.End()

	price, err := h.productService.GetPrice(newCtx, uuids)
	if err != nil {
		http2.AbortWithBadResponse(c, http2.MapErrorToCode(err), err)
		return
	}

	c.JSON(http.StatusOK, price)
}

func (h *Handler) GetProducts(c *gin.Context) {
	ctxTrace := global.GetTextMapPropagator().Extract(c, propagation.HeaderCarrier(c.Request.Header))
	_, newCtx, span := getProductTracerSpan(ctxTrace, ".GetList")
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

func (h *Handler) PostProductsList(c *gin.Context) {
	ctxTrace := global.GetTextMapPropagator().Extract(c, propagation.HeaderCarrier(c.Request.Header))
	_, newCtx, span := getProductTracerSpan(ctxTrace, ".Create")
	defer span.End()

	var body ProductCreate
	if err := c.BindJSON(&body); err != nil {
		http2.AbortWithBadResponse(c, http2.MapErrorToCode(err), err)
		return
	}

	domain := CreateToDomain(body)
	err := h.productService.Create(newCtx, &domain)
	if err != nil {
		http2.AbortWithBadResponse(c, http2.MapErrorToCode(err), err)
		return
	}

	c.JSON(http.StatusOK, http.NoBody)
}

func (h *Handler) PostProducts(c *gin.Context) {
	ctxTrace := global.GetTextMapPropagator().Extract(c, propagation.HeaderCarrier(c.Request.Header))
	_, newCtx, span := getProductTracerSpan(ctxTrace, ".Create")
	defer span.End()

	var body []ProductCreate
	if err := c.BindJSON(&body); err != nil {
		http2.AbortWithBadResponse(c, http2.MapErrorToCode(err), err)
		return
	}

	for _, product := range body {
		domain := CreateToDomain(product)
		err := h.productService.Create(newCtx, &domain)
		if err != nil {
			http2.AbortWithBadResponse(c, http2.MapErrorToCode(err), err)
			return
		}
	}

	c.JSON(http.StatusOK, http.NoBody)
}

func (h *Handler) PutProducts(c *gin.Context) {
	ctxTrace := global.GetTextMapPropagator().Extract(c, propagation.HeaderCarrier(c.Request.Header))
	_, newCtx, span := getProductTracerSpan(ctxTrace, ".Update")
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

func (h *Handler) PutProductsIdDisable(c *gin.Context, id openapitypes.UUID) {
	ctxTrace := global.GetTextMapPropagator().Extract(c, propagation.HeaderCarrier(c.Request.Header))
	_, newCtx, span := getProductTracerSpan(ctxTrace, ".Disable")
	defer span.End()

	err := h.productService.SetAvailability(newCtx, id, false)
	if err != nil {
		http2.AbortWithBadResponse(c, http2.MapErrorToCode(err), err)
		return
	}

	c.JSON(http.StatusOK, http.NoBody)
}
