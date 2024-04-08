package adapters

import (
	"context"
	"xor-go/services/courses/internal/domain"
)

type FinancesClient interface {
	RegisterProducts(ctx context.Context, products []domain.Product) ([]domain.Product, error)
	CreatePurchase(ctx context.Context, productIDs []domain.Product) (domain.PaymentRedirect, error)
}

type KafkaClient interface {
	SendMessage(ctx context.Context, message string) error
}
