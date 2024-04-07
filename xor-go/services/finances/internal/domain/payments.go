package domain

import (
	"github.com/google/uuid"
)

type PaymentsGetStatus struct {
	UUID   uuid.UUID `json:"id"`
	Status string    `json:"status"`
}

type PaymentsCreatePurchaseProduct struct {
	Description string  `json:"description"`
	Quantity    int     `json:"quantity"`
	Money       float32 `json:"money"`
	Currency    string  `json:"currency"`
	PaymentMode string  `json:"payment_mode"`
}

type PaymentsCreatePurchase struct {
	PaymentUUID uuid.UUID                       `json:"payment_uuid"`
	PaymentName string                          `json:"payment_name"`
	Money       float32                         `json:"money"`
	Currency    string                          `json:"currency"`
	FullName    string                          `json:"full_name"`
	Phone       string                          `json:"phone"`
	Email       string                          `json:"email"`
	Products    []PaymentsCreatePurchaseProduct `json:"products"`
}

type PaymentsCreatePayoutCard struct {
	First6        string `json:"first6"`
	Last4         string `json:"last4"`
	CardType      string `json:"card_type"`
	IssuerCountry string `json:"issuer_country"`
	IssuerName    string `json:"issuer_name"`
}

type PaymentsCreatePayout struct {
	PaymentUUID string                   `json:"payment_uuid"`
	PaymentName string                   `json:"payment_name"`
	Money       string                   `json:"money"`
	Currency    string                   `json:"currency"`
	FullName    string                   `json:"full_name"`
	Phone       string                   `json:"phone"`
	Email       string                   `json:"email"`
	CardInfo    PaymentsCreatePayoutCard `json:"card_info"`
	IsTest      bool                     `json:"is_test"`
}
