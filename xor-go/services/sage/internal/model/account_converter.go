package model

func (r *RegisterAccountDto) ToRegisterAccountEntity() *RegisterAccountEntity {
	return &RegisterAccountEntity{
		Login:    r.Login,
		Password: r.Password,
	}
}
