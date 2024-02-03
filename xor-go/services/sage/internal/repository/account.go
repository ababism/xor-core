package repository

import (
	"context"
	"github.com/google/uuid"
	"xor-go/services/sage/internal/model"
)

type AccountRepositoryI interface {
	LoginPresent(ctx context.Context, login string) (bool, error)
	GetPasswordHash(ctx context.Context, uuid uuid.UUID) (string, error)
	Create(ctx context.Context, account *model.AccountEntity) error
	UpdatePassword(ctx context.Context, uuid uuid.UUID, passwordHash string) error
}
