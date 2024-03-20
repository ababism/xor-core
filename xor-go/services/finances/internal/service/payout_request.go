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
	spanDefaultPayoutRequest = "payout-request/service."
)

var _ adapters.PayoutRequestService = &payoutRequestService{}

type payoutRequestService struct {
	r adapters.PayoutRequestRepository
}

func NewPayoutRequestService(payoutRequestRepository adapters.PayoutRequestRepository) adapters.PayoutRequestService {
	return &payoutRequestService{r: payoutRequestRepository}
}

func getPayoutRequestTracerSpan(ctx context.Context, name string) (trace.Tracer, context.Context, trace.Span) {
	tr := global.Tracer(adapters.ServiceNamePayoutRequest)
	newCtx, span := tr.Start(ctx, spanDefaultPayoutRequest+name)
	return tr, newCtx, span
}

func (s *payoutRequestService) Get(ctx context.Context, id uuid.UUID) (*domain.PayoutRequestGet, error) {
	_, newCtx, span := getPayoutRequestTracerSpan(ctx, ".Get")
	defer span.End()

	payoutRequest, err := s.r.Get(newCtx, id)
	if err != nil {
		return nil, err
	}

	return payoutRequest, nil
}

func (s *payoutRequestService) List(ctx context.Context, filter *domain.PayoutRequestFilter) ([]domain.PayoutRequestGet, error) {
	_, newCtx, span := getPayoutRequestTracerSpan(ctx, ".List")
	defer span.End()

	payoutRequests, err := s.r.List(newCtx, filter)
	if err != nil {
		return nil, err
	}

	return payoutRequests, nil
}

func (s *payoutRequestService) Create(ctx context.Context, payout *domain.PayoutRequestCreate) error {
	_, newCtx, span := getPayoutRequestTracerSpan(ctx, ".Create")
	defer span.End()

	err := s.r.Create(newCtx, payout)
	if err != nil {
		return err
	}

	return nil
}

func (s *payoutRequestService) Archive(ctx context.Context, id uuid.UUID) error {
	_, newCtx, span := getPayoutRequestTracerSpan(ctx, ".Archive")
	defer span.End()

	err := s.r.Archive(newCtx, id)
	if err != nil {
		return err
	}

	return nil
}
