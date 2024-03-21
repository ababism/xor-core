package service

import (
	"context"
	global "go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"xor-go/services/finances/internal/domain"
	"xor-go/services/finances/internal/service/adapters"
)

const (
	spanDefaultBankAccount = "bank-account/service."
)

var _ adapters.BankAccountService = &bankAccountService{}

type bankAccountService struct {
	r adapters.BankAccountRepository
}

func NewAccountService(accountRepository adapters.BankAccountRepository) adapters.BankAccountService {
	return &bankAccountService{r: accountRepository}
}

func getAccountTracerSpan(ctx context.Context, name string) (trace.Tracer, context.Context, trace.Span) {
	tr := global.Tracer(adapters.ServiceNameBankAccount)
	newCtx, span := tr.Start(ctx, spanDefaultBankAccount+name)
	return tr, newCtx, span
}

func (s *bankAccountService) Get(ctx context.Context, login string) (*domain.BankAccountGet, error) {
	_, newCtx, span := getAccountTracerSpan(ctx, ".Get")
	defer span.End()

	filter := domain.CreateBankAccountFilterLogin(&login)
	account, err := s.r.Get(newCtx, &filter)
	if err != nil {
		return nil, err
	}

	return account, err
}

func (s *bankAccountService) List(ctx context.Context, filter *domain.BankAccountFilter) ([]domain.BankAccountGet, error) {
	_, newCtx, span := getAccountTracerSpan(ctx, ".List")
	defer span.End()

	accounts, err := s.r.List(newCtx, filter)
	if err != nil {
		return nil, err
	}

	return accounts, err
}

func (s *bankAccountService) Create(ctx context.Context, account *domain.BankAccountCreate) error {
	_, newCtx, span := getAccountTracerSpan(ctx, ".Create")
	defer span.End()

	err := s.r.Create(newCtx, account)
	if err != nil {
		return err
	}

	return err
}

func (s *bankAccountService) Update(ctx context.Context, account *domain.BankAccountUpdate) error {
	_, newCtx, span := getAccountTracerSpan(ctx, ".Update")
	defer span.End()

	err := s.r.Update(newCtx, account)
	if err != nil {
		return err
	}

	return err
}

func (s *bankAccountService) ChangeFunds(ctx context.Context, login string, diff float32) error {
	_, newCtx, span := getAccountTracerSpan(ctx, ".AddDiffToFunds")
	defer span.End()

	account, err := s.Get(ctx, login)
	if err != nil {
		return err
	}

	account.Funds += diff

	err = s.r.Update(newCtx, domain.GetToBankAccountUpdateDomain(account))
	if err != nil {
		return err
	}

	return err
}
