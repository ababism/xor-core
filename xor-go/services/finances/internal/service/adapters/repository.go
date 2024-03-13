package adapters

import (
	"context"
	"github.com/google/uuid"
	"xor-go/services/finances/internal/domain"
)

type BankAccountRepository interface {
	GetAccountById(ctx context.Context, id uuid.UUID) (*domain.BankAccount, error)
	GetAccountByLogin(ctx context.Context, login string) (*domain.BankAccount, error)
	CreateAccount(ctx context.Context, account *domain.BankAccount) error
	UpdateAccount(ctx context.Context, account *domain.BankAccount) error
	DeleteAccount(ctx context.Context, id uuid.UUID) error
}

type PaymentsRepository interface {
	GetPayment(ctx context.Context, id uuid.UUID) (*domain.Payment, error)
	GetPaymentByFilter(ctx context.Context, filter domain.PaymentFilter) ([]domain.Payment, error)
	CreatePayment(ctx context.Context, payment *domain.Payment) error
	UpdatePayment(ctx context.Context, payment *domain.Payment) error
	DeletePayment(ctx context.Context, id uuid.UUID) error
}

type ProductRepository interface {
	GetProductById(ctx context.Context, id uuid.UUID) (*domain.Product, error)
	GetProductByName(ctx context.Context, name string) (*domain.Product, error)
	GetProductByProductFilter(ctx context.Context, filter domain.ProductFilter) ([]domain.Product, error)
	CreateProduct(ctx context.Context, product *domain.Product) error
	UpdateProduct(ctx context.Context, product *domain.Product) error
	DeleteProduct(ctx context.Context, id uuid.UUID) error
}

type DiscountRepository interface {
	GetDiscountById(ctx context.Context, id uuid.UUID) (*domain.Discount, error)
	GetDiscountByName(ctx context.Context, name string) (*domain.Discount, error)
	GetDiscountByDiscountFilter(ctx context.Context, filter domain.DiscountFilter) ([]domain.Discount, error)
	EndDiscount(ctx context.Context, id uuid.UUID) error
	CreateDiscount(ctx context.Context, discount *domain.Discount) error
	UpdateDiscount(ctx context.Context, discount *domain.Discount) error
	DeleteDiscount(ctx context.Context, id uuid.UUID) error
}

type PayoutRequestRepository interface {
	GetPayoutRequestById(ctx context.Context, id uuid.UUID) (*domain.PayoutRequest, error)
	GetPayoutRequestByPayoutRequestFilter(ctx context.Context, filter domain.PayoutRequestFilter) ([]domain.PayoutRequest, error)
	CreatePayoutRequest(ctx context.Context, payout *domain.PayoutRequest) error
	UpdatePayoutRequest(ctx context.Context, payout *domain.PayoutRequest) error
	DeletePayoutRequest(ctx context.Context, id uuid.UUID) error
}

type PurchaseRequestRepository interface {
	GetPurchaseRequestById(ctx context.Context, id uuid.UUID) (*domain.PurchaseRequest, error)
	GetPurchaseRequestByPurchaseRequestFilter(ctx context.Context, filter domain.PurchaseRequestFilter) ([]domain.PurchaseRequest, error)
	CreatePurchaseRequest(ctx context.Context, purchase *domain.PurchaseRequest) error
	UpdatePurchaseRequest(ctx context.Context, purchase *domain.PurchaseRequest) error
	DeletePurchaseRequest(ctx context.Context, id uuid.UUID) error
}
