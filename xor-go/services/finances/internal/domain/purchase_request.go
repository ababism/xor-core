package domain

import (
	"github.com/google/uuid"
	"time"
)

type PurchaseRequestGet struct {
	UUID       uuid.UUID
	Sender     *uuid.UUID
	Receiver   *uuid.UUID
	Amount     float32
	Products   []uuid.UUID
	WebhookURL string
	CreatedAt  time.Time
}

type PurchaseRequestCreate struct {
	Sender     *uuid.UUID
	Receiver   *uuid.UUID
	Products   []uuid.UUID
	WebhookURL string
	CreatedAt  time.Time
}

type PurchaseRequestFilter struct {
	UUID       *uuid.UUID
	Sender     *uuid.UUID
	Receiver   *uuid.UUID
	WebhookURL *string
	CreatedAt  *time.Time
}

type PurchaseRequestWebhook struct {
	Sender   uuid.UUID      `json:"buyer_id"`
	Receiver uuid.UUID      `json:"seller_id"`
	Products []ProductSmall `json:"products"`
}
