package dto

type RegisterAccountDto struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}
