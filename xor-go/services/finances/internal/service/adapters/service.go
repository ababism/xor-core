package adapters

import (
	"context"
	"github.com/google/uuid"
	"xor-go/services/finances/internal/domain"
)

const (
	ServiceNameBankAccount = "bank-account-service"
	ServiceNamePayment     = "payment-service"
	ServiceNameProduct     = "product-service"
	ServiceNameDiscount    = "discount-service"
)

type BankAccountService interface {
	GetByLogin(ctx context.Context, login string) (*domain.BankAccountGet, error)
	List(ctx context.Context, filter domain.BankAccountFilter) ([]domain.BankAccountGet, error)
	Create(ctx context.Context, account *domain.BankAccountGet) error
	Update(ctx context.Context, account *domain.BankAccountGet) error
	AddDiffToFunds(ctx context.Context, login string, newFunds float64) error
	//? UpdateStatus(ctx context.Context, account *domain.BankAccountGet) error
}

type PaymentService interface {
	Get(ctx context.Context, uuid uuid.UUID) (*domain.Payment, error)
	List(ctx context.Context, filter domain.PaymentFilter) ([]domain.Payment, error)
	Create(ctx context.Context, payment *domain.Payment) error
	Update(ctx context.Context, payment *domain.Payment) error
}

type ProductService interface {
	Get(ctx context.Context, id uuid.UUID) (*domain.Product, error)
	List(ctx context.Context, filter domain.ProductFilter) ([]domain.Product, error)
	Create(ctx context.Context, product *domain.Product) error
	Update(ctx context.Context, product *domain.Product) error
	Disable(ctx context.Context, id uuid.UUID) error
}

type DiscountService interface {
	Get(ctx context.Context, id uuid.UUID) (*domain.Discount, error)
	List(ctx context.Context, filter domain.DiscountFilter) ([]domain.Discount, error)
	EndDiscount(ctx context.Context, id uuid.UUID) error
	Create(ctx context.Context, discount *domain.Discount) error
	Update(ctx context.Context, discount *domain.Discount) error
}

type PayoutRequestService interface {
	Get(ctx context.Context, id uuid.UUID) (*domain.PayoutRequest, error)
	List(ctx context.Context, filter domain.PayoutRequestFilter) ([]domain.PayoutRequest, error)
	Create(ctx context.Context, payout *domain.PayoutRequest) error
	// TODO продумать как архивировать или
	Archive(ctx context.Context, id uuid.UUID) error
}

type PurchaseRequestService interface {
	Get(ctx context.Context, id uuid.UUID) (*domain.PurchaseRequest, error)
	List(ctx context.Context, filter domain.PurchaseRequestFilter) ([]domain.PurchaseRequest, error)
	Create(ctx context.Context, purchase *domain.PurchaseRequest) error
	Update(ctx context.Context, purchase *domain.PurchaseRequest) error
	Delete(ctx context.Context, id uuid.UUID) error
}
