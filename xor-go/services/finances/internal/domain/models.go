package domain

import (
	"github.com/google/uuid"
	"time"
)

type BankAccountData struct {
}

type BankAccount struct {
	UUID        uuid.UUID
	AccountUUID uuid.UUID
	Login       string
	Funds       float64
	Data        BankAccountData
	Status      string
	LastDealAt  time.Time
	Payments    []uuid.UUID
}

type BankAccountFilter struct {
	UUID        *uuid.UUID
	AccountUUID *uuid.UUID
	Login       *string
	Funds       *float64
	Status      *string
	LastDealAt  *time.Time
}

type PaymentData struct {
}

type Payment struct {
	UUID     uuid.UUID
	Sender   uuid.UUID
	Receiver uuid.UUID
	Data     PaymentData
	URL      string
	Status   string
	EndedAt  time.Time
}

type PaymentFilter struct {
	UUID      *uuid.UUID
	Sender    *uuid.UUID
	Receiver  *uuid.UUID
	URL       *string
	Status    *string
	StartedAt *time.Time
	EndedAt   *time.Time
}

type Product struct {
	UUID      uuid.UUID
	Name      string
	Price     float64
	CreatedAt time.Time
}

type ProductFilter struct {
	UUID      *uuid.UUID
	Price     *float64
	StartedAt *time.Time
	EndedAt   *time.Time
}

type Discount struct {
	UUID       uuid.UUID
	CreatedBy  uuid.UUID
	Percent    float64
	StandAlone bool
	StartedAt  time.Time
	EndedAt    time.Time
	Status     string
}

type DiscountFilter struct {
	UUID       *uuid.UUID
	CreatedBy  *uuid.UUID
	Percent    *float64
	StandAlone *bool
	StartedAt  *time.Time
	EndedAt    *time.Time
	Status     *string
}

type PayoutRequestData struct {
}

type PayoutRequest struct {
	UUID       uuid.UUID
	Receiver   uuid.UUID
	Amount     float64
	ReceivedAt time.Time
	Data       PayoutRequestData
}

type PayoutRequestFilter struct {
	UUID       *uuid.UUID
	Receiver   *uuid.UUID
	Amount     *float64
	ReceivedAt *time.Time
}
type PurchaseRequest struct {
	UUID       uuid.UUID
	Sender     uuid.UUID
	Receiver   uuid.UUID
	WebhookURL string
	ReceivedAt time.Time
}

type PurchaseRequestFilter struct {
	UUID       uuid.UUID
	Sender     *uuid.UUID
	Receiver   *uuid.UUID
	WebhookURL *string
	ReceivedAt *time.Time
}
