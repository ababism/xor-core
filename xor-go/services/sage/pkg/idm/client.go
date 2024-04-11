package idm

import (
	"errors"
	"github.com/go-resty/resty/v2"
)

const verifyEndpoint = "/admin/account/verify"

type ClientInterface interface {
	Verify(request *VerifyRequest) (*VerifyResponse, error)
}

type Client struct {
	ApiHost     string
	restyClient *resty.Client
}

func NewIdmClient(ApiUrl string) *Client {
	return &Client{
		ApiHost:     ApiUrl,
		restyClient: resty.New(),
	}
}

func (r *Client) Verify(request *VerifyRequest) (*VerifyResponse, error) {
	var verifyResponse VerifyResponse

	restyResponse, err := r.restyClient.R().
		SetAuthToken(request.AccessToken).
		SetResult(&verifyResponse).
		Get(r.getRequestUrl(verifyEndpoint))

	if err != nil {
		return nil, err
	}

	if restyResponse.IsError() {
		return nil, errors.New(restyResponse.String())
	}

	return &verifyResponse, nil
}

func (r *Client) getRequestUrl(endpoint string) string {
	return r.ApiHost + endpoint
}
