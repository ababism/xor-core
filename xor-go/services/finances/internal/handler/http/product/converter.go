package product

import (
	"xor-go/services/finances/internal/domain"
	"xor-go/services/finances/internal/handler/generated/product"
)

func DomainToGet(get domain.ProductGet) product.ProductGet {
	return product.ProductGet{
		UUID:          get.UUID,
		Name:          get.Name,
		Price:         get.Price,
		IsAvailable:   get.IsAvailable,
		CreatedAt:     get.CreatedAt,
		LastUpdatedAt: get.LastUpdatedAt,
	}
}

func CreateToDomain(create product.ProductCreate) domain.ProductCreate {
	return domain.ProductCreate{
		Name:  create.Name,
		Price: create.Price,
	}
}

func UpdateToDomain(update product.ProductUpdate) domain.ProductUpdate {
	return domain.ProductUpdate{
		UUID:        update.UUID,
		Name:        update.Name,
		Price:       update.Price,
		IsAvailable: update.IsAvailable,
	}
}

func FilterToDomain(filter *product.ProductFilter) *domain.ProductFilter {
	if filter == nil {
		return nil
	}
	return &domain.ProductFilter{
		UUID:  filter.UUID,
		Name:  filter.Name,
		Price: filter.Price,
	}
}
