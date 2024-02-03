package dto

import (
	"xor-go/services/sage/internal/domain"
)

func (r *RegisterAccountDto) ToRegisterAccount() *domain.RegisterAccount {
	return &domain.RegisterAccount{
		Login:    r.Login,
		Password: r.Password,
	}
}
