package idm

import "github.com/google/uuid"

type VerifyRequest struct {
	AccessToken string `json:"access_token"`
}

type VerifyResponse struct {
	AccountUuid  uuid.UUID `json:"uuid"`
	AccountEmail string    `json:"email"`
	Roles        []string  `json:"roles"`
}
