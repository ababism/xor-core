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

	"github.com/oapi-codegen/runtime"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// Driver defines model for Driver.
type Driver struct {
	// Id Идентификатор водителя
	Id *string `json:"id,omitempty"`

	// Lat Latitude in decimal degrees
	Lat float32 `json:"lat"`

	// Lng Longitude in decimal degrees
	Lng float32 `json:"lng"`
}

// LatLngLiteral An object describing a specific location with Latitude and Longitude in decimal degrees.
type LatLngLiteral struct {
	// Lat Latitude in decimal degrees
	Lat float32 `json:"lat"`

	// Lng Longitude in decimal degrees
	Lng float32 `json:"lng"`
}

// GetDriversParams defines parameters for GetDrivers.
type GetDriversParams struct {
	// Lat Latitude in decimal degrees
	Lat float32 `form:"lat" json:"lat"`

	// Lng Longitude in decimal degrees
	Lng float32 `form:"lng" json:"lng"`

	// Radius Radius of serach
	Radius float32 `form:"radius" json:"radius"`
}

// UpdateDriverLocationJSONRequestBody defines body for UpdateDriverLocation for application/json ContentType.
type UpdateDriverLocationJSONRequestBody = LatLngLiteral

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
	// GetDrivers request
	GetDrivers(ctx context.Context, params *GetDriversParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateDriverLocationWithBody request with any body
	UpdateDriverLocationWithBody(ctx context.Context, driverId openapi_types.UUID, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateDriverLocation(ctx context.Context, driverId openapi_types.UUID, body UpdateDriverLocationJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetDrivers(ctx context.Context, params *GetDriversParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetDriversRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UpdateDriverLocationWithBody(ctx context.Context, driverId openapi_types.UUID, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUpdateDriverLocationRequestWithBody(c.Server, driverId, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UpdateDriverLocation(ctx context.Context, driverId openapi_types.UUID, body UpdateDriverLocationJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUpdateDriverLocationRequest(c.Server, driverId, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetDriversRequest generates requests for GetDrivers
func NewGetDriversRequest(server string, params *GetDriversParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/drivers")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "lat", runtime.ParamLocationQuery, params.Lat); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "lng", runtime.ParamLocationQuery, params.Lng); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "radius", runtime.ParamLocationQuery, params.Radius); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

		queryURL.RawQuery = queryValues.Encode()
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewUpdateDriverLocationRequest calls the generic UpdateDriverLocation builder with application/json body
func NewUpdateDriverLocationRequest(server string, driverId openapi_types.UUID, body UpdateDriverLocationJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewUpdateDriverLocationRequestWithBody(server, driverId, "application/json", bodyReader)
}

// NewUpdateDriverLocationRequestWithBody generates requests for UpdateDriverLocation with any type of body
func NewUpdateDriverLocationRequestWithBody(server string, driverId openapi_types.UUID, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "driver_id", runtime.ParamLocationPath, driverId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/drivers/%s/location", pathParam0)
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
	// GetDriversWithResponse request
	GetDriversWithResponse(ctx context.Context, params *GetDriversParams, reqEditors ...RequestEditorFn) (*GetDriversResponse, error)

	// UpdateDriverLocationWithBodyWithResponse request with any body
	UpdateDriverLocationWithBodyWithResponse(ctx context.Context, driverId openapi_types.UUID, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateDriverLocationResponse, error)

	UpdateDriverLocationWithResponse(ctx context.Context, driverId openapi_types.UUID, body UpdateDriverLocationJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateDriverLocationResponse, error)
}

type GetDriversResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r GetDriversResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetDriversResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type UpdateDriverLocationResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r UpdateDriverLocationResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r UpdateDriverLocationResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetDriversWithResponse request returning *GetDriversResponse
func (c *ClientWithResponses) GetDriversWithResponse(ctx context.Context, params *GetDriversParams, reqEditors ...RequestEditorFn) (*GetDriversResponse, error) {
	rsp, err := c.GetDrivers(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetDriversResponse(rsp)
}

// UpdateDriverLocationWithBodyWithResponse request with arbitrary body returning *UpdateDriverLocationResponse
func (c *ClientWithResponses) UpdateDriverLocationWithBodyWithResponse(ctx context.Context, driverId openapi_types.UUID, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateDriverLocationResponse, error) {
	rsp, err := c.UpdateDriverLocationWithBody(ctx, driverId, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUpdateDriverLocationResponse(rsp)
}

func (c *ClientWithResponses) UpdateDriverLocationWithResponse(ctx context.Context, driverId openapi_types.UUID, body UpdateDriverLocationJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateDriverLocationResponse, error) {
	rsp, err := c.UpdateDriverLocation(ctx, driverId, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUpdateDriverLocationResponse(rsp)
}

// ParseGetDriversResponse parses an HTTP response from a GetDriversWithResponse call
func ParseGetDriversResponse(rsp *http.Response) (*GetDriversResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetDriversResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// ParseUpdateDriverLocationResponse parses an HTTP response from a UpdateDriverLocationWithResponse call
func ParseUpdateDriverLocationResponse(rsp *http.Response) (*UpdateDriverLocationResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &UpdateDriverLocationResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}
