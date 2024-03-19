package domain

import (
	"github.com/google/uuid"
	"time"
)

type BankAccountData struct {
}

type BankAccountGet struct {
	UUID         uuid.UUID
	AccountUUID  uuid.UUID
	Login        string
	Funds        float64
	Data         BankAccountData
	Status       string
	LastDealAt   time.Time
	CreatedAt    time.Time
	LastUpdateAt *time.Time
	Payments     []uuid.UUID
}

type BankAccountCreate struct {
	AccountUUID uuid.UUID
	Login       string
	Data        BankAccountData
	Payments    []uuid.UUID
}

type BankAccountPost struct {
	UUID        uuid.UUID
	AccountUUID uuid.UUID
	Login       string
	Funds       float64
	Data        BankAccountData
	Status      string
	LastDealAt  time.Time
	Payments    []uuid.UUID
}

type BankAccountFilter struct {
	UUID        *uuid.UUID
	AccountUUID *uuid.UUID
	Login       *string
	Funds       *float64
	Status      *string
}
