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
	Sender    *uuid.UUID
	Receiver  *uuid.UUID
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
	CreatedBy  *uuid.UUID
	PercentMin *float64
	PercentMax *float64
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
	Receiver  uuid.UUID
	AmountMin float64
	AmountMax float64
	StartedAt time.Time
	EndedAt   time.Time
}
type PurchaseRequest struct {
	UUID       uuid.UUID
	Sender     uuid.UUID
	Receiver   uuid.UUID
	WebhookURL string
	ReceivedAt time.Time
}

type PurchaseRequestFilter struct {
	Sender    uuid.UUID
	Receiver  uuid.UUID
	StartedAt time.Time
	EndedAt   time.Time
}
