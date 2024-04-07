package adapters

import (
	"context"
	"github.com/google/uuid"
	"xor-go/services/finances/internal/domain"
)

type BankAccountRepository interface {
	Present(ctx context.Context, filter *domain.BankAccountFilter) (bool, error)
	Get(ctx context.Context, filter *domain.BankAccountFilter) (*domain.BankAccountGet, error)
	List(ctx context.Context, filter *domain.BankAccountFilter) ([]domain.BankAccountGet, error)
	Create(ctx context.Context, account *domain.BankAccountCreate) error
	Update(ctx context.Context, account *domain.BankAccountUpdate) error
}

type PaymentRepository interface {
	Get(ctx context.Context, filter *domain.PaymentFilter) (*domain.PaymentGet, error)
	List(ctx context.Context, filter *domain.PaymentFilter) ([]domain.PaymentGet, error)
	Create(ctx context.Context, payment *domain.PaymentCreate) error
}

type ProductRepository interface {
	Get(ctx context.Context, id uuid.UUID) (*domain.ProductGet, error)
	List(ctx context.Context, filter *domain.ProductFilter) ([]domain.ProductGet, error)
	Create(ctx context.Context, product *domain.ProductCreate) error
	Update(ctx context.Context, product *domain.ProductUpdate) error
	SetAvailability(ctx context.Context, id uuid.UUID, isAvailable bool) error
}

type DiscountRepository interface {
	Get(ctx context.Context, id uuid.UUID) (*domain.DiscountGet, error)
	List(ctx context.Context, filter *domain.DiscountFilter) ([]domain.DiscountGet, error)
	EndDiscount(ctx context.Context, id uuid.UUID) error
	Create(ctx context.Context, discount *domain.DiscountCreate) error
	Update(ctx context.Context, discount *domain.DiscountUpdate) error
}

type PayoutRequestRepository interface {
	Get(ctx context.Context, id uuid.UUID) (*domain.PayoutRequestGet, error)
	List(ctx context.Context, filter *domain.PayoutRequestFilter) ([]domain.PayoutRequestGet, error)
	Create(ctx context.Context, payout *domain.PayoutRequestCreate) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type PurchaseRequestRepository interface {
	Get(ctx context.Context, id uuid.UUID) (*domain.PurchaseRequestGet, error)
	List(ctx context.Context, filter *domain.PurchaseRequestFilter) ([]domain.PurchaseRequestGet, error)
	Create(ctx context.Context, purchase *domain.PurchaseRequestCreate) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type PaymentsClient interface {
	GetStatus(ctx context.Context, id uuid.UUID) (*domain.PaymentsGetStatus, error)
	CreatePurchase(ctx context.Context, purchase *domain.PaymentsCreatePurchase) error
	CreatePayout(ctx context.Context, purchase *domain.PaymentsCreatePayout) error
}
