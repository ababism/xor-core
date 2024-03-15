package postgre

import (
	"context"
	"github.com/jmoiron/sqlx"
	global "go.opentelemetry.io/otel"
	"xor-go/pkg/xcommon"
	"xor-go/services/finances/internal/domain"
	"xor-go/services/finances/internal/repository/postgre/repo_models"
	"xor-go/services/finances/internal/service/adapters"
)

const (
	spanPaymentsDefault = "payments/repository/postgre"
)

const (
	baseGetPaymentQuery = `
		SELECT uuid, account_uuid, login, funds, data, status, last_deal_at, created_at, last_updated_at
		FROM  bank_accounts
		WHERE uuid = $1
	`
	createPaymentQuery = `
		INSERT INTO bank_accounts (account_uuid, login, data, status)
		VALUES ($1, $2, $3, $4)
	`
	// GPT
	updatePaymentQuery = `
		UPDATE bank_accounts SET password_hash = $1 WHERE uuid = $2
	`
	// ? deactivateQuery = `
	//	UPDATE bank_accounts SET active = false WHERE uuid = $1
	//`
)

var _ adapters.PaymentRepository = &paymentRepository{}

type paymentRepository struct {
	db *sqlx.DB
}

func NewPaymentRepository(db *sqlx.DB) adapters.PaymentRepository {
	return &paymentRepository{db: db}
}

func (r *paymentRepository) Get(ctx context.Context, filter *domain.PaymentFilter) (*domain.Payment, error) {
	tr := global.Tracer(adapters.ServiceNameBankAccount)
	_, span := tr.Start(ctx, spanPaymentsDefault+".Get")
	defer span.End()

	accounts, err := r.List(ctx, filter)
	if err != nil {
		return nil, err
	}
	return xcommon.EnsureSingle(accounts)
}

func (r *paymentRepository) List(ctx context.Context, filter *domain.PaymentFilter) ([]domain.Payment, error) {
	tr := global.Tracer(adapters.ServiceNameBankAccount)
	_, span := tr.Start(ctx, spanPaymentsDefault+".List")
	defer span.End()

	paramsMap := mapGetPaymentRequestParams(filter)
	query, args := xcommon.QueryWhereAnd(baseGetPaymentQuery, paramsMap)
	var payments []repo_models.Payment
	err := r.db.SelectContext(ctx, &payments, query, args...)
	if err != nil {
		return nil, err
	}
	return xcommon.ConvertSliceP(payments, repo_models.ToPayment), nil
}

func (r *paymentRepository) Create(ctx context.Context, account *domain.Payment) error {
	tr := global.Tracer(adapters.ServiceNameBankAccount)
	_, span := tr.Start(ctx, spanPaymentsDefault+".Create")
	defer span.End()

	accountPostgres := repo_models.ToPaymentPostgres(account)
	_, err := r.db.ExecContext(
		ctx,
		createPaymentQuery,
		// GPT
		account.AccountUUID,
		account.Login,
		account.Data,
		account.Status,
	)
	return err
}

func (r *paymentRepository) Update(ctx context.Context, account *domain.Payment) error {
	_, err := r.db.ExecContext(
		ctx,
		// GPT
		updatePaymentsQuery,
		account.AccountUUID,
		account.Login,
		account.Data,
		account.Status,
		account.UUID,
	)
	return err
}

func mapGetPaymentRequestParams(params *domain.PaymentFilter) map[string]any {
	paramsMap := make(map[string]any)
	if params.UUID != nil {
		paramsMap["uuid"] = params.UUID
	}
	// TODO GPT

	return paramsMap
}
