package repo_models

import (
	"github.com/google/uuid"
	"time"
	"xor-go/services/finances/internal/domain"
)

type BankAccountData struct {
}

type BankAccount struct {
	UUID          uuid.UUID       `db:"uuid"`
	AccountUUID   uuid.UUID       `db:"account_uuid"`
	Login         string          `db:"login"`
	Funds         float32         `db:"funds"`
	Data          BankAccountData `db:"data"`
	Status        string          `db:"status"`
	LastDealAt    time.Time       `db:"last_deal_at"`
	CreatedAt     time.Time       `db:"created_at"`
	LastUpdatedAt time.Time       `db:"updated_at"`
}

func CreateToBankAccountPostgres(model *domain.BankAccountCreate) *BankAccount {
	id, _ := uuid.NewUUID()
	return &BankAccount{
		UUID:        id,
		AccountUUID: model.AccountUUID,
		Login:       model.Login,
		Funds:       0,
		Data:        ToBankAccountDataPostgres(model.Data),
		Status:      domain.AccountStatusActive,
	}
}

func UpdateToBankAccountPostgres(model *domain.BankAccountUpdate) *BankAccount {
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

func ToBankAccountDomain(model *BankAccount) *domain.BankAccountGet {
	return &domain.BankAccountGet{
		UUID:         model.UUID,
		AccountUUID:  model.AccountUUID,
		Login:        model.Login,
		Funds:        model.Funds,
		Data:         ToBankAccountDataDomain(model.Data),
		Status:       model.Status,
		LastDealAt:   model.LastDealAt,
		CreatedAt:    model.CreatedAt,
		LastUpdateAt: model.LastUpdatedAt,
	}
}

func ToBankAccountDataDomain(model BankAccountData) domain.BankAccountData {
	return domain.BankAccountData{}
}
