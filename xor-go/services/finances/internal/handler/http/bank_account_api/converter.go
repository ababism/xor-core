package bank_account_api

import (
	"github.com/google/uuid"
	openapitypes "github.com/oapi-codegen/runtime/types"
	"xor-go/services/finances/internal/domain"
	bankaccount "xor-go/services/finances/internal/handler/generated/bank-account"
)

func ToBankAccountDataDomain(data bankaccount.BankAccountData) domain.BankAccountData {
	return domain.BankAccountData{}
}

func ToBankAccountDataResponse(data domain.BankAccountData) bankaccount.BankAccountData {
	return bankaccount.BankAccountData{}
}

func CreateToDomain(input bankaccount.BankAccountCreate) domain.BankAccountCreate {
	return domain.BankAccountCreate{
		AccountUUID: ToGoogleUUID(input.AccountUUID),
		Login:       input.Login,
		Data:        ToBankAccountDataDomain(input.Data),
		Payments:    ToGoogleUUIDs(input.Payments),
	}
}

func UpdateToDomain(input bankaccount.BankAccountUpdate) domain.BankAccountUpdate {
	return domain.BankAccountUpdate{
		UUID:        ToGoogleUUID(input.UUID),
		AccountUUID: ToGoogleUUID(input.AccountUUID),
		Login:       input.Login,
		Funds:       input.Funds,
		Data:        ToBankAccountDataDomain(input.Data),
		Status:      input.Status,
		LastDealAt:  input.LastDealAt,
		Payments:    ToGoogleUUIDs(input.Payments),
	}
}

func DomainToResponse(input domain.BankAccountGet) bankaccount.BankAccountGet {
	return bankaccount.BankAccountGet{
		UUID:         ToResponseUUID(input.UUID),
		AccountUUID:  ToResponseUUID(input.AccountUUID),
		Login:        input.Login,
		Funds:        input.Funds,
		Data:         ToBankAccountDataResponse(input.Data),
		Status:       input.Status,
		LastDealAt:   input.LastDealAt,
		CreatedAt:    input.CreatedAt,
		LastUpdateAt: input.LastUpdateAt,
		Payments:     ToResponseUUIDs(input.Payments),
	}
}

func FilterToDomain(input *bankaccount.BankAccountFilter) *domain.BankAccountFilter {
	if input == nil {
		return nil
	}
	return &domain.BankAccountFilter{
		UUID:        ToGoogleUUIDPointer(input.UUID),
		AccountUUID: ToGoogleUUIDPointer(input.AccountUUID),
		Login:       input.Login,
		Funds:       input.Funds,
		Status:      input.Status,
	}
}

func ToGoogleUUIDs(ids []openapitypes.UUID) []uuid.UUID {
	uuids := make([]uuid.UUID, len(ids))
	for i, id := range ids {
		uuids[i] = ToGoogleUUID(id)
	}
	return uuids
}

func ToResponseUUIDs(ids []openapitypes.UUID) []uuid.UUID {
	uuids := make([]uuid.UUID, len(ids))
	for i, id := range ids {
		uuids[i] = ToResponseUUID(id)
	}
	return uuids
}

func ToGoogleUUID(id openapitypes.UUID) uuid.UUID {
	return id
}

func ToGoogleUUIDPointer(id *openapitypes.UUID) *uuid.UUID {
	return id
}

func ToResponseUUID(id uuid.UUID) openapitypes.UUID {
	return id
}
