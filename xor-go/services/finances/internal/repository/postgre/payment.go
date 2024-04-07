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
	spanPaymentDefault  = "payment/repository/postgre"
	basePaymentGetQuery = `
		SELECT uuid, sender, receiver, data, url, status, ended_at, created_at
		FROM payments
		WHERE uuid = $1
	`
	createPaymentQuery = `
		INSERT INTO payments (uuid, sender, receiver, data, url, status, ended_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
)

var _ adapters.PaymentRepository = &paymentRepository{}

type paymentRepository struct {
	db *sqlx.DB
}

func NewPaymentRepository(db *sqlx.DB) adapters.PaymentRepository {
	return &paymentRepository{db: db}
}

func (r *paymentRepository) Get(ctx context.Context, filter *domain.PaymentFilter) (*domain.PaymentGet, error) {
	tr := global.Tracer(adapters.ServiceNamePayment)
	_, span := tr.Start(ctx, spanPaymentDefault+".Get")
	defer span.End()

	accounts, err := r.List(ctx, filter)
	if err != nil {
		return nil, err
	}
	return xcommon.EnsureSingle(accounts)
}

func (r *paymentRepository) List(ctx context.Context, filter *domain.PaymentFilter) ([]domain.PaymentGet, error) {
	tr := global.Tracer(adapters.ServiceNamePayment)
	_, span := tr.Start(ctx, spanPaymentDefault+".List")
	defer span.End()

	paramsMap := mapGetPaymentRequestParams(filter)
	query, args := xcommon.QueryWhereAnd(basePaymentGetQuery, paramsMap)
	var payments []repo_models.Payment
	err := r.db.SelectContext(ctx, &payments, query, args...)
	if err != nil {
		return nil, err
	}

	return xcommon.ConvertSliceP(payments, repo_models.ToPaymentDomain), nil
}

func (r *paymentRepository) Create(ctx context.Context, account *domain.PaymentCreate) error {
	tr := global.Tracer(adapters.ServiceNamePayment)
	_, span := tr.Start(ctx, spanPaymentDefault+".Create")
	defer span.End()

	accountPostgres := repo_models.CreateToPaymentPostgres(account)
	_, err := r.db.ExecContext(
		ctx,
		createPaymentQuery,
		accountPostgres.UUID,
		accountPostgres.Sender,
		accountPostgres.Receiver,
		accountPostgres.Data,
		accountPostgres.URL,
		accountPostgres.Status,
		accountPostgres.EndedAt,
	)
	return err
}

func mapGetPaymentRequestParams(params *domain.PaymentFilter) map[string]any {
	if params == nil {
		return map[string]any{}
	}
	paramsMap := make(map[string]any)
	if params.UUID != nil {
		paramsMap["uuid"] = params.UUID
	}
	if params.Sender != nil {
		paramsMap["sender"] = params.Sender
	}
	if params.Receiver != nil {
		paramsMap["receiver"] = params.Receiver
	}
	if params.URL != nil {
		paramsMap["url"] = params.URL
	}
	if params.Status != nil {
		paramsMap["status"] = params.Status
	}
	if params.EndedAt != nil {
		paramsMap["ended_at"] = params.EndedAt
	}
	return paramsMap
}
