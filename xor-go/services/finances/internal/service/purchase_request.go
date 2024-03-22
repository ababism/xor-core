package service

import (
	"context"
	"github.com/google/uuid"
	global "go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"xor-go/services/finances/internal/domain"
	"xor-go/services/finances/internal/service/adapters"
)

const (
	spanDefaultPurchaseRequest = "purchase-request/service."
)

var _ adapters.PurchaseRequestService = &purchaseRequestService{}

type purchaseRequestService struct {
	r adapters.PurchaseRequestRepository
}

func NewPurchaseRequestService(purchaseRequestRepository adapters.PurchaseRequestRepository) adapters.PurchaseRequestService {
	return &purchaseRequestService{r: purchaseRequestRepository}
}

func getPurchaseRequestTracerSpan(ctx context.Context, name string) (trace.Tracer, context.Context, trace.Span) {
	tr := global.Tracer(adapters.ServiceNamePurchaseRequest)
	newCtx, span := tr.Start(ctx, spanDefaultPurchaseRequest+name)
	return tr, newCtx, span
}

func (s *purchaseRequestService) Get(ctx context.Context, id uuid.UUID) (*domain.PurchaseRequestGet, error) {
	_, newCtx, span := getPurchaseRequestTracerSpan(ctx, ".Get")
	defer span.End()

	purchaseRequest, err := s.r.Get(newCtx, id)
	if err != nil {
		return nil, err
	}

	return purchaseRequest, nil
}

func (s *purchaseRequestService) List(ctx context.Context, filter *domain.PurchaseRequestFilter) ([]domain.PurchaseRequestGet, error) {
	_, newCtx, span := getPurchaseRequestTracerSpan(ctx, ".List")
	defer span.End()

	purchaseRequests, err := s.r.List(newCtx, filter)
	if err != nil {
		return nil, err
	}

	return purchaseRequests, nil
}

func (s *purchaseRequestService) Create(ctx context.Context, purchase *domain.PurchaseRequestCreate) error {
	_, newCtx, span := getPurchaseRequestTracerSpan(ctx, ".Create")
	defer span.End()

	err := s.r.Create(newCtx, purchase)
	if err != nil {
		return err
	}

	return nil
}

func (s *purchaseRequestService) Archive(ctx context.Context, id uuid.UUID) error {
	_, newCtx, span := getPurchaseRequestTracerSpan(ctx, ".Archive")
	defer span.End()

	err := s.r.Delete(newCtx, id)
	if err != nil {
		return err
	}

	return nil
}
