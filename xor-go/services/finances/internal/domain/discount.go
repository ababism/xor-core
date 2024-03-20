package domain

import (
	"github.com/google/uuid"
	"time"
)

type DiscountGet struct {
	UUID         uuid.UUID
	CreatedBy    uuid.UUID
	Percent      float64
	StandAlone   bool
	StartedAt    time.Time
	EndedAt      time.Time
	Status       string
	CreatedAt    time.Time
	LastUpdateAt time.Time
}

type DiscountCreate struct {
	CreatedBy  uuid.UUID
	Percent    float64
	StandAlone bool
	StartedAt  time.Time
	EndedAt    time.Time
	Status     string
}

type DiscountUpdate struct {
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
	Status     *string
}
