package dto

import (
	"github.com/google/uuid"
	"time"
)

type BankAccountData struct {
}

type BankAccount struct {
	UUID        uuid.UUID       `json:"uuid"`
	AccountUUID uuid.UUID       `json:"account_uuid"`
	Login       string          `json:"login"`
	Funds       float32         `json:"funds"`
	Data        BankAccountData `json:"data"`
	Status      string          `json:"status"`
	LastDealAt  time.Time       `json:"last_deal_at"`
	Payments    []uuid.UUID     `json:"payments"`
}

type PaymentData struct {
}

type Payment struct {
	UUID     uuid.UUID   `json:"uuid"`
	Sender   uuid.UUID   `json:"sender"`
	Receiver uuid.UUID   `json:"receiver"`
	Data     PaymentData `json:"data"`
	URL      string      `json:"url"`
	Status   string      `json:"status"`
	EndedAt  time.Time   `json:"ended_at"`
}

type Product struct {
	UUID      uuid.UUID `json:"uuid"`
	Name      string    `json:"name"`
	Price     float32   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

type Discount struct {
	UUID       uuid.UUID `json:"uuid"`
	CreatedBy  uuid.UUID `json:"created_by"`
	Percent    float32   `json:"percent"`
	StandAlone bool      `json:"stand_alone"`
	StartedAt  time.Time `json:"started_at"`
	EndedAt    time.Time `json:"ended_at"`
	Status     string    `json:"status"`
}

type PayoutRequestData struct {
}

type PayoutRequest struct {
	UUID      uuid.UUID         `json:"uuid"`
	Receiver  uuid.UUID         `json:"receiver"`
	Amount    float32           `json:"amount"`
	StartedAt time.Time         `json:"started_at"`
	Data      PayoutRequestData `json:"data"`
}

type PurchaseRequest struct {
	UUID       uuid.UUID `json:"uuid"`
	Sender     uuid.UUID `json:"sender"`
	Receiver   uuid.UUID `json:"receiver"`
	WebhookURL string    `json:"webhook_url"`
	StartedAt  time.Time `json:"started_at"`
}
