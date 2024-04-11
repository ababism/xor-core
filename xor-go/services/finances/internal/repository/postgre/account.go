package postgre

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
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
		SELECT uuid, account_uuid, login, funds, data, status, last_deal_at, created_at, updated_at
		FROM  bank_accounts
	`
	createBankAccountQuery = `
		INSERT INTO bank_accounts (account_uuid, login, funds, data, status)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING uuid
	`
	updateBankAccountQuery = `
		UPDATE bank_accounts
		SET 
		    account_uuid = $2,
		    login = $3,
		    funds = $4,
		    data = $5,
		    status = $6,
		    last_deal_at = $7
		WHERE uuid = $1;
	`
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

func (r *bankAccountRepository) Get(
	ctx context.Context,
	filter *domain.BankAccountFilter,
) (*domain.BankAccountGet, error) {
	tr := global.Tracer(adapters.ServiceNameBankAccount)
	_, span := tr.Start(ctx, spanDefaultBankAccount+".GetByLogin")
	defer span.End()

	accounts, err := r.List(ctx, filter)
	if err != nil {
		return nil, err
	}
	return xcommon.EnsureSingle(accounts)
}

func (r *bankAccountRepository) List(
	ctx context.Context,
	filter *domain.BankAccountFilter,
) ([]domain.BankAccountGet, error) {
	tr := global.Tracer(adapters.ServiceNameBankAccount)
	_, span := tr.Start(ctx, spanDefaultBankAccount+".List")
	defer span.End()

	paramsMap := mapGetBankAccountRequestParams(filter)
	query, args := xcommon.QueryWhereAnd(
		baseBankAccountGetQuery,
		paramsMap,
	)
	var accounts []repo_models.BankAccount
	err := r.db.SelectContext(ctx, &accounts, query, args...)
	if err != nil {
		return nil, err
	}
	return xcommon.ConvertSliceP(accounts, repo_models.ToBankAccountDomain), nil
}

func (r *bankAccountRepository) Create(ctx context.Context, account *domain.BankAccountCreate) (*uuid.UUID, error) {
	tr := global.Tracer(adapters.ServiceNameBankAccount)
	_, span := tr.Start(ctx, spanDefaultBankAccount+".Create")
	defer span.End()

	accountPostgres := repo_models.CreateToBankAccountPostgres(account)
	data, err := json.Marshal(accountPostgres.Data)
	if err != nil {
		return nil, err
	}

	row := r.db.QueryRow(
		createBankAccountQuery,
		accountPostgres.AccountUUID,
		accountPostgres.Login,
		accountPostgres.Funds,
		string(data),
		accountPostgres.Status,
	)
	var id uuid.UUID
	err = row.Scan(&id)
	if err != nil {
		return nil, err
	}
	return &id, err
}

func (r *bankAccountRepository) Update(ctx context.Context, account *domain.BankAccountUpdate) error {
	accountPostgres := repo_models.UpdateToBankAccountPostgres(account)
	_, err := r.db.ExecContext(
		ctx,
		updateBankAccountQuery,
		accountPostgres.UUID,
		accountPostgres.AccountUUID,
		accountPostgres.Login,
		accountPostgres.Data,
		accountPostgres.Status,
		accountPostgres.LastDealAt,
	)
	return err
}

func mapGetBankAccountRequestParams(params *domain.BankAccountFilter) map[string]any {
	if params == nil {
		return map[string]any{}
	}
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
