package idm

import (
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
)

type ClientInterface interface {
	Verify(request *VerifyRequest) (*VerifyResponse, error)
}

type Client struct {
	baseUrl string
	client  *resty.Client
}

func NewClient(baseURL string, client *resty.Client) *Client {
	return &Client{
		baseUrl: baseURL,
		client:  client,
	}
}

func (r *Client) Verify(request *VerifyRequest) (*VerifyResponse, error) {
	url := r.baseUrl + "/admin/account/verify"
	var targetResponse VerifyResponse

	resp, err := r.client.R().
		SetAuthToken(request.AccessToken).
		SetResult(&targetResponse).
		Get(url)
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, errors.New("client")
	}

	fmt.Println(resp.StatusCode())

	return &targetResponse, nil
}
