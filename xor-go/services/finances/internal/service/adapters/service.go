package adapters

import (
	"context"
	"github.com/google/uuid"
	"xor-go/services/finances/internal/domain"
)

const (
	ServiceNameBankAccount     = "bank-account-service"
	ServiceNamePayment         = "payment-service"
	ServiceNamePaymentsClient  = "payments-client"
	ServiceNameProduct         = "product-service"
	ServiceNameDiscount        = "discount-service"
	ServiceNamePayoutRequest   = "payout-request-service"
	ServiceNamePurchaseRequest = "purchase-request-service"
)

type BankAccountService interface {
	GetByLogin(ctx context.Context, login string) (*domain.BankAccountGet, error)
	GetById(ctx context.Context, id uuid.UUID) (*domain.BankAccountGet, error)
	List(ctx context.Context, filter *domain.BankAccountFilter) ([]domain.BankAccountGet, error)
	Create(ctx context.Context, account *domain.BankAccountCreate) (*uuid.UUID, error)
	Update(ctx context.Context, account *domain.BankAccountUpdate) error
	ChangeFunds(ctx context.Context, id uuid.UUID, newFunds float32) error
}

type PaymentService interface {
	Get(ctx context.Context, uuid uuid.UUID) (*domain.PaymentGet, error)
	List(ctx context.Context, filter *domain.PaymentFilter) ([]domain.PaymentGet, error)
	Create(ctx context.Context, payment *domain.PaymentCreate) (*uuid.UUID, error)
}

type ProductService interface {
	Get(ctx context.Context, id uuid.UUID) (*domain.ProductGet, error)
	GetPrice(ctx context.Context, productUUIDs []uuid.UUID) (*float32, error)
	List(ctx context.Context, filter *domain.ProductFilter) ([]domain.ProductGet, error)
	Create(ctx context.Context, product *domain.ProductCreate) (*uuid.UUID, error)
	Update(ctx context.Context, product *domain.ProductUpdate) error
	SetAvailability(ctx context.Context, id uuid.UUID, isAvailable bool) error
}

type DiscountService interface {
	Get(ctx context.Context, id uuid.UUID) (*domain.DiscountGet, error)
	List(ctx context.Context, filter *domain.DiscountFilter) ([]domain.DiscountGet, error)
	EndDiscount(ctx context.Context, id uuid.UUID) error
	Create(ctx context.Context, discount *domain.DiscountCreate) (*uuid.UUID, error)
	Update(ctx context.Context, discount *domain.DiscountUpdate) error
}

type PayoutRequestService interface {
	Get(ctx context.Context, id uuid.UUID) (*domain.PayoutRequestGet, error)
	List(ctx context.Context, filter *domain.PayoutRequestFilter) ([]domain.PayoutRequestGet, error)
	Create(ctx context.Context, payout *domain.PayoutRequestCreate) (*uuid.UUID, error)
	Archive(ctx context.Context, id uuid.UUID) error
}

type PurchaseRequestService interface {
	Get(ctx context.Context, id uuid.UUID) (*domain.PurchaseRequestGet, error)
	List(ctx context.Context, filter *domain.PurchaseRequestFilter) ([]domain.PurchaseRequestGet, error)
	Create(ctx context.Context, purchase *domain.PurchaseRequestCreate) (*uuid.UUID, error)
	Archive(ctx context.Context, id uuid.UUID) error
}
