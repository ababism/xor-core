package domain

import (
	"github.com/google/uuid"
	"time"
)

type PayoutRequestData struct {
}

type PayoutRequestGet struct {
	UUID       uuid.UUID
	Receiver   uuid.UUID
	Amount     float32
	Data       PayoutRequestData
	ReceivedAt time.Time
}

type PayoutRequestCreate struct {
	Receiver   uuid.UUID
	Amount     float32
	Data       PayoutRequestData
	ReceivedAt time.Time
}

type PayoutRequestFilter struct {
	UUID       *uuid.UUID
	Receiver   *uuid.UUID
	Amount     *float32
	ReceivedAt *time.Time
}
