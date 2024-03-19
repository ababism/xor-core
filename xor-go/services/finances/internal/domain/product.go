package domain

import (
	"github.com/google/uuid"
	"time"
)

type ProductGet struct {
	UUID          uuid.UUID
	Name          string
	Price         float64
	CreatedAt     time.Time
	LastUpdatedAt time.Time
}

type ProductCreate struct {
	Name  string
	Price float64
}

type ProductUpdate struct {
	UUID  uuid.UUID
	Name  string
	Price float64
}

type ProductFilter struct {
	UUID  *uuid.UUID
	Name  *string
	Price *float64
}
