package domain

import (
	"github.com/google/uuid"
	"time"
)

type ProductGet struct {
	UUID          uuid.UUID
	Name          string
	Price         float32
	IsAvailable   bool
	CreatedAt     time.Time
	LastUpdatedAt time.Time
}

type ProductCreate struct {
	Name        string
	Price       float32
	IsAvailable bool
}

type ProductUpdate struct {
	UUID        uuid.UUID
	Name        string
	Price       float32
	IsAvailable bool
}

type ProductFilter struct {
	UUID        *uuid.UUID
	Name        *string
	Price       *float32
	IsAvailable *bool
}
