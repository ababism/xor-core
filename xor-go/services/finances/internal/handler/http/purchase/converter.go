package purchase

import (
	"xor-go/services/finances/internal/domain"
)

func CreateToDomain(create PurchaseRequestCreate) domain.PurchaseRequestCreate {
	return domain.PurchaseRequestCreate{
		Sender:     create.Sender,
		Receiver:   create.Receiver,
		Status:     create.Status,
		Products:   create.Products,
		WebhookURL: create.WebhookURL,
		CreatedAt:  create.CreatedAt,
	}
}

func DomainToGet(get domain.PurchaseRequestGet) PurchaseRequestGet {
	return PurchaseRequestGet{
		UUID:       get.UUID,
		Sender:     get.Sender,
		Products:   get.Products,
		Status:     get.Status,
		Receiver:   get.Receiver,
		WebhookURL: get.WebhookURL,
		CreatedAt:  get.CreatedAt,
	}
}

func FilterToDomain(filter *PurchaseRequestFilter) *domain.PurchaseRequestFilter {
	if filter == nil {
		return nil
	}
	return &domain.PurchaseRequestFilter{
		UUID:       filter.UUID,
		Sender:     filter.Sender,
		Receiver:   filter.Receiver,
		Status:     filter.Status,
		WebhookURL: filter.WebhookURL,
		CreatedAt:  filter.CreatedAt,
	}
}
