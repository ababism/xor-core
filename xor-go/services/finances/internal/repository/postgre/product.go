package postgre

import (
	"context"
	"database/sql"
	"errors"
	"xor-go/services/finances/internal/domain"
	"xor-go/services/finances/internal/repository/postgre/repo_models"
	"xor-go/services/finances/internal/service/adapters"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	global "go.opentelemetry.io/otel"
	"xor-go/pkg/xcommon"
)

const (
	spanDefaultProduct = "product/repository.postgre"
)

const (
	baseProductGetQuery = `
		SELECT uuid, name, price, is_available, created_at, updated_at
		FROM products
	`
	createProductQuery = `
		INSERT INTO products (name, price, info, is_available)
		VALUES ($1, $2, $3, $4)
		RETURNING uuid
	`
	updateProductQuery = `
		UPDATE products
		SET 
		    name = $2,
		    price = $3,
		    info = $4,
		    is_available = $5
		WHERE uuid = $1;
	`
	updateProductAvailabilityQuery = `
		UPDATE products
		SET 
		    is_available = $2
		WHERE uuid = $1;
	`
	getPrice = `
		SELECT
			SUM(
				CASE
					WHEN dp.product_uuid IS NOT NULL THEN
						CASE
							WHEN d.percent IS NULL THEN p.price
							ELSE p.price - (p.price * d.percent / 100)
						END
					ELSE
						p.price
					END
				) AS total_price
			FROM
				products p
			LEFT JOIN
				(
					SELECT
						dp.product_uuid,
						MAX(d.percent) AS max_percent
					FROM
						discounts_products dp
					JOIN
						discounts d ON dp.discount_uuid = d.uuid
					WHERE
						d.started_at <= CURRENT_TIMESTAMP
						AND d.ended_at >= CURRENT_TIMESTAMP
						AND d.status = 'active'
					GROUP BY
						dp.product_uuid
				) AS max_discount ON p.uuid = max_discount.product_uuid
			LEFT JOIN
				discounts d ON max_discount.max_percent = d.percent
			LEFT JOIN
				discounts_products dp ON d.uuid = dp.discount_uuid AND dp.product_uuid = p.uuid
			WHERE
				p.uuid IN ($1);
`
)

var _ adapters.ProductRepository = &productRepository{}

type productRepository struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) adapters.ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) Get(ctx context.Context, id uuid.UUID) (*domain.ProductGet, error) {
	tr := global.Tracer(adapters.ServiceNameProduct)
	_, span := tr.Start(ctx, spanDefaultProduct+".GetByLogin")
	defer span.End()

	products, err := r.List(ctx, &domain.ProductFilter{UUID: &id})
	if err != nil {
		return nil, err
	}
	return xcommon.EnsureSingle(products)
}

func (r *productRepository) GetPrice(ctx context.Context, productUUIDs []uuid.UUID) (*float32, error) {
	tr := global.Tracer(adapters.ServiceNamePurchaseRequest)
	_, span := tr.Start(ctx, spanDefaultPurchaseRequest+".GetPrice")
	defer span.End()

	args := make([]interface{}, len(productUUIDs))
	for i, id := range productUUIDs {
		args[i] = id
	}

	var totalPrice float32
	err := r.db.Get(&totalPrice, r.db.Rebind(getPrice), args...)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	return &totalPrice, nil
}

func (r *productRepository) List(ctx context.Context, filter *domain.ProductFilter) ([]domain.ProductGet, error) {
	tr := global.Tracer(adapters.ServiceNameProduct)
	_, span := tr.Start(ctx, spanDefaultProduct+".List")
	defer span.End()

	paramsMap := mapGetProductRequestParams(filter)
	query, args := xcommon.QueryWhereAnd(baseProductGetQuery, paramsMap)
	var products []repo_models.Product
	err := r.db.SelectContext(ctx, &products, query, args...)
	if err != nil {
		return nil, err
	}
	return xcommon.ConvertSliceP(products, repo_models.ToProductDomain), nil
}

func (r *productRepository) Create(ctx context.Context, product *domain.ProductCreate) (*uuid.UUID, error) {
	tr := global.Tracer(adapters.ServiceNameProduct)
	_, span := tr.Start(ctx, spanDefaultProduct+".Create")
	defer span.End()

	productPostgres := repo_models.CreateToProductPostgres(product)
	row := r.db.QueryRow(
		createProductQuery,
		productPostgres.Name,
		productPostgres.Price,
		productPostgres.Info,
		productPostgres.IsAvailable,
	)

	var id uuid.UUID
	err := row.Scan(&id)
	if err != nil {
		return nil, err
	}
	return &id, err
}

func (r *productRepository) Update(ctx context.Context, product *domain.ProductUpdate) error {
	productPostgres := repo_models.UpdateToProductPostgres(product)
	_, err := r.db.ExecContext(
		ctx,
		updateProductQuery,
		productPostgres.UUID,
		productPostgres.Name,
		productPostgres.Price,
		productPostgres.Info,
		productPostgres.IsAvailable,
	)
	return err
}

func (r *productRepository) SetAvailability(ctx context.Context, id uuid.UUID, isAvailable bool) error {
	_, err := r.db.ExecContext(
		ctx,
		updateProductAvailabilityQuery,
		id,
		isAvailable,
	)
	return err
}

func mapGetProductRequestParams(params *domain.ProductFilter) map[string]interface{} {
	if params == nil {
		return map[string]any{}
	}
	paramsMap := make(map[string]interface{})
	if params.UUID != nil {
		paramsMap["uuid"] = *params.UUID
	}
	if params.Name != nil {
		paramsMap["name"] = *params.Name
	}
	if params.Price != nil {
		paramsMap["price"] = *params.Price
	}
	return paramsMap
}
