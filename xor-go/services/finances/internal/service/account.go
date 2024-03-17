package service

import (
	"context"
	global "go.opentelemetry.io/otel"
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

func (s *bankAccountService) GetByLogin(ctx context.Context, login string) (*domain.BankAccountGet, error) {
	tr := global.Tracer(adapters.ServiceNameBankAccount)
	newCtx, span := tr.Start(ctx, spanDefaultBankAccount+".GetByLogin")
	defer span.End()

	filter := domain.CreateBankAccountFilterLogin(&login)
	account, err := s.r.Get(newCtx, filter)
	if err != nil {
		return nil, err
	}

	return account, err
}

func (s *bankAccountService) List(ctx context.Context, filter domain.BankAccountFilter) ([]domain.BankAccountGet, error) {
	tr := global.Tracer(adapters.ServiceNameBankAccount)
	newCtx, span := tr.Start(ctx, spanDefaultBankAccount+".List")
	defer span.End()

	accounts, err := s.r.List(newCtx, filter)
	if err != nil {
		return nil, err
	}

	return accounts, err
}

func (s *bankAccountService) Create(ctx context.Context, account *domain.BankAccountGet) error {
	tr := global.Tracer(adapters.ServiceNameBankAccount)
	newCtx, span := tr.Start(ctx, spanDefaultBankAccount+".Create")
	defer span.End()

	err := s.r.Create(newCtx, account)
	if err != nil {
		return err
	}

	return err
}

func (s *bankAccountService) Update(ctx context.Context, account *domain.BankAccountGet) error {
	tr := global.Tracer(adapters.ServiceNameBankAccount)
	newCtx, span := tr.Start(ctx, spanDefaultBankAccount+".Update")
	defer span.End()

	err := s.r.Update(newCtx, account)
	if err != nil {
		return err
	}

	return err
}

func (s *bankAccountService) AddDiffToFunds(ctx context.Context, login string, diff float64) error {
	tr := global.Tracer(adapters.ServiceNameBankAccount)
	newCtx, span := tr.Start(ctx, spanDefaultBankAccount+".AddDiffToFunds")
	defer span.End()

	account, err := s.GetByLogin(ctx, login)
	if err != nil {
		return err
	}

	account.Funds += diff

	err = s.r.Update(newCtx, account)
	if err != nil {
		return err
	}

	return err
}
