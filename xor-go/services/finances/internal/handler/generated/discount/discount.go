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
	CreatedBy  openapi_types.UUID `json:"CreatedBy"`
	EndedAt    time.Time          `json:"EndedAt"`
	Percent    float64            `json:"Percent"`
	StandAlone bool               `json:"StandAlone"`
	StartedAt  time.Time          `json:"StartedAt"`
	Status     string             `json:"Status"`
}

// DiscountFilter defines model for DiscountFilter.
type DiscountFilter struct {
	CreatedBy  *openapi_types.UUID `json:"CreatedBy,omitempty"`
	Percent    *float64            `json:"Percent,omitempty"`
	StandAlone *bool               `json:"StandAlone,omitempty"`
	Status     *string             `json:"Status,omitempty"`
	UUID       *openapi_types.UUID `json:"UUID,omitempty"`
}

// DiscountGet defines model for DiscountGet.
type DiscountGet struct {
	CreatedAt    time.Time          `json:"CreatedAt"`
	CreatedBy    openapi_types.UUID `json:"CreatedBy"`
	EndedAt      time.Time          `json:"EndedAt"`
	LastUpdateAt time.Time          `json:"LastUpdateAt"`
	Percent      float64            `json:"Percent"`
	StandAlone   bool               `json:"StandAlone"`
	StartedAt    time.Time          `json:"StartedAt"`
	Status       string             `json:"Status"`
	UUID         openapi_types.UUID `json:"UUID"`
}

// DiscountUpdate defines model for DiscountUpdate.
type DiscountUpdate struct {
	CreatedBy  openapi_types.UUID `json:"CreatedBy"`
	EndedAt    time.Time          `json:"EndedAt"`
	Percent    float64            `json:"Percent"`
	StandAlone bool               `json:"StandAlone"`
	StartedAt  time.Time          `json:"StartedAt"`
	Status     string             `json:"Status"`
	UUID       openapi_types.UUID `json:"UUID"`
}

// GetListParams defines parameters for GetList.
type GetListParams struct {
	Filter *DiscountFilter `form:"filter,omitempty" json:"filter,omitempty"`
}

// CreateParams defines parameters for Create.
type CreateParams struct {
	Model DiscountCreate `form:"model" json:"model"`
}

// UpdateParams defines parameters for Update.
type UpdateParams struct {
	Model DiscountUpdate `form:"model" json:"model"`
}

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// List discounts
	// (GET /discounts)
	GetList(c *gin.Context, params GetListParams)
	// Create a discount
	// (POST /discounts)
	Create(c *gin.Context, params CreateParams)
	// Update a discount
	// (PUT /discounts)
	Update(c *gin.Context, params UpdateParams)
	// Get discount by ID
	// (GET /discounts/{id})
	Get(c *gin.Context, id openapi_types.UUID)
	// End a discount
	// (PATCH /discounts/{id}/end)
	End(c *gin.Context, id openapi_types.UUID)
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

	// ------------- Required query parameter "model" -------------

	if paramValue := c.Query("model"); paramValue != "" {

	} else {
		siw.ErrorHandler(c, fmt.Errorf("Query argument model is required, but not found"), http.StatusBadRequest)
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "model", c.Request.URL.Query(), &params.Model)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter model: %w", err), http.StatusBadRequest)
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

// Update operation middleware
func (siw *ServerInterfaceWrapper) Update(c *gin.Context) {

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params UpdateParams

	// ------------- Required query parameter "model" -------------

	if paramValue := c.Query("model"); paramValue != "" {

	} else {
		siw.ErrorHandler(c, fmt.Errorf("Query argument model is required, but not found"), http.StatusBadRequest)
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "model", c.Request.URL.Query(), &params.Model)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter model: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.Update(c, params)
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

// End operation middleware
func (siw *ServerInterfaceWrapper) End(c *gin.Context) {

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

	siw.Handler.End(c, id)
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

	router.GET(options.BaseURL+"/discounts", wrapper.GetList)
	router.POST(options.BaseURL+"/discounts", wrapper.Create)
	router.PUT(options.BaseURL+"/discounts", wrapper.Update)
	router.GET(options.BaseURL+"/discounts/:id", wrapper.Get)
	router.PATCH(options.BaseURL+"/discounts/:id/end", wrapper.End)
}
