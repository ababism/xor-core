package payout

import (
	"xor-go/services/finances/internal/domain"
)

func DataToDomain(data PayoutRequestData) domain.PayoutRequestData {
	return domain.PayoutRequestData{}
}

func DataToPayment(data domain.PayoutRequestData) PayoutRequestData {
	return PayoutRequestData{}
}

func CreateToDomain(create PayoutRequestCreate) domain.PayoutRequestCreate {
	return domain.PayoutRequestCreate{
		Receiver:   create.Receiver,
		Amount:     create.Amount,
		Data:       DataToDomain(create.Data),
		ReceivedAt: create.ReceivedAt,
	}
}

func DomainToGet(get domain.PayoutRequestGet) PayoutRequestGet {
	return PayoutRequestGet{
		UUID:       get.UUID,
		Receiver:   get.Receiver,
		Amount:     get.Amount,
		Data:       DataToPayment(get.Data),
		ReceivedAt: get.ReceivedAt,
	}
}

func FilterToDomain(filter *PayoutRequestFilter) *domain.PayoutRequestFilter {
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
