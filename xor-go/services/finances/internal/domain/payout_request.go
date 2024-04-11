package domain

import (
	"github.com/google/uuid"
	"time"
)

type PayoutRequestData struct {
}

type PayoutRequestGet struct {
	UUID      uuid.UUID
	Receiver  uuid.UUID
	Status    string
	Amount    float32
	Data      PayoutRequestData
	CreatedAt time.Time
}

type PayoutRequestCreate struct {
	Receiver  uuid.UUID
	Amount    float32
	Data      PayoutRequestData
	CreatedAt time.Time
}

type PayoutRequestFilter struct {
	UUID      *uuid.UUID
	Receiver  *uuid.UUID
	Status    *string
	Amount    *float32
	CreatedAt *time.Time
}
