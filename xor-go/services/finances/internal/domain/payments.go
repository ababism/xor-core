package domain

import (
	"github.com/google/uuid"
	openapitypes "github.com/oapi-codegen/runtime/types"
	"time"
)

type PaymentsGetStatus struct {
	UUID   uuid.UUID
	Status string
}

type CardInfo struct {
	CardType      *string
	First6        string
	IssuerCountry *string
	IssuerName    string
	Last4         string
}

type PaymentsCreatePurchaseProduct struct {
	Description string
	Quantity    int
	Money       float32
	Currency    string
	PaymentMode string
}

type PaymentsCreatePurchase struct {
	PaymentUUID uuid.UUID
	PaymentName string
	Money       float32
	Currency    string
	FullName    string
	Phone       string
	Email       string
	Products    []PaymentsCreatePurchaseProduct
}

type Amount struct {
	Currency *string
	Value    *float32
}

type PayoutDestination struct {
	Card *CardInfo
	Type *string
}

type CreatePayoutResponse struct {
	Amount            Amount
	CreatedAt         time.Time
	Description       string
	Id                openapitypes.UUID
	Metadata          *map[string]interface{}
	PayoutDestination PayoutDestination
	Status            string
	Test              bool
}

type PaymentsCreatePayoutCard struct {
	First6        string
	Last4         string
	CardType      string
	IssuerCountry string
	IssuerName    string
}

type PaymentsCreatePayout struct {
	PaymentUUID uuid.UUID
	PaymentName string
	Money       float32
	Currency    string
	FullName    string
	Phone       string
	Email       string
	CardInfo    PaymentsCreatePayoutCard
	IsTest      bool
}

type Confirmation struct {
	ConfirmationUrl *string
	ReturnUrl       *string
	Type            *string
}

type PaymentMethod struct {
	Id    *openapitypes.UUID
	Saved *bool
	Type  *string
}

type Recipient struct {
	AccountId *string
	GatewayId *string
}

type CreatePurchaseResponse struct {
	Amount        Amount
	Confirmation  Confirmation
	CreatedAt     time.Time
	Description   string
	Id            openapitypes.UUID
	Metadata      *map[string]interface{}
	Paid          bool
	PaymentMethod PaymentMethod
	Recipient     Recipient
	Refundable    bool
	Status        string
	Test          bool
}
