package domain

import (
	"github.com/google/uuid"
	"time"
)

type PurchaseRequestGet struct {
	UUID       uuid.UUID
	Sender     uuid.UUID
	Receiver   uuid.UUID
	Products   []uuid.UUID
	WebhookURL string
	ReceivedAt time.Time
}

type PurchaseRequestCreate struct {
	Sender     uuid.UUID
	Receiver   uuid.UUID
	Products   []uuid.UUID
	WebhookURL string
	ReceivedAt time.Time
}

type PurchaseRequestFilter struct {
	UUID       *uuid.UUID
	Sender     *uuid.UUID
	Receiver   *uuid.UUID
	WebhookURL *string
	ReceivedAt *time.Time
}
