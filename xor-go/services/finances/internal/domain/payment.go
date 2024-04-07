package domain

import (
	"github.com/google/uuid"
	"time"
)

const (
	STATUS_COMPLETED = "completed"
	STATUS_FAILED    = "failed"
)

type PaymentData struct {
}

type PaymentGet struct {
	UUID      uuid.UUID
	Sender    uuid.UUID
	Receiver  uuid.UUID
	Data      PaymentData
	URL       string
	Status    string
	EndedAt   time.Time
	CreatedAt time.Time
}

type PaymentCreate struct {
	Sender   uuid.UUID
	Receiver uuid.UUID
	Data     PaymentData
	URL      string
	Status   string
	EndedAt  time.Time
}

type PaymentFilter struct {
	UUID     *uuid.UUID
	Sender   *uuid.UUID
	Receiver *uuid.UUID
	URL      *string
	Status   *string
	EndedAt  *time.Time
}
