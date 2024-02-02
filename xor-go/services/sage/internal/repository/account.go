package repository

import (
	"context"
	"github.com/google/uuid"
	"xor-go/services/sage/internal/models"
)

type AccountRepositoryI interface {
	LoginPresent(ctx context.Context, login string) (bool, error)
	GetPasswordHash(ctx context.Context, uuid uuid.UUID) (string, error)
	Create(ctx context.Context, account *models.AccountEntity) error
	UpdatePassword(ctx context.Context, uuid uuid.UUID, passwordHash string) error
}
