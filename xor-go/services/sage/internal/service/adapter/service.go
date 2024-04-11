package adapter

import (
	"context"
	"xor-go/services/sage/internal/domain"
)

type GatewayService interface {
	Verify(ctx context.Context, passSecureResourceInfo *domain.PassSecureResourceInfo) (*domain.IdmVerifyResponse, error)
	PassSecure(ctx context.Context, passResourceRequest *domain.PassSecureResourceRequest) (*domain.InternalResourceResponse, error)
	PassInsecure(ctx context.Context, PassInsecureResourceRequest *domain.PassInsecureResourceRequest) (*domain.InternalResourceResponse, error)
}
