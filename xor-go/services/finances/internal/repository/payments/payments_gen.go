// Package payments provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.0.0 DO NOT EDIT.
package payments

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

// CardInfo defines model for CardInfo.
type CardInfo struct {
	CardType      *string `json:"card_type,omitempty"`
	First6        string  `json:"first6"`
	IssuerCountry *string `json:"issuer_country,omitempty"`
	IssuerName    string  `json:"issuer_name"`
	Last4         string  `json:"last4"`
}

// CreatePayout defines model for CreatePayout.
type CreatePayout struct {
	CardInfo    CardInfo           `json:"card_info"`
	Currency    *string            `json:"currency,omitempty"`
	Email       *string            `json:"email,omitempty"`
	FullName    string             `json:"full_name"`
	IsTest      *bool              `json:"is_test,omitempty"`
	Money       float32            `json:"money"`
	PaymentName string             `json:"payment_name"`
	PaymentUuid openapi_types.UUID `json:"payment_uuid"`
	Phone       *string            `json:"phone,omitempty"`
}

// CreatePayoutResponse defines model for CreatePayoutResponse.
type CreatePayoutResponse struct {
	Amount struct {
		Currency *string `json:"currency,omitempty"`
		Value    *string `json:"value,omitempty"`
	} `json:"amount"`
	CreatedAt         time.Time               `json:"created_at"`
	Description       string                  `json:"description"`
	Id                openapi_types.UUID      `json:"id"`
	Metadata          *map[string]interface{} `json:"metadata,omitempty"`
	PayoutDestination struct {
		Card *CardInfo `json:"card,omitempty"`
		Type *string   `json:"type,omitempty"`
	} `json:"payout_destination"`
	Status string `json:"status"`
	Test   bool   `json:"test"`
}

// CreatePurchase defines model for CreatePurchase.
type CreatePurchase struct {
	Currency    *string                 `json:"currency,omitempty"`
	Email       string                  `json:"email"`
	FullName    string                  `json:"full_name"`
	Money       float32                 `json:"money"`
	PaymentName string                  `json:"payment_name"`
	PaymentUuid string                  `json:"payment_uuid"`
	Phone       string                  `json:"phone"`
	Products    []CreatePurchaseProduct `json:"products"`
}

// CreatePurchaseProduct defines model for CreatePurchaseProduct.
type CreatePurchaseProduct struct {
	Currency    *string `json:"currency,omitempty"`
	Description string  `json:"description"`
	Money       float32 `json:"money"`
	PaymentMode *string `json:"payment_mode,omitempty"`
	Quantity    int     `json:"quantity"`
}

// CreatePurchaseResponse defines model for CreatePurchaseResponse.
type CreatePurchaseResponse struct {
	Amount struct {
		Currency *string `json:"currency,omitempty"`
		Value    *string `json:"value,omitempty"`
	} `json:"amount"`
	Confirmation struct {
		ConfirmationUrl *string `json:"confirmation_url,omitempty"`
		ReturnUrl       *string `json:"return_url,omitempty"`
		Type            *string `json:"type,omitempty"`
	} `json:"confirmation"`
	CreatedAt     time.Time               `json:"created_at"`
	Description   string                  `json:"description"`
	Id            openapi_types.UUID      `json:"id"`
	Metadata      *map[string]interface{} `json:"metadata,omitempty"`
	Paid          bool                    `json:"paid"`
	PaymentMethod struct {
		Id    *openapi_types.UUID `json:"id,omitempty"`
		Saved *bool               `json:"saved,omitempty"`
		Type  *string             `json:"type,omitempty"`
	} `json:"payment_method"`
	Recipient struct {
		AccountId *string `json:"account_id,omitempty"`
		GatewayId *string `json:"gateway_id,omitempty"`
	} `json:"recipient"`
	Refundable bool   `json:"refundable"`
	Status     string `json:"status"`
	Test       bool   `json:"test"`
}

// GetStatusPayload defines model for GetStatusPayload.
type GetStatusPayload struct {
	PaymentId openapi_types.UUID `json:"payment_id"`
}

// GetStatusResult defines model for GetStatusResult.
type GetStatusResult struct {
	Id     *openapi_types.UUID `json:"id,omitempty"`
	Status *string             `json:"status,omitempty"`
}

// CreatePayoutRequestJSONRequestBody defines body for CreatePayoutRequest for application/json ContentType.
type CreatePayoutRequestJSONRequestBody = CreatePayout

// CreatePurchaseRequestJSONRequestBody defines body for CreatePurchaseRequest for application/json ContentType.
type CreatePurchaseRequestJSONRequestBody = CreatePurchase

// GetStatusJSONRequestBody defines body for GetStatus for application/json ContentType.
type GetStatusJSONRequestBody = GetStatusPayload

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// CreatePayoutRequestWithBody request with any body
	CreatePayoutRequestWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreatePayoutRequest(ctx context.Context, body CreatePayoutRequestJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreatePurchaseRequestWithBody request with any body
	CreatePurchaseRequestWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreatePurchaseRequest(ctx context.Context, body CreatePurchaseRequestJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetStatusWithBody request with any body
	GetStatusWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	GetStatus(ctx context.Context, body GetStatusJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) CreatePayoutRequestWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreatePayoutRequestRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreatePayoutRequest(ctx context.Context, body CreatePayoutRequestJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreatePayoutRequestRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreatePurchaseRequestWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreatePurchaseRequestRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreatePurchaseRequest(ctx context.Context, body CreatePurchaseRequestJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreatePurchaseRequestRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetStatusWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetStatusRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetStatus(ctx context.Context, body GetStatusJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetStatusRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewCreatePayoutRequestRequest calls the generic CreatePayoutRequest builder with application/json body
func NewCreatePayoutRequestRequest(server string, body CreatePayoutRequestJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewCreatePayoutRequestRequestWithBody(server, "application/json", bodyReader)
}

// NewCreatePayoutRequestRequestWithBody generates requests for CreatePayoutRequest with any type of body
func NewCreatePayoutRequestRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/payout")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewCreatePurchaseRequestRequest calls the generic CreatePurchaseRequest builder with application/json body
func NewCreatePurchaseRequestRequest(server string, body CreatePurchaseRequestJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewCreatePurchaseRequestRequestWithBody(server, "application/json", bodyReader)
}

// NewCreatePurchaseRequestRequestWithBody generates requests for CreatePurchaseRequest with any type of body
func NewCreatePurchaseRequestRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/purchase")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewGetStatusRequest calls the generic GetStatus builder with application/json body
func NewGetStatusRequest(server string, body GetStatusJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewGetStatusRequestWithBody(server, "application/json", bodyReader)
}

// NewGetStatusRequestWithBody generates requests for GetStatus with any type of body
func NewGetStatusRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/status")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// CreatePayoutRequestWithBodyWithResponse request with any body
	CreatePayoutRequestWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreatePayoutRequestResponse, error)

	CreatePayoutRequestWithResponse(ctx context.Context, body CreatePayoutRequestJSONRequestBody, reqEditors ...RequestEditorFn) (*CreatePayoutRequestResponse, error)

	// CreatePurchaseRequestWithBodyWithResponse request with any body
	CreatePurchaseRequestWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreatePurchaseRequestResponse, error)

	CreatePurchaseRequestWithResponse(ctx context.Context, body CreatePurchaseRequestJSONRequestBody, reqEditors ...RequestEditorFn) (*CreatePurchaseRequestResponse, error)

	// GetStatusWithBodyWithResponse request with any body
	GetStatusWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*GetStatusResponse, error)

	GetStatusWithResponse(ctx context.Context, body GetStatusJSONRequestBody, reqEditors ...RequestEditorFn) (*GetStatusResponse, error)
}

type CreatePayoutRequestResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *CreatePayoutResponse
}

// Status returns HTTPResponse.Status
func (r CreatePayoutRequestResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CreatePayoutRequestResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type CreatePurchaseRequestResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *CreatePurchaseResponse
}

// Status returns HTTPResponse.Status
func (r CreatePurchaseRequestResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CreatePurchaseRequestResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetStatusResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *GetStatusResult
}

// Status returns HTTPResponse.Status
func (r GetStatusResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetStatusResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// CreatePayoutRequestWithBodyWithResponse request with arbitrary body returning *CreatePayoutRequestResponse
func (c *ClientWithResponses) CreatePayoutRequestWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreatePayoutRequestResponse, error) {
	rsp, err := c.CreatePayoutRequestWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreatePayoutRequestResponse(rsp)
}

func (c *ClientWithResponses) CreatePayoutRequestWithResponse(ctx context.Context, body CreatePayoutRequestJSONRequestBody, reqEditors ...RequestEditorFn) (*CreatePayoutRequestResponse, error) {
	rsp, err := c.CreatePayoutRequest(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreatePayoutRequestResponse(rsp)
}

// CreatePurchaseRequestWithBodyWithResponse request with arbitrary body returning *CreatePurchaseRequestResponse
func (c *ClientWithResponses) CreatePurchaseRequestWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreatePurchaseRequestResponse, error) {
	rsp, err := c.CreatePurchaseRequestWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreatePurchaseRequestResponse(rsp)
}

func (c *ClientWithResponses) CreatePurchaseRequestWithResponse(ctx context.Context, body CreatePurchaseRequestJSONRequestBody, reqEditors ...RequestEditorFn) (*CreatePurchaseRequestResponse, error) {
	rsp, err := c.CreatePurchaseRequest(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreatePurchaseRequestResponse(rsp)
}

// GetStatusWithBodyWithResponse request with arbitrary body returning *GetStatusResponse
func (c *ClientWithResponses) GetStatusWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*GetStatusResponse, error) {
	rsp, err := c.GetStatusWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetStatusResponse(rsp)
}

func (c *ClientWithResponses) GetStatusWithResponse(ctx context.Context, body GetStatusJSONRequestBody, reqEditors ...RequestEditorFn) (*GetStatusResponse, error) {
	rsp, err := c.GetStatus(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetStatusResponse(rsp)
}

// ParseCreatePayoutRequestResponse parses an HTTP response from a CreatePayoutRequestWithResponse call
func ParseCreatePayoutRequestResponse(rsp *http.Response) (*CreatePayoutRequestResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CreatePayoutRequestResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest CreatePayoutResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseCreatePurchaseRequestResponse parses an HTTP response from a CreatePurchaseRequestWithResponse call
func ParseCreatePurchaseRequestResponse(rsp *http.Response) (*CreatePurchaseRequestResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CreatePurchaseRequestResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest CreatePurchaseResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseGetStatusResponse parses an HTTP response from a GetStatusWithResponse call
func ParseGetStatusResponse(rsp *http.Response) (*GetStatusResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetStatusResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest GetStatusResult
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
