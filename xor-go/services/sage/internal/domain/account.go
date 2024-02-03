package domain

import (
	"github.com/google/uuid"
	"time"
)

type Account struct {
	Uuid         uuid.UUID
	Login        string
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Contacts     *AccountContacts
}

type AccountContacts struct {
	TgLogin string
}

type RegisterAccount struct {
	Login    string
	Password string
}
