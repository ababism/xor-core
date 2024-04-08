package postgre

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	global "go.opentelemetry.io/otel"
	"xor-go/pkg/xcommon"
	"xor-go/services/finances/internal/domain"
	"xor-go/services/finances/internal/repository/postgre/repo_models"
	"xor-go/services/finances/internal/service/adapters"
)

const (
	spanDefaultPayoutRequest = "payout-request/repository.postgre"
)

const (
	basePayoutRequestGetQuery = `
		SELECT uuid, receiver, amount, data, created_at
		FROM payout_requests
	`
	createPayoutRequestQuery = `
		INSERT INTO payout_requests (receiver, amount, data, created_at)
		VALUES ($1, $2, $3, $4)
		RETURNING uuid
	`
	deletePayoutRequestQuery = `
		DELETE FROM payout_requests WHERE uuid = $1
	`
)

var _ adapters.PayoutRequestRepository = &payoutRequestRepository{}

type payoutRequestRepository struct {
	db *sqlx.DB
}

func NewPayoutRequestRepository(db *sqlx.DB) adapters.PayoutRequestRepository {
	return &payoutRequestRepository{db: db}
}

func (r *payoutRequestRepository) Get(ctx context.Context, id uuid.UUID) (*domain.PayoutRequestGet, error) {
	tr := global.Tracer(adapters.ServiceNamePayoutRequest)
	_, span := tr.Start(ctx, spanDefaultPayoutRequest+".Get")
	defer span.End()

	payouts, err := r.List(ctx, &domain.PayoutRequestFilter{UUID: &id})
	if err != nil {
		return nil, err
	}
	return xcommon.EnsureSingle(payouts)
}

func (r *payoutRequestRepository) List(
	ctx context.Context,
	filter *domain.PayoutRequestFilter,
) ([]domain.PayoutRequestGet, error) {
	tr := global.Tracer(adapters.ServiceNamePayoutRequest)
	_, span := tr.Start(ctx, spanDefaultPayoutRequest+".List")
	defer span.End()

	paramsMap := mapGetPayoutRequestRequestParams(filter)
	query, args := xcommon.QueryWhereAnd(basePayoutRequestGetQuery, paramsMap)
	var payouts []repo_models.PayoutRequest
	err := r.db.SelectContext(ctx, &payouts, query, args...)
	if err != nil {
		return nil, err
	}
	return xcommon.ConvertSliceP(payouts, repo_models.ToPayoutRequestDomain), nil
}

func (r *payoutRequestRepository) Create(ctx context.Context, payout *domain.PayoutRequestCreate) (*uuid.UUID, error) {
	tr := global.Tracer(adapters.ServiceNamePayoutRequest)
	_, span := tr.Start(ctx, spanDefaultPayoutRequest+".Create")
	defer span.End()

	payoutPostgres := repo_models.CreateToPayoutRequestPostgres(payout)
	data, err := json.Marshal(payoutPostgres.Data)
	if err != nil {
		return nil, err
	}

	row := r.db.QueryRow(
		createPayoutRequestQuery,
		payoutPostgres.Receiver,
		payoutPostgres.Amount,
		string(data),
		payoutPostgres.CreatedAt,
	)
	var id uuid.UUID
	err = row.Scan(&id)
	if err != nil {
		return nil, err
	}

	return &id, nil
}

func (r *payoutRequestRepository) Delete(ctx context.Context, id uuid.UUID) error {
	tr := global.Tracer(adapters.ServiceNamePayoutRequest)
	_, span := tr.Start(ctx, spanDefaultPayoutRequest+".Delete")
	defer span.End()

	_, err := r.db.ExecContext(ctx, deletePayoutRequestQuery, id)
	return err
}

func mapGetPayoutRequestRequestParams(params *domain.PayoutRequestFilter) map[string]interface{} {
	if params == nil {
		return map[string]any{}
	}
	paramsMap := make(map[string]interface{})
	if params.UUID != nil {
		paramsMap["uuid"] = *params.UUID
	}
	if params.Receiver != nil {
		paramsMap["receiver"] = *params.Receiver
	}
	if params.Amount != nil {
		paramsMap["amount"] = *params.Amount
	}
	if params.ReceivedAt != nil {
		paramsMap["created_at"] = *params.ReceivedAt
	}
	return paramsMap
}
