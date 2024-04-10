// Package bank provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.0.0 DO NOT EDIT.
package bank

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/oapi-codegen/runtime"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// BankAccountCreate defines model for BankAccountCreate.
type BankAccountCreate struct {
	AccountUUID openapi_types.UUID   `json:"AccountUUID"`
	Data        BankAccountData      `json:"Data"`
	Login       string               `json:"Login"`
	Payments    []openapi_types.UUID `json:"Payments"`
}

// BankAccountData defines model for BankAccountData.
type BankAccountData = map[string]interface{}

// BankAccountFilter defines model for BankAccountFilter.
type BankAccountFilter struct {
	AccountUUID *openapi_types.UUID `json:"AccountUUID,omitempty"`
	Funds       *float32            `json:"Funds,omitempty"`
	Login       *string             `json:"Login,omitempty"`
	Status      *string             `json:"Status,omitempty"`
	UUID        *openapi_types.UUID `json:"UUID,omitempty"`
}

// BankAccountGet defines model for BankAccountGet.
type BankAccountGet struct {
	AccountUUID  openapi_types.UUID   `json:"AccountUUID"`
	CreatedAt    time.Time            `json:"CreatedAt"`
	Data         BankAccountData      `json:"Data"`
	Funds        float32              `json:"Funds"`
	LastDealAt   *time.Time           `json:"LastDealAt,omitempty"`
	LastUpdateAt time.Time            `json:"LastUpdateAt"`
	Login        string               `json:"Login"`
	Payments     []openapi_types.UUID `json:"Payments"`
	Status       string               `json:"Status"`
	UUID         openapi_types.UUID   `json:"UUID"`
}

// BankAccountUpdate defines model for BankAccountUpdate.
type BankAccountUpdate struct {
	AccountUUID openapi_types.UUID   `json:"AccountUUID"`
	Data        BankAccountData      `json:"Data"`
	Funds       float32              `json:"Funds"`
	LastDealAt  *time.Time           `json:"LastDealAt,omitempty"`
	Login       string               `json:"Login"`
	Payments    []openapi_types.UUID `json:"Payments"`
	Status      string               `json:"Status"`
	UUID        openapi_types.UUID   `json:"UUID"`
}

// PutBankAccountsLoginChangeFundsParams defines parameters for PutBankAccountsLoginChangeFunds.
type PutBankAccountsLoginChangeFundsParams struct {
	NewFunds float32 `form:"newFunds" json:"newFunds"`
}

// GetBankAccountsJSONRequestBody defines body for GetBankAccounts for application/json ContentType.
type GetBankAccountsJSONRequestBody = BankAccountFilter

// PostBankAccountsJSONRequestBody defines body for PostBankAccounts for application/json ContentType.
type PostBankAccountsJSONRequestBody = BankAccountCreate

// PutBankAccountsJSONRequestBody defines body for PutBankAccounts for application/json ContentType.
type PutBankAccountsJSONRequestBody = BankAccountUpdate

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// List bank accounts
	// (GET /bank-accounts)
	GetBankAccounts(c *gin.Context)
	// Create a bank account
	// (POST /bank-accounts)
	PostBankAccounts(c *gin.Context)
	// Update a bank account
	// (PUT /bank-accounts)
	PutBankAccounts(c *gin.Context)
	// Get bank account by login
	// (GET /bank-accounts/{login})
	GetBankAccountsLogin(c *gin.Context, login string)
	// Change bank account funds
	// (PUT /bank-accounts/{login}/change-funds)
	PutBankAccountsLoginChangeFunds(c *gin.Context, login string, params PutBankAccountsLoginChangeFundsParams)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// GetBankAccounts operation middleware
func (siw *ServerInterfaceWrapper) GetBankAccounts(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetBankAccounts(c)
}

// PostBankAccounts operation middleware
func (siw *ServerInterfaceWrapper) PostBankAccounts(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostBankAccounts(c)
}

// PutBankAccounts operation middleware
func (siw *ServerInterfaceWrapper) PutBankAccounts(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PutBankAccounts(c)
}

// GetBankAccountsLogin operation middleware
func (siw *ServerInterfaceWrapper) GetBankAccountsLogin(c *gin.Context) {

	var err error

	// ------------- Path parameter "login" -------------
	var login string

	err = runtime.BindStyledParameter("simple", false, "login", c.Param("login"), &login)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter login: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetBankAccountsLogin(c, login)
}

// PutBankAccountsLoginChangeFunds operation middleware
func (siw *ServerInterfaceWrapper) PutBankAccountsLoginChangeFunds(c *gin.Context) {

	var err error

	// ------------- Path parameter "login" -------------
	var login string

	err = runtime.BindStyledParameter("simple", false, "login", c.Param("login"), &login)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter login: %w", err), http.StatusBadRequest)
		return
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params PutBankAccountsLoginChangeFundsParams

	// ------------- Required query parameter "newFunds" -------------

	if paramValue := c.Query("newFunds"); paramValue != "" {

	} else {
		siw.ErrorHandler(c, fmt.Errorf("Query argument newFunds is required, but not found"), http.StatusBadRequest)
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "newFunds", c.Request.URL.Query(), &params.NewFunds)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter newFunds: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PutBankAccountsLoginChangeFunds(c, login, params)
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

	router.GET(options.BaseURL+"/bank-accounts", wrapper.GetBankAccounts)
	router.POST(options.BaseURL+"/bank-accounts", wrapper.PostBankAccounts)
	router.PUT(options.BaseURL+"/bank-accounts", wrapper.PutBankAccounts)
	router.GET(options.BaseURL+"/bank-accounts/:login", wrapper.GetBankAccountsLogin)
	router.PUT(options.BaseURL+"/bank-accounts/:login/change-funds", wrapper.PutBankAccountsLoginChangeFunds)
}
