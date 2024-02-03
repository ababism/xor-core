package entity

import (
	"github.com/google/uuid"
	"time"
)

type AccountMongo struct {
	Uuid         uuid.UUID             `bson:"uuid"`
	Login        string                `bson:"login"`
	PasswordHash string                `bson:"password_hash"`
	CreatedAt    time.Time             `bson:"created_at"`
	UpdatedAt    time.Time             `bson:"updated_at"`
	Contacts     *AccountContactsMongo `bson:"contacts"`
}

type AccountContactsMongo struct {
	TgLogin string `bson:"tg_login"`
}
