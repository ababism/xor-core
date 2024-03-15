package adapters

import (
	"context"
	"github.com/google/uuid"
	"xor-go/services/finances/internal/domain"
)

type BankAccountRepository interface {
	Present(ctx context.Context, filter domain.BankAccountFilter) (bool, error)
	Get(ctx context.Context, filter domain.BankAccountFilter) (*domain.BankAccount, error)
	List(ctx context.Context, filter domain.BankAccountFilter) ([]domain.BankAccount, error)
	Create(ctx context.Context, account *domain.BankAccount) error
	Update(ctx context.Context, account *domain.BankAccount) error
	//? Deactivate(ctx context.Context, id uuid.UUID) error
}

type PaymentRepository interface {
	Get(ctx context.Context, filter domain.PaymentFilter) (*domain.Payment, error)
	List(ctx context.Context, filter domain.PaymentFilter) ([]domain.Payment, error)
	Create(ctx context.Context, payment *domain.Payment) error
	Update(ctx context.Context, payment *domain.Payment) error
}

type ProductRepository interface {
	Get(ctx context.Context, id uuid.UUID) (*domain.Product, error)
	List(ctx context.Context, filter domain.ProductFilter) ([]domain.Product, error)
	Create(ctx context.Context, product *domain.Product) error
	Update(ctx context.Context, product *domain.Product) error
	Disable(ctx context.Context, id uuid.UUID) error
}

type DiscountRepository interface {
	Get(ctx context.Context, id uuid.UUID) (*domain.Discount, error)
	List(ctx context.Context, filter domain.DiscountFilter) ([]domain.Discount, error)
	EndDiscount(ctx context.Context, id uuid.UUID) error
	Create(ctx context.Context, discount *domain.Discount) error
	Update(ctx context.Context, discount *domain.Discount) error
}

type PayoutRequestRepository interface {
	Get(ctx context.Context, id uuid.UUID) (*domain.PayoutRequest, error)
	List(ctx context.Context, filter domain.PayoutRequestFilter) ([]domain.PayoutRequest, error)
	Create(ctx context.Context, payout *domain.PayoutRequest) error
	Update(ctx context.Context, payout *domain.PayoutRequest) error
	Archive(ctx context.Context, id uuid.UUID) error
}

type PurchaseRequestRepository interface {
	Get(ctx context.Context, id uuid.UUID) (*domain.PurchaseRequest, error)
	List(ctx context.Context, filter domain.PurchaseRequestFilter) ([]domain.PurchaseRequest, error)
	Create(ctx context.Context, purchase *domain.PurchaseRequest) error
	Update(ctx context.Context, purchase *domain.PurchaseRequest) error
	Delete(ctx context.Context, id uuid.UUID) error
}
