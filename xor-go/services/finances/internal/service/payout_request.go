package service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	global "go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"time"
	"xor-go/services/finances/internal/domain"
	"xor-go/services/finances/internal/log"
	"xor-go/services/finances/internal/service/adapters"
)

const (
	spanDefaultPayoutRequest = "payout-request/service."
)

var _ adapters.PayoutRequestService = &payoutRequestService{}

type payoutRequestService struct {
	rPayout      adapters.PayoutRequestRepository
	rPayment     adapters.PaymentRepository
	rBankAccount adapters.BankAccountService
	cPayments    adapters.PaymentsClient
}

func NewPayoutRequestService(
	payoutRequestRepo adapters.PayoutRequestRepository,
	paymentRepo adapters.PaymentRepository,
	paymentsClient adapters.PaymentsClient,
	rBankAccount adapters.BankAccountService,
) adapters.PayoutRequestService {
	return &payoutRequestService{
		rPayout:      payoutRequestRepo,
		rPayment:     paymentRepo,
		rBankAccount: rBankAccount,
		cPayments:    paymentsClient,
	}
}

func getPayoutRequestTracerSpan(ctx context.Context, name string) (trace.Tracer, context.Context, trace.Span) {
	tr := global.Tracer(adapters.ServiceNamePayoutRequest)
	newCtx, span := tr.Start(ctx, spanDefaultPayoutRequest+name)
	return tr, newCtx, span
}

func (s *payoutRequestService) Get(ctx context.Context, id uuid.UUID) (*domain.PayoutRequestGet, error) {
	_, newCtx, span := getPayoutRequestTracerSpan(ctx, ".GetByLogin")
	defer span.End()

	payoutRequest, err := s.rPayout.Get(newCtx, id)
	if err != nil {
		return nil, err
	}

	return payoutRequest, nil
}

func (s *payoutRequestService) List(
	ctx context.Context,
	filter *domain.PayoutRequestFilter,
) ([]domain.PayoutRequestGet, error) {
	_, newCtx, span := getPayoutRequestTracerSpan(ctx, ".List")
	defer span.End()

	payoutRequests, err := s.rPayout.List(newCtx, filter)
	if err != nil {
		return nil, err
	}

	return payoutRequests, nil
}

func (s *payoutRequestService) Create(ctx context.Context, payout *domain.PayoutRequestCreate) (*uuid.UUID, error) {
	_, newCtx, span := getPayoutRequestTracerSpan(ctx, ".Create")
	defer span.End()

	id, err := s.rPayout.Create(newCtx, payout)
	if err != nil {
		return nil, err
	}

	createPayout, err := s.cPayments.CreatePayout(ctx, &domain.PaymentsCreatePayout{
		PaymentUUID: *id,
		PaymentName: "",
		Money:       payout.Amount,
		Currency:    "RUB",
		FullName:    payout.Receiver.String(),
		Phone:       "",
		Email:       "",
		CardInfo:    domain.PaymentsCreatePayoutCard{},
		IsTest:      false,
	})
	if err != nil {
		return nil, err
	}
	log.Logger.Info(fmt.Sprintf("Payout created: %v", createPayout))

	time.Sleep(5 * time.Second)

	log.Logger.Info(fmt.Sprintf("Payout status: %v", createPayout.Status))
	if createPayout.Status == domain.PaymentsStatusSucceeded {
		err = s.rBankAccount.ChangeFunds(ctx, payout.Receiver, -1*payout.Amount)
		if err != nil {
			return nil, err
		}
	}

	return id, nil
}

func (s *payoutRequestService) Archive(ctx context.Context, id uuid.UUID) error {
	_, newCtx, span := getPayoutRequestTracerSpan(ctx, ".Archive")
	defer span.End()

	payout, err := s.rPayout.Get(ctx, id)
	if err != nil {
		return err
	}

	_, err = s.rPayment.Create(ctx, &domain.PaymentCreate{
		Sender:   uuid.Nil,
		Receiver: payout.Receiver,
		Data:     domain.PaymentData{},
		URL:      "",
		Status:   domain.STATUS_COMPLETED,
		EndedAt:  time.Now(),
	})
	if err != nil {
		return err
	}

	err = s.rPayout.Delete(newCtx, id)
	if err != nil {
		return err
	}

	return nil
}
