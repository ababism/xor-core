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
	if contacts == nil {
		return nil
	}
	return &AccountContactsMongo{
		TgLogin: contacts.TgLogin,
	}
}

func ToAccount(accountMongo *AccountMongo) *domain.Account {
	return &domain.Account{
		Uuid:         accountMongo.Uuid,
		Login:        accountMongo.Login,
		PasswordHash: accountMongo.PasswordHash,
		CreatedAt:    accountMongo.CreatedAt,
		UpdatedAt:    accountMongo.UpdatedAt,
		Contacts:     ToAccountContacts(accountMongo.Contacts),
	}
}

func ToAccountContacts(contactsMongo *AccountContactsMongo) *domain.AccountContacts {
	if contactsMongo == nil {
		return nil
	}
	return &domain.AccountContacts{
		TgLogin: contactsMongo.TgLogin,
	}
}
