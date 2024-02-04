package adapter

import (
	"context"
	"github.com/google/uuid"
	"xor-go/services/sage/internal/domain"
)

type AccountService interface {
	List(ctx context.Context, filter *domain.AccountFilter) ([]domain.Account, error)
	Create(ctx context.Context, registerAccount *domain.RegisterAccount) error
	UpdatePassword(ctx context.Context, uuid uuid.UUID, password string) error
	Deactivate(ctx context.Context, uuid uuid.UUID) error
}
