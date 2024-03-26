package postgre

import (
	"context"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	global "go.opentelemetry.io/otel"
	"time"
	"xor-go/pkg/xcommon"
	"xor-go/services/finances/internal/domain"
	"xor-go/services/finances/internal/repository/postgre/repo_models"
	"xor-go/services/finances/internal/service/adapters"
)

const (
	spanDefaultDiscount = "discount/repository.postgre"
)

const (
	baseDiscountGetQuery = `
		SELECT uuid, created_by, percent, stand_alone, started_at, ended_at, status, created_at, updated_at
		FROM discounts
		WHERE uuid = $1
	`
	createDiscountQuery = `
		INSERT INTO discounts (created_by, percent, stand_alone, started_at, ended_at, status)
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	updateDiscountQuery = `
		UPDATE discounts
		SET 
		    created_by = $2,
		    percent = $3,
		    stand_alone = $4,
		    started_at = $5,
		    ended_at = $6,
		    status = $7
		WHERE uuid = $1;
	`
	finishDiscountQuery = `
		UPDATE discounts
		SET 
		    ended_at = $2
		WHERE uuid = $1;
	`
)

var _ adapters.DiscountRepository = &discountRepository{}

type discountRepository struct {
	db *sqlx.DB
}

func NewDiscountRepository(db *sqlx.DB) adapters.DiscountRepository {
	return &discountRepository{db: db}
}

func (r *discountRepository) Get(ctx context.Context, id uuid.UUID) (*domain.DiscountGet, error) {
	tr := global.Tracer(adapters.ServiceNameDiscount)
	_, span := tr.Start(ctx, spanDefaultDiscount+".Get")
	defer span.End()

	discounts, err := r.List(ctx, &domain.DiscountFilter{UUID: &id})
	if err != nil {
		return nil, err
	}
	return xcommon.EnsureSingle(discounts)
}

func (r *discountRepository) List(ctx context.Context, filter *domain.DiscountFilter) ([]domain.DiscountGet, error) {
	tr := global.Tracer(adapters.ServiceNameDiscount)
	_, span := tr.Start(ctx, spanDefaultDiscount+".List")
	defer span.End()

	paramsMap := mapGetDiscountRequestParams(filter)
	query, args := xcommon.QueryWhereAnd(baseDiscountGetQuery, paramsMap)
	var discounts []repo_models.Discount
	err := r.db.SelectContext(ctx, &discounts, query, args...)
	if err != nil {
		return nil, err
	}
	return xcommon.ConvertSliceP(discounts, repo_models.ToDiscountDomain), nil
}

func (r *discountRepository) Create(ctx context.Context, discount *domain.DiscountCreate) error {
	tr := global.Tracer(adapters.ServiceNameDiscount)
	_, span := tr.Start(ctx, spanDefaultDiscount+".Create")
	defer span.End()

	discountPostgres := repo_models.CreateToDiscountPostgres(discount)
	_, err := r.db.ExecContext(
		ctx,
		createDiscountQuery,
		discountPostgres.CreatedBy,
		discountPostgres.Percent,
		discountPostgres.StandAlone,
		discountPostgres.StartedAt,
		discountPostgres.EndedAt,
		discountPostgres.Status,
	)
	return err
}

func (r *discountRepository) Update(ctx context.Context, discount *domain.DiscountUpdate) error {
	discountPostgres := repo_models.UpdateToDiscountPostgres(discount)
	_, err := r.db.ExecContext(
		ctx,
		updateDiscountQuery,
		discountPostgres.UUID,
		discountPostgres.CreatedBy,
		discountPostgres.Percent,
		discountPostgres.StandAlone,
		discountPostgres.StartedAt,
		discountPostgres.EndedAt,
		discountPostgres.Status,
	)
	return err
}

func (r *discountRepository) EndDiscount(ctx context.Context, id uuid.UUID) error {
	_, err := r.db.ExecContext(
		ctx,
		finishDiscountQuery,
		id,
		time.Now(),
	)
	return err
}

func mapGetDiscountRequestParams(params *domain.DiscountFilter) map[string]interface{} {
	paramsMap := make(map[string]interface{})
	if params.UUID != nil {
		paramsMap["uuid"] = *params.UUID
	}
	if params.CreatedBy != nil {
		paramsMap["created_by"] = *params.CreatedBy
	}
	if params.Percent != nil {
		paramsMap["percent"] = *params.Percent
	}
	if params.StandAlone != nil {
		paramsMap["stand_alone"] = *params.StandAlone
	}
	if params.Status != nil {
		paramsMap["status"] = *params.Status
	}
	return paramsMap
}
