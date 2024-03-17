package repo_models

import "xor-go/services/finances/internal/domain"

func ToBankAccountPostgres(model *domain.BankAccountGet) *BankAccount {
	return &BankAccount{
		UUID:        model.UUID,
		AccountUUID: model.AccountUUID,
		Login:       model.Login,
		Funds:       model.Funds,
		Data:        ToBankAccountDataPostgres(model.Data),
		Status:      model.Status,
		LastDealAt:  model.LastDealAt,
	}
}

func ToBankAccountDataPostgres(model domain.BankAccountData) BankAccountData {
	return BankAccountData{}
}

func ToBankAccount(model *BankAccount) *domain.BankAccountGet {
	return &domain.BankAccountGet{
		UUID:         model.UUID,
		AccountUUID:  model.AccountUUID,
		Login:        model.Login,
		Funds:        model.Funds,
		Data:         ToBankAccountData(model.Data),
		Status:       model.Status,
		LastDealAt:   model.LastDealAt,
		CreatedAt:    model.CreatedAt,
		LastUpdateAt: model.LastUpdatedAt,
	}
}

func ToBankAccountData(model BankAccountData) domain.BankAccountData {
	return domain.BankAccountData{}
}

// TODO GPT
