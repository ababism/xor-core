package payment_api

import (
	"xor-go/services/finances/internal/domain"
	"xor-go/services/finances/internal/handler/generated/payment"
)

func DataToDomain(data payment.PaymentData) domain.PaymentData {
	return domain.PaymentData{}
}

func DataToPayment(data domain.PaymentData) payment.PaymentData {
	return payment.PaymentData{}
}

func CreateToDomain(create payment.PaymentCreate) domain.PaymentCreate {
	return domain.PaymentCreate{
		Sender:   create.Sender,
		Receiver: create.Receiver,
		Data:     DataToDomain(create.Data),
		URL:      create.URL,
		Status:   create.Status,
		EndedAt:  create.EndedAt,
	}
}

func DomainToGet(get domain.PaymentGet) payment.PaymentGet {
	return payment.PaymentGet{
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

func FilterToDomain(filter *payment.PaymentFilter) *domain.PaymentFilter {
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

func FilterToPayment(filter domain.PaymentFilter) payment.PaymentFilter {
	return payment.PaymentFilter{
		UUID:     filter.UUID,
		Sender:   filter.Sender,
		Receiver: filter.Receiver,
		URL:      filter.URL,
		Status:   filter.Status,
		EndedAt:  filter.EndedAt,
	}
}
