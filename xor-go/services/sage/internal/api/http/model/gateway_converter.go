package model

import "xor-go/services/sage/internal/domain"

func ToPassSecureResourceInfo(m *PassSecureResourceRequest) *domain.PassSecureResourceInfo {
	return &domain.PassSecureResourceInfo{
		AccessToken: m.AccessToken,
		Resource:    m.Resource,
		Route:       m.Route,
	}
}

func ToPassResourceResponse(m *domain.InternalResourceResponse) *PassResourceResponse {
	return &PassResourceResponse{
		Status: m.Status,
		Body:   m.Body,
	}
}
