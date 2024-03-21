package purchase_request_api

import (
	"xor-go/services/finances/internal/domain"
	purchaserequest "xor-go/services/finances/internal/handler/generated/purchase-request"
)

func CreateToDomain(create purchaserequest.PurchaseRequestCreate) domain.PurchaseRequestCreate {
	return domain.PurchaseRequestCreate{
		Sender:     create.Sender,
		Receiver:   create.Receiver,
		WebhookURL: create.WebhookURL,
		ReceivedAt: create.ReceivedAt,
	}
}

func DomainToGet(get domain.PurchaseRequestGet) purchaserequest.PurchaseRequestGet {
	return purchaserequest.PurchaseRequestGet{
		UUID:       get.UUID,
		Sender:     get.Sender,
		Receiver:   get.Receiver,
		WebhookURL: get.WebhookURL,
		ReceivedAt: get.ReceivedAt,
	}
}

func FilterToDomain(filter *purchaserequest.PurchaseRequestFilter) *domain.PurchaseRequestFilter {
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
