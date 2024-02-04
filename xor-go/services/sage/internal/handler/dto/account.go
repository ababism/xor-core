package dto

import (
	"github.com/google/uuid"
)

type RegisterAccountDto struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AccountDto struct {
	Uuid     uuid.UUID    `json:"uuid"`
	Login    string       `json:"login"`
	Contacts *ContactsDto `json:"contacts"`
	Active   bool         `json:"active"`
}

type ContactsDto struct {
	TgLogin string `json:"tg_login"`
}
