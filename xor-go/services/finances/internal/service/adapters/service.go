package adapters

import (
	"context"
	"github.com/google/uuid"
	"xor-go/services/finances/internal/domain"
)

const (
	ServiceNameBankAccount     = "bank-account-service"
	ServiceNamePayment         = "payment-service"
	ServiceNameProduct         = "product-service"
	ServiceNameDiscount        = "discount-service"
	ServiceNamePayoutRequest   = "payout-request-service"
	ServiceNamePurchaseRequest = "purchase-request-service"
)

type BankAccountService interface {
	GetByLogin(ctx context.Context, login string) (*domain.BankAccountGet, error)
	List(ctx context.Context, filter *domain.BankAccountFilter) ([]domain.BankAccountGet, error)
	Create(ctx context.Context, account *domain.BankAccountCreate) error
	Update(ctx context.Context, account *domain.BankAccountUpdate) error
	AddDiffToFunds(ctx context.Context, login string, newFunds float64) error
}

type PaymentService interface {
	Get(ctx context.Context, uuid uuid.UUID) (*domain.PaymentGet, error)
	List(ctx context.Context, filter *domain.PaymentFilter) ([]domain.PaymentGet, error)
	Create(ctx context.Context, payment *domain.PaymentCreate) error
}

type ProductService interface {
	Get(ctx context.Context, id uuid.UUID) (*domain.ProductGet, error)
	List(ctx context.Context, filter *domain.ProductFilter) ([]domain.ProductGet, error)
	Create(ctx context.Context, product *domain.ProductCreate) error
	Update(ctx context.Context, product *domain.ProductUpdate) error
	Disable(ctx context.Context, id uuid.UUID) error
}

type DiscountService interface {
	Get(ctx context.Context, id uuid.UUID) (*domain.DiscountGet, error)
	List(ctx context.Context, filter *domain.DiscountFilter) ([]domain.DiscountGet, error)
	EndDiscount(ctx context.Context, id uuid.UUID) error
	Create(ctx context.Context, discount *domain.DiscountCreate) error
	Update(ctx context.Context, discount *domain.DiscountUpdate) error
}

type PayoutRequestService interface {
	Get(ctx context.Context, id uuid.UUID) (*domain.PayoutRequestGet, error)
	List(ctx context.Context, filter *domain.PayoutRequestFilter) ([]domain.PayoutRequestGet, error)
	Create(ctx context.Context, payout *domain.PayoutRequestCreate) error
	Archive(ctx context.Context, id uuid.UUID) error
}

type PurchaseRequestService interface {
	Get(ctx context.Context, id uuid.UUID) (*domain.PurchaseRequestGet, error)
	List(ctx context.Context, filter *domain.PurchaseRequestFilter) ([]domain.PurchaseRequestGet, error)
	Create(ctx context.Context, purchase *domain.PurchaseRequestCreate) error
	Archive(ctx context.Context, id uuid.UUID) error
}