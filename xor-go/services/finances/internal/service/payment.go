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
	spanDefaultPayment = "payment/service."
)

var _ adapters.PaymentService = &paymentService{}

type paymentService struct {
	r adapters.PaymentRepository
	c adapters.PaymentsClient
}

func NewPaymentService(
	paymentRepository adapters.PaymentRepository,
	paymentsClient adapters.PaymentsClient,
) adapters.PaymentService {
	return &paymentService{r: paymentRepository, c: paymentsClient}
}

func getPaymentTracerSpan(ctx context.Context, name string) (trace.Tracer, context.Context, trace.Span) {
	tr := global.Tracer(adapters.ServiceNamePayment)
	newCtx, span := tr.Start(ctx, spanDefaultPayment+name)
	return tr, newCtx, span
}

func (s *paymentService) Get(ctx context.Context, uuid uuid.UUID) (*domain.PaymentGet, error) {
	_, newCtx, span := getPaymentTracerSpan(ctx, ".Get")
	defer span.End()

	filter := domain.CreatePaymentFilterId(&uuid)
	payment, err := s.r.Get(newCtx, &filter)
	if err != nil {
		return nil, err
	}

	return payment, nil
}

func (s *paymentService) List(ctx context.Context, filter *domain.PaymentFilter) ([]domain.PaymentGet, error) {
	_, newCtx, span := getPaymentTracerSpan(ctx, ".List")
	defer span.End()

	payments, err := s.r.List(newCtx, filter)
	if err != nil {
		return nil, err
	}

	return payments, nil
}

func (s *paymentService) Create(ctx context.Context, payment *domain.PaymentCreate) error {
	_, newCtx, span := getPaymentTracerSpan(ctx, ".Create")
	defer span.End()

	err := s.r.Create(newCtx, payment)
	if err != nil {
		return err
	}

	return nil
}
