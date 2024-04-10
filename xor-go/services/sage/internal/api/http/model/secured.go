package model

//type RequestInfo struct {
//	Resource string `json:"resource"`
//	Path     string `json:"path"`
//	Method   string `json:"method"`
//}

//type T struct {
//	RequestInfo struct {
//		ApiResource string `json:"api_resource"`
//		ApiUrl      string `json:"api_url"`
//		Method      string `json:"method"`
//		Body        struct {
//		} `json:"body"`
//	} `json:"request_info"`
//}

type SecuredAccessRequest struct {
	AccessToken string         `json:"access_token" binding:"required"`
	ApiResource string         `json:"api_resource" binding:"required"`
	ApiUrl      string         `json:"api_url" binding:"required"`
	Method      string         `json:"method" binding:"required"`
	Body        map[string]any `json:"body" binding:"required"`
}

type SecuredAccessResponse struct {
	Status int `json:"status"`
	Body   any `json:"body"`
}

func NewSecuredAccessResponse(status int, body any) *SecuredAccessResponse {
	return &SecuredAccessResponse{Status: status, Body: body}
}

//type PostRequest struct {
//	RequestInfo *RequestInfo `json:"request_info"`
//}

//type PayoutRequestData struct {
//}
//
//type PayoutRequest struct {
//	UUID      uuid.UUID         `json:"uuid"`
//	Receiver  uuid.UUID         `json:"receiver"`
//	Amount    float32           `json:"amount"`
//	StartedAt time.Time         `json:"started_at"`
//	Data      PayoutRequestData `json:"data"`
//}
