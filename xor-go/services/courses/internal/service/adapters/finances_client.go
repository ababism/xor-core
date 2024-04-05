package adapters

import (
	"context"
	"github.com/google/uuid"

	"xor-go/services/courses/internal/domain"
)

type FinancesClient interface {
	RegisterProducts(ctx context.Context, lessonIDs []uuid.UUID, teacherID uuid.UUID, price float64) ([]domain.Product, error)
	CreatePurchase(ctx context.Context, productIDs []domain.Product) (domain.PaymentRedirect, error)
}
