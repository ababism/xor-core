package service

import (
	"context"
	"errors"
	"github.com/google/uuid"
	global "go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"xor-go/services/finances/internal/domain"
	"xor-go/services/finances/internal/service/adapters"
)

const (
	spanDefaultProduct = "product/service."
)

var _ adapters.ProductService = &productService{}

type productService struct {
	r adapters.ProductRepository
}

func NewProductService(productRepository adapters.ProductRepository) adapters.ProductService {
	return &productService{r: productRepository}
}

func getProductTracerSpan(ctx context.Context, name string) (trace.Tracer, context.Context, trace.Span) {
	tr := global.Tracer(adapters.ServiceNameProduct)
	newCtx, span := tr.Start(ctx, spanDefaultProduct+name)
	return tr, newCtx, span
}

func (s *productService) Get(ctx context.Context, id uuid.UUID) (*domain.ProductGet, error) {
	_, newCtx, span := getProductTracerSpan(ctx, ".GetByLogin")
	defer span.End()

	product, err := s.r.Get(newCtx, id)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *productService) GetPrice(ctx context.Context, productUUIDs []uuid.UUID) (*float32, error) {
	_, newCtx, span := getProductTracerSpan(ctx, ".GetPrice")
	defer span.End()

	price, err := s.r.GetPrice(newCtx, productUUIDs)
	if err != nil {
		return nil, err
	}

	return price, nil
}

func (s *productService) List(ctx context.Context, filter *domain.ProductFilter) ([]domain.ProductGet, error) {
	_, newCtx, span := getProductTracerSpan(ctx, ".List")
	defer span.End()

	products, err := s.r.List(newCtx, filter)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (s *productService) Create(ctx context.Context, product *domain.ProductCreate) (*uuid.UUID, error) {
	_, newCtx, span := getProductTracerSpan(ctx, ".Create")
	defer span.End()

	id, err := s.r.Create(newCtx, product)
	if err != nil {
		return nil, err
	}

	return id, nil
}

func (s *productService) Update(ctx context.Context, product *domain.ProductUpdate) error {
	_, newCtx, span := getProductTracerSpan(ctx, ".Update")
	defer span.End()

	err := s.r.Update(newCtx, product)
	if err != nil {
		return err
	}

	return nil
}

func (s *productService) SetAvailability(ctx context.Context, id uuid.UUID, isAvailable bool) error {
	_, newCtx, span := getProductTracerSpan(ctx, ".SetAvailability")
	defer span.End()

	if id == uuid.Nil {
		return errors.New("product ID cannot be nil")
	}

	err := s.r.SetAvailability(newCtx, id, isAvailable)
	if err != nil {
		return err
	}

	return nil
}
