package models

import (
	"github.com/google/uuid"
	"time"
)

type AccountEntity struct {
	Uuid         uuid.UUID              `bson:"uuid"`
	Login        string                 `bson:"login"`
	PasswordHash string                 `bson:"password_hash"`
	CreatedAt    time.Time              `bson:"created_at"`
	UpdatedAt    time.Time              `bson:"updated_at"`
	Contacts     *AccountContactsEntity `bson:"contacts"`
}

type AccountContactsEntity struct {
	TgLogin string `bson:"tg_login"`
}

type RegisterAccountDto struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterAccountEntity struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
