package postgres

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"xor-go/pkg/xcommon"
	"xor-go/services/sage/internal/domain"
	"xor-go/services/sage/internal/repository/postgres/entity"
	"xor-go/services/sage/internal/service/adapter"
)

const (
	getAnyQuery = `
		SELECT uuid, login, password_hash, contacts, active FROM account
	`
	createQuery = `
		INSERT INTO account (uuid, login, password_hash, contacts, active)
		VALUES ($1, $2, $3, $4, $5)
	`
	updatePasswordQuery = `
		UPDATE account SET password_hash = $1 WHERE uuid = $2
	`
	deactivateQuery = `
		UPDATE account SET active = false WHERE uuid = $1
	`
)

var _ adapter.AccountRepository = &accountRepository{}

type accountRepository struct {
	logger *zap.Logger
	db     *sqlx.DB
}

func NewAccountPostgresRepository(logger *zap.Logger, db *sqlx.DB) adapter.AccountRepository {
	return &accountRepository{logger: logger, db: db}
}

func (r *accountRepository) Present(ctx context.Context, filter *domain.AccountFilter) (bool, error) {
	argsMap := make(map[string]any)
	if filter.Uuid != nil {
		argsMap["uuid"] = filter.Uuid
	}
	if filter.Login != nil {
		argsMap["login"] = filter.Login
	}
	selectQuery := "SELECT * FROM account"
	query, args := xcommon.QueryWhereAnd(selectQuery, argsMap)
	presentQuery := fmt.Sprintf("SELECT EXISTS (%s)", query)
	var present bool
	err := r.db.GetContext(ctx, &present, presentQuery, args...)
	if err != nil {
		return false, err
	}
	return present, nil
}

func (r *accountRepository) List(ctx context.Context, filter *domain.AccountFilter) ([]domain.Account, error) {
	paramsMap := mapGetRequestParams(filter)
	query, args := xcommon.QueryWhereAnd(getAnyQuery, paramsMap)
	var accountPostgres []entity.AccountPostgres
	err := r.db.SelectContext(ctx, &accountPostgres, query, args...)
	if err != nil {
		return nil, err
	}
	return xcommon.ConvertSliceP(accountPostgres, entity.ToAccount), nil
}

func (r *accountRepository) Get(ctx context.Context, filter *domain.AccountFilter) (*domain.Account, error) {
	accounts, err := r.List(ctx, filter)
	if err != nil {
		return nil, err
	}
	return xcommon.EnsureSingle(accounts)
}

func (r *accountRepository) Create(ctx context.Context, account *domain.Account) error {
	accountPostgres := entity.ToAccountPostgres(account)
	_, err := r.db.ExecContext(
		ctx,
		createQuery,
		accountPostgres.Uuid,
		accountPostgres.Login,
		accountPostgres.PasswordHash,
		accountPostgres.Contacts,
		accountPostgres.Active,
	)
	return err
}

func (r *accountRepository) UpdatePassword(ctx context.Context, uuid uuid.UUID, passwordHash string) error {
	_, err := r.db.ExecContext(
		ctx,
		updatePasswordQuery,
		passwordHash,
		uuid,
	)
	return err
}

func (r *accountRepository) Deactivate(ctx context.Context, uuid uuid.UUID) error {
	_, err := r.db.ExecContext(ctx, deactivateQuery, uuid)
	return err
}

func mapGetRequestParams(params *domain.AccountFilter) map[string]any {
	paramsMap := make(map[string]any)
	if params.Uuid != nil {
		paramsMap["uuid"] = params.Uuid
	}
	if params.Login != nil {
		paramsMap["login"] = params.Login
	}
	return paramsMap
}
