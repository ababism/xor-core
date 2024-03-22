package purchase

import (
	"xor-go/services/finances/internal/domain"
)

func CreateToDomain(create PurchaseRequestCreate) domain.PurchaseRequestCreate {
	return domain.PurchaseRequestCreate{
		Sender:     create.Sender,
		Receiver:   create.Receiver,
		WebhookURL: create.WebhookURL,
		ReceivedAt: create.ReceivedAt,
	}
}

func DomainToGet(get domain.PurchaseRequestGet) PurchaseRequestGet {
	return PurchaseRequestGet{
		UUID:       get.UUID,
		Sender:     get.Sender,
		Receiver:   get.Receiver,
		WebhookURL: get.WebhookURL,
		ReceivedAt: get.ReceivedAt,
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
		WebhookURL: filter.WebhookURL,
		ReceivedAt: filter.ReceivedAt,
	}
}
