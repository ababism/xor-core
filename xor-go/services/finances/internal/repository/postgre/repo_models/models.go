package repo_models

import (
	"github.com/google/uuid"
	"time"
)

type BankAccountData struct {
}

type BankAccount struct {
	UUID          uuid.UUID       `db:"uuid"`
	AccountUUID   uuid.UUID       `db:"account_uuid"`
	Login         string          `db:"login"`
	Funds         float64         `db:"funds"`
	Data          BankAccountData `db:"data"`
	Status        string          `db:"status"`
	LastDealAt    time.Time       `db:"last_deal_at"`
	CreatedAt     time.Time       `db:"created_at"`
	LastUpdatedAt time.Time       `db:"last_updated_at"`
}

type PaymentData struct {
}

type Payment struct {
	UUID          uuid.UUID   `db:"uuid"`
	Sender        uuid.UUID   `db:"sender"`
	Receiver      uuid.UUID   `db:"receiver"`
	Data          PaymentData `db:"data"`
	URL           string      `db:"url"`
	Status        string      `db:"status"`
	EndedAt       time.Time   `db:"ended_at"`
	CreatedAt     time.Time   `db:"created_at"`
	LastUpdatedAt time.Time   `db:"last_updated_at"`
}

type Product struct {
	UUID          uuid.UUID `db:"uuid"`
	Name          string    `db:"name"`
	Price         float64   `db:"price"`
	CreatedAt     time.Time `db:"created_at"`
	LastUpdatedAt time.Time `db:"last_updated_at"`
}

type Discount struct {
	UUID          uuid.UUID `db:"uuid"`
	CreatedBy     uuid.UUID `db:"created_by"`
	Percent       float64   `db:"percent"`
	StandAlone    bool      `db:"stand_alone"`
	StartedAt     time.Time `db:"started_at"`
	EndedAt       time.Time `db:"ended_at"`
	Status        string    `db:"status"`
	CreatedAt     time.Time `db:"created_at"`
	LastUpdatedAt time.Time `db:"last_updated_at"`
}

type PayoutRequestData struct {
}

type PayoutRequest struct {
	UUID          uuid.UUID         `db:"uuid"`
	Receiver      uuid.UUID         `db:"receiver"`
	Amount        float64           `db:"amount"`
	StartedAt     time.Time         `db:"started_at"`
	Data          PayoutRequestData `db:"data"`
	CreatedAt     time.Time         `db:"created_at"`
	LastUpdatedAt time.Time         `db:"last_updated_at"`
}

type PurchaseRequest struct {
	UUID          uuid.UUID `db:"uuid"`
	Sender        uuid.UUID `db:"sender"`
	Receiver      uuid.UUID `db:"receiver"`
	WebhookURL    string    `db:"webhook_url"`
	StartedAt     time.Time `db:"started_at"`
	CreatedAt     time.Time `db:"created_at"`
	LastUpdatedAt time.Time `db:"last_updated_at"`
}
