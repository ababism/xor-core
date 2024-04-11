package domain

import "github.com/google/uuid"

type IdmVerifyResponse struct {
	AccountUuid  uuid.UUID
	AccountEmail string
	Roles        []string
}

type PassSecureResourceInfo struct {
	AccessToken string
	Resource    string
	Route       string
}

type PassSecureResourceRequest struct {
	RequestUuid  uuid.UUID
	Resource     string
	Route        string
	Method       string
	Body         map[string]any
	AccountUuid  uuid.UUID
	AccountEmail string
	Roles        []string
}

type PassInsecureResourceRequest struct {
	RequestUuid uuid.UUID
	Resource    string
	Route       string
	Method      string
	Body        map[string]any
}

type InternalResourceResponse struct {
	Status int
	Body   any
}
