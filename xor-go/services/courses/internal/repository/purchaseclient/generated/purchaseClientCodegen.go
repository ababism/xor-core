// Package generated provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.0.0 DO NOT EDIT.
package generated

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

	"github.com/oapi-codegen/runtime"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// PurchaseRequestCreate defines model for PurchaseRequestCreate.
type PurchaseRequestCreate struct {
	CreatedAt  time.Time            `json:"CreatedAt"`
	Products   []openapi_types.UUID `json:"Products"`
	Receiver   *openapi_types.UUID  `json:"Receiver,omitempty"`
	Sender     *openapi_types.UUID  `json:"Sender,omitempty"`
	WebhookURL string               `json:"WebhookURL"`
}

// PurchaseRequestFilter defines model for PurchaseRequestFilter.
type PurchaseRequestFilter struct {
	CreatedAt  *time.Time          `json:"CreatedAt,omitempty"`
	Receiver   *openapi_types.UUID `json:"Receiver,omitempty"`
	Sender     *openapi_types.UUID `json:"Sender,omitempty"`
	UUID       *openapi_types.UUID `json:"UUID,omitempty"`
	WebhookURL *string             `json:"WebhookURL,omitempty"`
}

// PurchaseRequestGet defines model for PurchaseRequestGet.
type PurchaseRequestGet struct {
	CreatedAt  time.Time            `json:"CreatedAt"`
	Products   []openapi_types.UUID `json:"Products"`
	Receiver   *openapi_types.UUID  `json:"Receiver,omitempty"`
	Sender     *openapi_types.UUID  `json:"Sender,omitempty"`
	UUID       openapi_types.UUID   `json:"UUID"`
	WebhookURL string               `json:"WebhookURL"`
}

// GetListJSONRequestBody defines body for GetList for application/json ContentType.
type GetListJSONRequestBody = PurchaseRequestFilter

// CreateJSONRequestBody defines body for Create for application/json ContentType.
type CreateJSONRequestBody = PurchaseRequestCreate

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
	// GetListWithBody request with any body
	GetListWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	GetList(ctx context.Context, body GetListJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateWithBody request with any body
	CreateWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	Create(ctx context.Context, body CreateJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// Get request
	Get(ctx context.Context, id openapi_types.UUID, reqEditors ...RequestEditorFn) (*http.Response, error)

	// Archive request
	Archive(ctx context.Context, id openapi_types.UUID, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetListWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetListRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetList(ctx context.Context, body GetListJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetListRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) Create(ctx context.Context, body CreateJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) Get(ctx context.Context, id openapi_types.UUID, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) Archive(ctx context.Context, id openapi_types.UUID, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewArchiveRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetListRequest calls the generic GetList builder with application/json body
func NewGetListRequest(server string, body GetListJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewGetListRequestWithBody(server, "application/json", bodyReader)
}

// NewGetListRequestWithBody generates requests for GetList with any type of body
func NewGetListRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/purchase-requests")
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

// NewCreateRequest calls the generic Create builder with application/json body
func NewCreateRequest(server string, body CreateJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewCreateRequestWithBody(server, "application/json", bodyReader)
}

// NewCreateRequestWithBody generates requests for Create with any type of body
func NewCreateRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/purchase-requests")
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

// NewGetRequest generates requests for Get
func NewGetRequest(server string, id openapi_types.UUID) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/purchase-requests/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewArchiveRequest generates requests for Archive
func NewArchiveRequest(server string, id openapi_types.UUID) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/purchase-requests/%s/archive", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

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
	// GetListWithBodyWithResponse request with any body
	GetListWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*GetListResponse, error)

	GetListWithResponse(ctx context.Context, body GetListJSONRequestBody, reqEditors ...RequestEditorFn) (*GetListResponse, error)

	// CreateWithBodyWithResponse request with any body
	CreateWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateResponse, error)

	CreateWithResponse(ctx context.Context, body CreateJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateResponse, error)

	// GetWithResponse request
	GetWithResponse(ctx context.Context, id openapi_types.UUID, reqEditors ...RequestEditorFn) (*GetResponse, error)

	// ArchiveWithResponse request
	ArchiveWithResponse(ctx context.Context, id openapi_types.UUID, reqEditors ...RequestEditorFn) (*ArchiveResponse, error)
}

type GetListResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]PurchaseRequestGet
}

// Status returns HTTPResponse.Status
func (r GetListResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetListResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type CreateResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r CreateResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CreateResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *PurchaseRequestGet
}

// Status returns HTTPResponse.Status
func (r GetResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type ArchiveResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r ArchiveResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ArchiveResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetListWithBodyWithResponse request with arbitrary body returning *GetListResponse
func (c *ClientWithResponses) GetListWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*GetListResponse, error) {
	rsp, err := c.GetListWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetListResponse(rsp)
}

func (c *ClientWithResponses) GetListWithResponse(ctx context.Context, body GetListJSONRequestBody, reqEditors ...RequestEditorFn) (*GetListResponse, error) {
	rsp, err := c.GetList(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetListResponse(rsp)
}

// CreateWithBodyWithResponse request with arbitrary body returning *CreateResponse
func (c *ClientWithResponses) CreateWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateResponse, error) {
	rsp, err := c.CreateWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateResponse(rsp)
}

func (c *ClientWithResponses) CreateWithResponse(ctx context.Context, body CreateJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateResponse, error) {
	rsp, err := c.Create(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateResponse(rsp)
}

// GetWithResponse request returning *GetResponse
func (c *ClientWithResponses) GetWithResponse(ctx context.Context, id openapi_types.UUID, reqEditors ...RequestEditorFn) (*GetResponse, error) {
	rsp, err := c.Get(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetResponse(rsp)
}

// ArchiveWithResponse request returning *ArchiveResponse
func (c *ClientWithResponses) ArchiveWithResponse(ctx context.Context, id openapi_types.UUID, reqEditors ...RequestEditorFn) (*ArchiveResponse, error) {
	rsp, err := c.Archive(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseArchiveResponse(rsp)
}

// ParseGetListResponse parses an HTTP response from a GetListWithResponse call
func ParseGetListResponse(rsp *http.Response) (*GetListResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetListResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []PurchaseRequestGet
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseCreateResponse parses an HTTP response from a CreateWithResponse call
func ParseCreateResponse(rsp *http.Response) (*CreateResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CreateResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// ParseGetResponse parses an HTTP response from a GetWithResponse call
func ParseGetResponse(rsp *http.Response) (*GetResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest PurchaseRequestGet
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseArchiveResponse parses an HTTP response from a ArchiveWithResponse call
func ParseArchiveResponse(rsp *http.Response) (*ArchiveResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ArchiveResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}
