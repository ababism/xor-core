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
	_, newCtx, span := getProductTracerSpan(ctx, ".Get")
	defer span.End()

	product, err := s.r.Get(newCtx, id)
	if err != nil {
		return nil, err
	}

	return product, nil
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

func (s *productService) Create(ctx context.Context, product *domain.ProductCreate) error {
	_, newCtx, span := getProductTracerSpan(ctx, ".Create")
	defer span.End()

	err := s.r.Create(newCtx, product)
	if err != nil {
		return err
	}

	return nil
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

func (s *productService) Disable(ctx context.Context, id uuid.UUID) error {
	_, newCtx, span := getProductTracerSpan(ctx, ".Disable")
	defer span.End()

	if id == uuid.Nil {
		return errors.New("product ID cannot be nil")
	}

	err := s.r.Disable(newCtx, id)
	if err != nil {
		return err
	}

	return nil
}