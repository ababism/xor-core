package postgre

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	global "go.opentelemetry.io/otel"
	"xor-go/pkg/xcommon"
	"xor-go/services/finances/internal/domain"
	"xor-go/services/finances/internal/repository/postgre/repo_models"
	"xor-go/services/finances/internal/service/adapters"
)

const (
	spanDefaultBankAccount = "bank-account/repository.postgre"
)

const (
	baseBankAccountGetQuery = `
		SELECT uuid, account_uuid, login, funds, data, status, last_deal_at, created_at, last_updated_at
		FROM  bank_accounts
		WHERE uuid = $1
	`
	createBankAccountQuery = `
		INSERT INTO bank_accounts (account_uuid, login, data, status)
		VALUES ($1, $2, $3, $4)
	`
	// GPT
	updateBankAccountQuery = `
		UPDATE bank_accounts SET password_hash = $1 WHERE uuid = $2
	`
	// ? deactivateQuery = `
	//	UPDATE bank_accounts SET active = false WHERE uuid = $1
	//`
)

var _ adapters.BankAccountRepository = &bankAccountRepository{}

type bankAccountRepository struct {
	db *sqlx.DB
}

func NewBankAccountRepository(db *sqlx.DB) adapters.BankAccountRepository {
	return &bankAccountRepository{db: db}
}

func (r *bankAccountRepository) Present(ctx context.Context, filter *domain.BankAccountFilter) (bool, error) {
	tr := global.Tracer(adapters.ServiceNameBankAccount)
	_, span := tr.Start(ctx, spanDefaultBankAccount+".Present")
	defer span.End()

	paramsMap := mapGetBankAccountRequestParams(filter)

	query, args := xcommon.QueryWhereAnd(
		baseBankAccountGetQuery,
		paramsMap,
	)
	presentQuery := fmt.Sprintf("SELECT EXISTS (%s)", query)
	var present bool
	err := r.db.GetContext(ctx, &present, presentQuery, args...)
	if err != nil {
		return false, err
	}
	return present, nil
}

func (r *bankAccountRepository) Get(ctx context.Context, filter *domain.BankAccountFilter) (*domain.BankAccountGet, error) {
	tr := global.Tracer(adapters.ServiceNameBankAccount)
	_, span := tr.Start(ctx, spanDefaultBankAccount+".Get")
	defer span.End()

	accounts, err := r.List(ctx, filter)
	if err != nil {
		return nil, err
	}
	return xcommon.EnsureSingle(accounts)
}

func (r *bankAccountRepository) List(ctx context.Context, filter *domain.BankAccountFilter) ([]domain.BankAccountGet, error) {
	tr := global.Tracer(adapters.ServiceNameBankAccount)
	_, span := tr.Start(ctx, spanDefaultBankAccount+".List")
	defer span.End()

	paramsMap := mapGetBankAccountRequestParams(filter)
	query, args := xcommon.QueryWhereAnd(baseBankAccountGetQuery, paramsMap)
	var accounts []repo_models.BankAccount
	err := r.db.SelectContext(ctx, &accounts, query, args...)
	if err != nil {
		return nil, err
	}
	return xcommon.ConvertSliceP(accounts, repo_models.ToBankAccountDomain), nil
}

func (r *bankAccountRepository) Create(ctx context.Context, account *domain.BankAccountCreate) error {
	tr := global.Tracer(adapters.ServiceNameBankAccount)
	_, span := tr.Start(ctx, spanDefaultBankAccount+".Create")
	defer span.End()

	accountPostgres := repo_models.CreateToBankAccountPostgres(account)
	_, err := r.db.ExecContext(
		ctx,
		createBankAccountQuery,
		accountPostgres.AccountUUID,
		accountPostgres.Login,
		accountPostgres.Data,
		accountPostgres.Status,
	)
	return err
}

func (r *bankAccountRepository) Update(ctx context.Context, account *domain.BankAccountPost) error {
	accountPostgres := repo_models.UpdateToBankAccountPostgres(account)
	_, err := r.db.ExecContext(
		ctx,
		updateBankAccountQuery,
		accountPostgres.AccountUUID,
		accountPostgres.Login,
		accountPostgres.Data,
		accountPostgres.Status,
		accountPostgres.UUID,
	)
	return err
}

func mapGetBankAccountRequestParams(params *domain.BankAccountFilter) map[string]any {
	paramsMap := make(map[string]any)
	if params.UUID != nil {
		paramsMap["uuid"] = params.UUID
	}
	if params.AccountUUID != nil {
		paramsMap["account_uuid"] = params.AccountUUID
	}
	if params.Login != nil {
		paramsMap["login"] = params.Login
	}
	if params.Funds != nil {
		paramsMap["funds"] = params.Funds
	}
	if params.Status != nil {
		paramsMap["status"] = params.Status
	}
	return paramsMap
}
