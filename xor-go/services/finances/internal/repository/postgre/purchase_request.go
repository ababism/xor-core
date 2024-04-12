package postgre

import (
	"context"
	"fmt"
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
		SELECT uuid, sender, receiver, status, webhook_url, created_at
		FROM purchase_requests
	`
	createPurchaseRequestQuery = `
	INSERT INTO purchase_requests (sender, receiver, status, amount, webhook_url, created_at)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING uuid
	`
	createPurchaseProductsQuery = `
		INSERT INTO purchase_requests_products (request_uuid, product_uuid)
		VALUES
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

	purchase, err := xcommon.EnsureSingle(purchaseRequests)
	if err != nil {
		return nil, err
	}

	return purchase, nil
}

func (r *purchaseRequestRepository) List(
	ctx context.Context,
	filter *domain.PurchaseRequestFilter,
) ([]domain.PurchaseRequestGet, error) {
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

func (r *purchaseRequestRepository) Create(
	ctx context.Context,
	purchase *domain.PurchaseRequestCreate,
	amount float32,
) (*uuid.UUID, error) {
	tr := global.Tracer(adapters.ServiceNamePurchaseRequest)
	_, span := tr.Start(ctx, spanDefaultPurchaseRequest+".Create")
	defer span.End()

	purchasePostgres := repo_models.CreateToPurchaseRequestPostgres(purchase, amount)
	row := r.db.QueryRow(
		createPurchaseRequestQuery,
		purchasePostgres.Sender,
		purchasePostgres.Receiver,
		"pending",
		purchasePostgres.Amount,
		purchasePostgres.WebhookURL,
		purchasePostgres.CreatedAt,
	)
	var id uuid.UUID
	err := row.Scan(&id)
	if err != nil {
		return nil, err
	}

	rows := ""
	for i, productId := range purchase.Products {
		rows += fmt.Sprintf("('%s', '%s')", id, productId)
		if i != len(purchase.Products)-1 {
			rows += ", "
		}
	}

	query := createPurchaseProductsQuery + rows
	_, err = r.db.ExecContext(
		ctx,
		query,
	)
	if err != nil {
		return nil, err
	}

	return &id, nil
}

func (r *purchaseRequestRepository) Delete(ctx context.Context, id uuid.UUID) error {
	tr := global.Tracer(adapters.ServiceNamePurchaseRequest)
	_, span := tr.Start(ctx, spanDefaultPurchaseRequest+".Delete")
	defer span.End()

	_, err := r.db.ExecContext(ctx, deletePurchaseRequestQuery, id)
	return err
}

func mapGetPurchaseRequestRequestParams(params *domain.PurchaseRequestFilter) map[string]interface{} {
	if params == nil {
		return map[string]any{}
	}
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
	if params.Status != nil {
		paramsMap["status"] = *params.Status
	}
	if params.WebhookURL != nil {
		paramsMap["webhook_url"] = *params.WebhookURL
	}
	if params.CreatedAt != nil {
		paramsMap["created_at"] = *params.CreatedAt
	}
	return paramsMap
}
