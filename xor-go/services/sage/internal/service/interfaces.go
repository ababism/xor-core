package service

import (
	"context"
	"github.com/google/uuid"
	"xor-go/services/sage/internal/domain"
)

type AccountService interface {
	Create(ctx context.Context, registerAccount *domain.RegisterAccount) error
	UpdatePassword(ctx context.Context, uuid uuid.UUID, password string) error
}
