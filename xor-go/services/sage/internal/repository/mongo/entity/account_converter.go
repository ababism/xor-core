package entity

import (
	"xor-go/services/sage/internal/domain"
)

func ToAccountMongo(account *domain.Account) *AccountMongo {
	return &AccountMongo{
		Uuid:         account.Uuid,
		Login:        account.Login,
		PasswordHash: account.PasswordHash,
		CreatedAt:    account.CreatedAt,
		UpdatedAt:    account.UpdatedAt,
		Contacts:     ToAccountContactsMongo(account.Contacts),
	}
}

func ToAccountContactsMongo(contacts *domain.AccountContacts) *AccountContactsMongo {
	return &AccountContactsMongo{
		TgLogin: contacts.TgLogin,
	}
}
