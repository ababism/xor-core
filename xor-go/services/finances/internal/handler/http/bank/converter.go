package bank

import (
	"xor-go/services/finances/internal/domain"
)

func ToBankAccountDataDomain(data BankAccountData) domain.BankAccountData {
	return domain.BankAccountData{}
}

func ToBankAccountDataGet(data domain.BankAccountData) BankAccountData {
	return BankAccountData{}
}

func CreateToDomain(input BankAccountCreate) domain.BankAccountCreate {
	return domain.BankAccountCreate{
		AccountUUID: input.AccountUUID,
		Login:       input.Login,
		Data:        ToBankAccountDataDomain(input.Data),
		Payments:    input.Payments,
	}
}

func UpdateToDomain(input BankAccountUpdate) domain.BankAccountUpdate {
	return domain.BankAccountUpdate{
		UUID:        input.UUID,
		AccountUUID: input.AccountUUID,
		Login:       input.Login,
		Funds:       input.Funds,
		Data:        ToBankAccountDataDomain(input.Data),
		Status:      input.Status,
		LastDealAt:  input.LastDealAt,
		Payments:    input.Payments,
	}
}

func DomainToGet(input domain.BankAccountGet) BankAccountGet {
	return BankAccountGet{
		UUID:         input.UUID,
		AccountUUID:  input.AccountUUID,
		Login:        input.Login,
		Funds:        input.Funds,
		Data:         ToBankAccountDataGet(input.Data),
		Status:       input.Status,
		LastDealAt:   input.LastDealAt,
		CreatedAt:    input.CreatedAt,
		LastUpdateAt: input.LastUpdateAt,
		Payments:     input.Payments,
	}
}

func FilterToDomain(input *BankAccountFilter) *domain.BankAccountFilter {
	if input == nil {
		return nil
	}
	return &domain.BankAccountFilter{
		UUID:        input.UUID,
		AccountUUID: input.AccountUUID,
		Login:       input.Login,
		Funds:       input.Funds,
		Status:      input.Status,
	}
}
