package repo_models

import (
	"github.com/google/uuid"
	"time"
	"xor-go/services/finances/internal/domain"
)

type Product struct {
	UUID          uuid.UUID `db:"uuid"`
	Name          string    `db:"name"`
	Price         float32   `db:"price"`
	IsAvailable   bool      `db:"is_available"`
	CreatedAt     time.Time `db:"created_at"`
	LastUpdatedAt time.Time `db:"last_updated_at"`
}

func CreateToProductPostgres(model *domain.ProductCreate) *Product {
	id, _ := uuid.NewUUID()
	return &Product{
		UUID:          id,
		Name:          model.Name,
		Price:         model.Price,
		IsAvailable:   true,
		CreatedAt:     time.Now(),
		LastUpdatedAt: time.Now(),
	}
}

func UpdateToProductPostgres(model *domain.ProductUpdate) *Product {
	return &Product{
		UUID:          model.UUID,
		Name:          model.Name,
		IsAvailable:   model.IsAvailable,
		Price:         model.Price,
		LastUpdatedAt: time.Now(),
	}
}

func ToProductDomain(model *Product) *domain.ProductGet {
	return &domain.ProductGet{
		UUID:          model.UUID,
		Name:          model.Name,
		Price:         model.Price,
		IsAvailable:   model.IsAvailable,
		CreatedAt:     model.CreatedAt,
		LastUpdatedAt: model.LastUpdatedAt,
	}
}
