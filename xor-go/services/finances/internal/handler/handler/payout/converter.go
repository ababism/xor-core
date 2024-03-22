package payout

import (
	"xor-go/services/finances/internal/domain"
	payoutrequest "xor-go/services/finances/internal/handler/generated/payout-request"
)

func DataToDomain(data payoutrequest.PayoutRequestData) domain.PayoutRequestData {
	return domain.PayoutRequestData{}
}

func DataToPayment(data domain.PayoutRequestData) payoutrequest.PayoutRequestData {
	return payoutrequest.PayoutRequestData{}
}

func CreateToDomain(create payoutrequest.PayoutRequestCreate) domain.PayoutRequestCreate {
	return domain.PayoutRequestCreate{
		Receiver:   create.Receiver,
		Amount:     create.Amount,
		Data:       DataToDomain(create.Data),
		ReceivedAt: create.ReceivedAt,
	}
}

func DomainToGet(get domain.PayoutRequestGet) payoutrequest.PayoutRequestGet {
	return payoutrequest.PayoutRequestGet{
		UUID:       get.UUID,
		Receiver:   get.Receiver,
		Amount:     get.Amount,
		Data:       DataToPayment(get.Data),
		ReceivedAt: get.ReceivedAt,
	}
}

func FilterToDomain(filter *payoutrequest.PayoutRequestFilter) *domain.PayoutRequestFilter {
	if filter == nil {
		return nil
	}
	return &domain.PayoutRequestFilter{
		UUID:       filter.UUID,
		Receiver:   filter.Receiver,
		Amount:     filter.Amount,
		ReceivedAt: filter.ReceivedAt,
	}
}
