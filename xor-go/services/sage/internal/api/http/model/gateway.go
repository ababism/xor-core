package model

type PassSecureResourceRequest struct {
	AccessToken string         `json:"access_token"`
	Resource    string         `json:"resource" binding:"required"`
	Route       string         `json:"route" binding:"required"`
	Method      string         `json:"method" binding:"required"`
	Body        map[string]any `json:"body" binding:"required"`
}

type PassInsecureResourceRequest struct {
	Resource string         `json:"resource" binding:"required"`
	Route    string         `json:"route" binding:"required"`
	Method   string         `json:"method" binding:"required"`
	Body     map[string]any `json:"body" binding:"required"`
}

type PassResourceResponse struct {
	Status int `json:"status"`
	Body   any `json:"body"`
}
