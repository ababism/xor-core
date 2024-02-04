package dto

import (
	"xor-go/services/sage/internal/domain"
)

func ToRegisterAccount(dto *RegisterAccountDto) *domain.RegisterAccount {
	return &domain.RegisterAccount{
		Login:    dto.Login,
		Password: dto.Password,
	}
}

func ToAccountDto(domain *domain.Account) *AccountDto {
	var contacts *ContactsDto
	if domain.Contacts != nil {
		contacts = ToContactsDto(domain.Contacts)
	}
	return &AccountDto{
		Uuid:     domain.Uuid,
		Login:    domain.Login,
		Contacts: contacts,
		Active:   domain.Active,
	}
}

func ToContactsDto(domain *domain.Contacts) *ContactsDto {
	return &ContactsDto{TgLogin: domain.TgLogin}
}
