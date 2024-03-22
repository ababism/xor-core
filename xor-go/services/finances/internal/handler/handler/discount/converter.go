package discount

import (
	"xor-go/services/finances/internal/domain"
	"xor-go/services/finances/internal/handler/generated/discount"
)

func CreateToDomain(create discount.DiscountCreate) domain.DiscountCreate {
	return domain.DiscountCreate{
		CreatedBy:  create.CreatedBy,
		Percent:    create.Percent,
		StandAlone: create.StandAlone,
		StartedAt:  create.StartedAt,
		EndedAt:    create.EndedAt,
		Status:     create.Status,
	}
}

func DomainToGet(get domain.DiscountGet) discount.DiscountGet {
	return discount.DiscountGet{
		UUID:         get.UUID,
		CreatedBy:    get.CreatedBy,
		Percent:      get.Percent,
		StandAlone:   get.StandAlone,
		StartedAt:    get.StartedAt,
		EndedAt:      get.EndedAt,
		Status:       get.Status,
		CreatedAt:    get.CreatedAt,
		LastUpdateAt: get.LastUpdateAt,
	}
}

func UpdateToDomain(update discount.DiscountUpdate) domain.DiscountUpdate {
	return domain.DiscountUpdate{
		UUID:       update.UUID,
		CreatedBy:  update.CreatedBy,
		Percent:    update.Percent,
		StandAlone: update.StandAlone,
		StartedAt:  update.StartedAt,
		EndedAt:    update.EndedAt,
		Status:     update.Status,
	}
}

func FilterToDomain(filter *discount.DiscountFilter) *domain.DiscountFilter {
	if filter == nil {
		return nil
	}
	return &domain.DiscountFilter{
		UUID:       filter.UUID,
		CreatedBy:  filter.CreatedBy,
		Percent:    filter.Percent,
		StandAlone: filter.StandAlone,
		Status:     filter.Status,
	}
}
