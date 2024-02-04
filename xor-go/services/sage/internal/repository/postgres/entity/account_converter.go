package entity

import (
	"xor-go/services/sage/internal/domain"
)

func ToAccountPostgres(domain *domain.Account) *AccountPostgres {
	var contacts *ContactsPostgres
	if domain.Contacts != nil {
		contacts = ToContactsPostgres(domain.Contacts)
	}
	return &AccountPostgres{
		Uuid:         domain.Uuid,
		Login:        domain.Login,
		PasswordHash: domain.PasswordHash,
		Contacts:     contacts,
		Active:       domain.Active,
	}
}

func ToContactsPostgres(domain *domain.Contacts) *ContactsPostgres {
	return &ContactsPostgres{
		TgLogin: domain.TgLogin,
	}
}

func ToAccount(entity *AccountPostgres) *domain.Account {
	var contacts *domain.Contacts
	if entity.Contacts != nil {
		contacts = ToContacts(entity.Contacts)
	}
	return &domain.Account{
		Uuid:         entity.Uuid,
		Login:        entity.Login,
		PasswordHash: entity.PasswordHash,
		Contacts:     contacts,
		Active:       entity.Active,
	}
}

func ToContacts(entity *ContactsPostgres) *domain.Contacts {
	return &domain.Contacts{
		TgLogin: entity.TgLogin,
	}
}
