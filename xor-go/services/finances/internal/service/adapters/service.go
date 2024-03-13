package adapters

import (
	"context"
	"github.com/google/uuid"
	"xor-go/services/finances/internal/domain"
)

type BankAccountService interface {
	GetAccountById(ctx context.Context, id uuid.UUID) (*domain.BankAccount, error)
	GetAccountByLogin(ctx context.Context, login string) (*domain.BankAccount, error)
	CreateAccount(ctx context.Context, account *domain.BankAccount) error
	UpdateAccount(ctx context.Context, account *domain.BankAccount) error
	DeleteAccount(ctx context.Context, id uuid.UUID) error
}
