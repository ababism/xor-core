package adapter

import (
	"context"
	"github.com/google/uuid"
	"xor-go/services/sage/internal/domain"
)

type AccountRepository interface {
	LoginPresent(ctx context.Context, login string) (bool, error)
	Get(ctx context.Context, uuid uuid.UUID) (*domain.Account, error)
	Create(ctx context.Context, account *domain.Account) error
	UpdatePassword(ctx context.Context, uuid uuid.UUID, passwordHash string) error
}
