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
	Amount     float64
	ReceivedAt time.Time
	Data       PayoutRequestData
}

type PayoutRequestCreate struct {
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
