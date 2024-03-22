package payment

import (
	"xor-go/services/finances/internal/domain"
)

func DataToDomain(data PaymentData) domain.PaymentData {
	return domain.PaymentData{}
}

func DataToPayment(data domain.PaymentData) PaymentData {
	return PaymentData{}
}

func CreateToDomain(create PaymentCreate) domain.PaymentCreate {
	return domain.PaymentCreate{
		Sender:   create.Sender,
		Receiver: create.Receiver,
		Data:     DataToDomain(create.Data),
		URL:      create.URL,
		Status:   create.Status,
		EndedAt:  create.EndedAt,
	}
}

func DomainToGet(get domain.PaymentGet) PaymentGet {
	return PaymentGet{
		UUID:      get.UUID,
		Sender:    get.Sender,
		Receiver:  get.Receiver,
		Data:      DataToPayment(get.Data),
		URL:       get.URL,
		Status:    get.Status,
		EndedAt:   get.EndedAt,
		CreatedAt: get.CreatedAt,
	}
}

func FilterToDomain(filter *PaymentFilter) *domain.PaymentFilter {
	if filter == nil {
		return nil
	}
	return &domain.PaymentFilter{
		UUID:     filter.UUID,
		Sender:   filter.Sender,
		Receiver: filter.Receiver,
		URL:      filter.URL,
		Status:   filter.Status,
		EndedAt:  filter.EndedAt,
	}
}

func FilterToPayment(filter domain.PaymentFilter) PaymentFilter {
	return PaymentFilter{
		UUID:     filter.UUID,
		Sender:   filter.Sender,
		Receiver: filter.Receiver,
		URL:      filter.URL,
		Status:   filter.Status,
		EndedAt:  filter.EndedAt,
	}
}
