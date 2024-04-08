package product

import (
	"xor-go/services/finances/internal/domain"
)

func DomainToGet(get domain.ProductGet) ProductGet {
	return ProductGet{
		UUID:          get.UUID,
		Name:          get.Name,
		Price:         get.Price,
		IsAvailable:   get.IsAvailable,
		CreatedAt:     get.CreatedAt,
		LastUpdatedAt: get.LastUpdatedAt,
	}
}

func CreateToDomain(create ProductCreate) domain.ProductCreate {
	return domain.ProductCreate{
		Name:        create.Name,
		Price:       create.Price,
		IsAvailable: create.IsAvailable,
	}
}

func UpdateToDomain(update ProductUpdate) domain.ProductUpdate {
	return domain.ProductUpdate{
		UUID:        update.UUID,
		Name:        update.Name,
		Price:       update.Price,
		IsAvailable: update.IsAvailable,
	}
}

func FilterToDomain(filter *ProductFilter) *domain.ProductFilter {
	if filter == nil {
		return nil
	}
	return &domain.ProductFilter{
		UUID:        filter.UUID,
		Name:        filter.Name,
		Price:       filter.Price,
		IsAvailable: filter.IsAvailable,
	}
}
