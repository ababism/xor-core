package postgre

import (
	"context"
	"xor-go/services/finances/internal/domain"
	"xor-go/services/finances/internal/repository/postgre/repo_models"
	"xor-go/services/finances/internal/service/adapters"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	global "go.opentelemetry.io/otel"
	"xor-go/pkg/xcommon"
)

const (
	spanDefaultPurchaseRequest = "purchase-request/repository.postgre"
)

const (
	basePurchaseRequestGetQuery = `
		SELECT uuid, sender, receiver, webhook_url, received_at
		FROM purchase_requests
		WHERE uuid = $1
	`
	createPurchaseRequestQuery = `
		INSERT INTO purchase_requests (sender, receiver, webhook_url, received_at)
		VALUES ($1, $2, $3, $4)
	`
	createPurchaseProductsQuery = `
		INSERT INTO purchase_requests_products (request_uuid, product_uuid)
		VALUES ($1, $2)
	`
	deletePurchaseRequestQuery = `
		DELETE FROM purchase_requests WHERE uuid = $1
	`
)

var _ adapters.PurchaseRequestRepository = &purchaseRequestRepository{}

type purchaseRequestRepository struct {
	db *sqlx.DB
}

func NewPurchaseRequestRepository(db *sqlx.DB) adapters.PurchaseRequestRepository {
	return &purchaseRequestRepository{db: db}
}

func (r *purchaseRequestRepository) Get(ctx context.Context, id uuid.UUID) (*domain.PurchaseRequestGet, error) {
	tr := global.Tracer(adapters.ServiceNamePurchaseRequest)
	_, span := tr.Start(ctx, spanDefaultPurchaseRequest+".Get")
	defer span.End()

	purchaseRequests, err := r.List(ctx, &domain.PurchaseRequestFilter{UUID: &id})
	if err != nil {
		return nil, err
	}
	return xcommon.EnsureSingle(purchaseRequests)
}

func (r *purchaseRequestRepository) List(ctx context.Context, filter *domain.PurchaseRequestFilter) ([]domain.PurchaseRequestGet, error) {
	tr := global.Tracer(adapters.ServiceNamePurchaseRequest)
	_, span := tr.Start(ctx, spanDefaultPurchaseRequest+".List")
	defer span.End()

	paramsMap := mapGetPurchaseRequestRequestParams(filter)
	query, args := xcommon.QueryWhereAnd(basePurchaseRequestGetQuery, paramsMap)
	var purchaseRequests []repo_models.PurchaseRequest
	err := r.db.SelectContext(ctx, &purchaseRequests, query, args...)
	if err != nil {
		return nil, err
	}
	return xcommon.ConvertSliceP(purchaseRequests, repo_models.ToPurchaseRequestDomain), nil
}

func (r *purchaseRequestRepository) Create(ctx context.Context, purchase *domain.PurchaseRequestCreate) error {
	tr := global.Tracer(adapters.ServiceNamePurchaseRequest)
	_, span := tr.Start(ctx, spanDefaultPurchaseRequest+".Create")
	defer span.End()

	purchasePostgres := repo_models.CreateToPurchaseRequestPostgres(purchase)
	_, err := r.db.ExecContext(
		ctx,
		createPurchaseRequestQuery,
		purchasePostgres.Sender,
		purchasePostgres.Receiver,
		purchasePostgres.WebhookURL,
		purchasePostgres.CreatedAt,
	)
	return err
}

func (r *purchaseRequestRepository) Delete(ctx context.Context, id uuid.UUID) error {
	tr := global.Tracer(adapters.ServiceNamePurchaseRequest)
	_, span := tr.Start(ctx, spanDefaultPurchaseRequest+".Delete")
	defer span.End()

	_, err := r.db.ExecContext(ctx, deletePurchaseRequestQuery, id)
	return err
}

func mapGetPurchaseRequestRequestParams(params *domain.PurchaseRequestFilter) map[string]interface{} {
	paramsMap := make(map[string]interface{})
	if params.UUID != nil {
		paramsMap["uuid"] = *params.UUID
	}
	if params.Sender != nil {
		paramsMap["sender"] = *params.Sender
	}
	if params.Receiver != nil {
		paramsMap["receiver"] = *params.Receiver
	}
	if params.WebhookURL != nil {
		paramsMap["webhook_url"] = *params.WebhookURL
	}
	if params.ReceivedAt != nil {
		paramsMap["received_at"] = *params.ReceivedAt
	}
	return paramsMap
}
