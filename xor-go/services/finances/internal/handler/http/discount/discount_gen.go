// Package discount provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.0.0 DO NOT EDIT.
package discount

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/oapi-codegen/runtime"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// DiscountCreate defines model for DiscountCreate.
type DiscountCreate struct {
	CreatedBy openapi_types.UUID `json:"CreatedBy"`
	EndedAt   time.Time          `json:"EndedAt"`
	Percent   float32            `json:"Percent"`
	StartedAt time.Time          `json:"StartedAt"`
	Status    string             `json:"Status"`
}

// DiscountFilter defines model for DiscountFilter.
type DiscountFilter struct {
	CreatedBy *openapi_types.UUID `json:"CreatedBy,omitempty"`
	Percent   *float32            `json:"Percent,omitempty"`
	Status    *string             `json:"Status,omitempty"`
	UUID      *openapi_types.UUID `json:"UUID,omitempty"`
}

// DiscountGet defines model for DiscountGet.
type DiscountGet struct {
	CreatedAt    time.Time          `json:"CreatedAt"`
	CreatedBy    openapi_types.UUID `json:"CreatedBy"`
	EndedAt      time.Time          `json:"EndedAt"`
	LastUpdateAt time.Time          `json:"LastUpdateAt"`
	Percent      float32            `json:"Percent"`
	StartedAt    time.Time          `json:"StartedAt"`
	Status       string             `json:"Status"`
	UUID         openapi_types.UUID `json:"UUID"`
}

// DiscountUpdate defines model for DiscountUpdate.
type DiscountUpdate struct {
	CreatedBy openapi_types.UUID `json:"CreatedBy"`
	EndedAt   time.Time          `json:"EndedAt"`
	Percent   float32            `json:"Percent"`
	StartedAt time.Time          `json:"StartedAt"`
	Status    string             `json:"Status"`
	UUID      openapi_types.UUID `json:"UUID"`
}

// GetDiscountsJSONRequestBody defines body for GetDiscounts for application/json ContentType.
type GetDiscountsJSONRequestBody = DiscountFilter

// PostDiscountsJSONRequestBody defines body for PostDiscounts for application/json ContentType.
type PostDiscountsJSONRequestBody = DiscountCreate

// PutDiscountsJSONRequestBody defines body for PutDiscounts for application/json ContentType.
type PutDiscountsJSONRequestBody = DiscountUpdate

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// List discounts
	// (GET /discounts)
	GetDiscounts(c *gin.Context)
	// Create a discount
	// (POST /discounts)
	PostDiscounts(c *gin.Context)
	// Update a discount
	// (PUT /discounts)
	PutDiscounts(c *gin.Context)
	// Get discount by ID
	// (GET /discounts/{id})
	GetDiscountsId(c *gin.Context, id openapi_types.UUID)
	// End a discount
	// (PATCH /discounts/{id}/end)
	PatchDiscountsIdEnd(c *gin.Context, id openapi_types.UUID)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// GetDiscounts operation middleware
func (siw *ServerInterfaceWrapper) GetDiscounts(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetDiscounts(c)
}

// PostDiscounts operation middleware
func (siw *ServerInterfaceWrapper) PostDiscounts(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostDiscounts(c)
}

// PutDiscounts operation middleware
func (siw *ServerInterfaceWrapper) PutDiscounts(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PutDiscounts(c)
}

// GetDiscountsId operation middleware
func (siw *ServerInterfaceWrapper) GetDiscountsId(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id openapi_types.UUID

	err = runtime.BindStyledParameter("simple", false, "id", c.Param("id"), &id)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetDiscountsId(c, id)
}

// PatchDiscountsIdEnd operation middleware
func (siw *ServerInterfaceWrapper) PatchDiscountsIdEnd(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id openapi_types.UUID

	err = runtime.BindStyledParameter("simple", false, "id", c.Param("id"), &id)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PatchDiscountsIdEnd(c, id)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router gin.IRouter, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router gin.IRouter, si ServerInterface, options GinServerOptions) {
	errorHandler := options.ErrorHandler
	if errorHandler == nil {
		errorHandler = func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"msg": err.Error()})
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandler:       errorHandler,
	}

	router.GET(options.BaseURL+"/discounts", wrapper.GetDiscounts)
	router.POST(options.BaseURL+"/discounts", wrapper.PostDiscounts)
	router.PUT(options.BaseURL+"/discounts", wrapper.PutDiscounts)
	router.GET(options.BaseURL+"/discounts/:id", wrapper.GetDiscountsId)
	router.PATCH(options.BaseURL+"/discounts/:id/end", wrapper.PatchDiscountsIdEnd)
}
