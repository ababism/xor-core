package repository

import (
	"context"
	"github.com/google/uuid"
	"xor-go/services/sage/internal/domain"
)

type AccountRepository interface {
	LoginPresent(ctx context.Context, login string) (bool, error)
	GetPasswordHash(ctx context.Context, uuid uuid.UUID) (string, error)
	Create(ctx context.Context, account *domain.Account) error
	UpdatePassword(ctx context.Context, uuid uuid.UUID, passwordHash string) error
}
