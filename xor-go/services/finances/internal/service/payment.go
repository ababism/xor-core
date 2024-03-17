package service

import (
	"context"
	"errors"
	"github.com/google/uuid"
	global "go.opentelemetry.io/otel"
	"xor-go/services/finances/internal/domain"
	"xor-go/services/finances/internal/service/adapters"
)

const (
	spanDefaultPayment = "payment/service."
)

var _ adapters.PaymentService = &paymentService{}

type paymentService struct {
	r adapters.PaymentRepository
}

func NewPaymentService(paymentRepository adapters.PaymentRepository) adapters.PaymentService {
	return &paymentService{r: paymentRepository}
}

func (s *paymentService) Get(ctx context.Context, uuid uuid.UUID) (*domain.Payment, error) {
	tr := global.Tracer(adapters.ServiceNamePayment)
	newCtx, span := tr.Start(ctx, spanDefaultPayment+".Get")
	defer span.End()

	filter := domain.CreatePaymentFilterId(&uuid)
	payment, err := s.r.Get(newCtx, filter)
	if err != nil {
		return nil, err
	}

	return payment, nil
}

func (s *paymentService) List(ctx context.Context, filter domain.PaymentFilter) ([]domain.Payment, error) {
	tr := global.Tracer(adapters.ServiceNamePayment)
	newCtx, span := tr.Start(ctx, spanDefaultPayment+".List")
	defer span.End()

	payments, err := s.r.List(newCtx, filter)
	if err != nil {
		return nil, err
	}

	return payments, nil
}

func (s *paymentService) Create(ctx context.Context, payment *domain.Payment) error {
	tr := global.Tracer(adapters.ServiceNamePayment)
	newCtx, span := tr.Start(ctx, spanDefaultPayment+".Create")
	defer span.End()

	if payment.UUID == uuid.Nil {
		return errors.New("payment UUID cannot be nil")
	}

	err := s.r.Create(newCtx, payment)
	if err != nil {
		return err
	}

	return nil
}

func (s *paymentService) Update(ctx context.Context, payment *domain.Payment) error {
	tr := global.Tracer(adapters.ServiceNamePayment)
	newCtx, span := tr.Start(ctx, spanDefaultPayment+".Update")
	defer span.End()

	if payment.UUID == uuid.Nil {
		return errors.New("payment UUID cannot be nil")
	}

	err := s.r.Update(newCtx, payment)
	if err != nil {
		return err
	}

	return nil
}
