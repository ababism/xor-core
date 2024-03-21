// Package purchase_request provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.0.0 DO NOT EDIT.
package purchase_request

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/oapi-codegen/runtime"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// PurchaseRequestCreate defines model for PurchaseRequestCreate.
type PurchaseRequestCreate struct {
	ReceivedAt time.Time          `json:"ReceivedAt"`
	Receiver   openapi_types.UUID `json:"Receiver"`
	Sender     openapi_types.UUID `json:"Sender"`
	WebhookURL string             `json:"WebhookURL"`
}

// PurchaseRequestFilter defines model for PurchaseRequestFilter.
type PurchaseRequestFilter struct {
	ReceivedAt *time.Time          `json:"ReceivedAt,omitempty"`
	Receiver   *openapi_types.UUID `json:"Receiver,omitempty"`
	Sender     *openapi_types.UUID `json:"Sender,omitempty"`
	UUID       *openapi_types.UUID `json:"UUID,omitempty"`
	WebhookURL *string             `json:"WebhookURL,omitempty"`
}

// PurchaseRequestGet defines model for PurchaseRequestGet.
type PurchaseRequestGet struct {
	ReceivedAt time.Time          `json:"ReceivedAt"`
	Receiver   openapi_types.UUID `json:"Receiver"`
	Sender     openapi_types.UUID `json:"Sender"`
	UUID       openapi_types.UUID `json:"UUID"`
	WebhookURL string             `json:"WebhookURL"`
}

// GetListParams defines parameters for GetList.
type GetListParams struct {
	Filter *PurchaseRequestFilter `form:"filter,omitempty" json:"filter,omitempty"`
}

// CreateParams defines parameters for Create.
type CreateParams struct {
	Filter PurchaseRequestCreate `form:"filter" json:"filter"`
}

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// List purchase requests
	// (GET /purchase-requests)
	GetList(c *gin.Context, params GetListParams)
	// Create a purchase request
	// (POST /purchase-requests)
	Create(c *gin.Context, params CreateParams)
	// Get purchase request by ID
	// (GET /purchase-requests/{id})
	Get(c *gin.Context, id openapi_types.UUID)
	// Archive a purchase request
	// (PUT /purchase-requests/{id}/archive)
	Archive(c *gin.Context, id openapi_types.UUID)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// GetList operation middleware
func (siw *ServerInterfaceWrapper) GetList(c *gin.Context) {

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetListParams

	// ------------- Optional query parameter "filter" -------------

	err = runtime.BindQueryParameter("form", true, false, "filter", c.Request.URL.Query(), &params.Filter)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter filter: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetList(c, params)
}

// Create operation middleware
func (siw *ServerInterfaceWrapper) Create(c *gin.Context) {

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params CreateParams

	// ------------- Required query parameter "filter" -------------

	if paramValue := c.Query("filter"); paramValue != "" {

	} else {
		siw.ErrorHandler(c, fmt.Errorf("Query argument filter is required, but not found"), http.StatusBadRequest)
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "filter", c.Request.URL.Query(), &params.Filter)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter filter: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.Create(c, params)
}

// Get operation middleware
func (siw *ServerInterfaceWrapper) Get(c *gin.Context) {

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

	siw.Handler.Get(c, id)
}

// Archive operation middleware
func (siw *ServerInterfaceWrapper) Archive(c *gin.Context) {

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

	siw.Handler.Archive(c, id)
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

	router.GET(options.BaseURL+"/purchase-requests", wrapper.GetList)
	router.POST(options.BaseURL+"/purchase-requests", wrapper.Create)
	router.GET(options.BaseURL+"/purchase-requests/:id", wrapper.Get)
	router.PUT(options.BaseURL+"/purchase-requests/:id/archive", wrapper.Archive)
}
