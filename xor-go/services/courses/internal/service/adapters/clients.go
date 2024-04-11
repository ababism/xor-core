package adapters

import (
	"context"
	"github.com/google/uuid"
	"xor-go/services/courses/internal/domain"
)

type FinancesClient interface {
	RegisterProducts(ctx context.Context, products []domain.Product) ([]domain.Product, error)
}

type PurchaseClient interface {
	CreatePurchase(ctx context.Context, productIDs []domain.Product, buyerID, ownerID uuid.UUID) (domain.PaymentRedirect, error)
}

type KafkaClient interface {
	SendMessage(ctx context.Context, message string) error
}
