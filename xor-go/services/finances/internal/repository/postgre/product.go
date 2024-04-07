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
	spanDefaultProduct = "product/repository.postgre"
)

const (
	baseProductGetQuery = `
		SELECT uuid, name, price, is_available, created_at, updated_at
		FROM products
		WHERE uuid = $1
	`
	createProductQuery = `
		INSERT INTO products (uuid, name, price)
		VALUES ($1, $2, $3)
	`
	updateProductQuery = `
		UPDATE products
		SET 
		    name = $2,
		    price = $3,
		    is_available = $4
		WHERE uuid = $1;
	`
	updateProductAvailabilityQuery = `
		UPDATE products
		SET 
		    is_available = $2
		WHERE uuid = $1;
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
	_, span := tr.Start(ctx, spanDefaultProduct+".Get")
	defer span.End()

	products, err := r.List(ctx, &domain.ProductFilter{UUID: &id})
	if err != nil {
		return nil, err
	}
	return xcommon.EnsureSingle(products)
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

func (r *productRepository) Create(ctx context.Context, product *domain.ProductCreate) error {
	tr := global.Tracer(adapters.ServiceNameProduct)
	_, span := tr.Start(ctx, spanDefaultProduct+".Create")
	defer span.End()

	productPostgres := repo_models.CreateToProductPostgres(product)
	_, err := r.db.ExecContext(
		ctx,
		createProductQuery,
		productPostgres.UUID,
		productPostgres.Name,
		productPostgres.Price,
	)
	return err
}

func (r *productRepository) Update(ctx context.Context, product *domain.ProductUpdate) error {
	productPostgres := repo_models.UpdateToProductPostgres(product)
	_, err := r.db.ExecContext(
		ctx,
		updateProductQuery,
		productPostgres.UUID,
		productPostgres.Name,
		productPostgres.Price,
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
