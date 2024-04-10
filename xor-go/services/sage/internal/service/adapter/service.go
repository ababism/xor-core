package adapter

import "context"

type IdmService interface {
	Verify(ctx context.Context)
	//Get(ctx context.Context, login string) (*domain.BankAccountGet, error)
	//List(ctx context.Context, filter *domain.BankAccountFilter) ([]domain.BankAccountGet, error)
	//Create(ctx context.Context, account *domain.BankAccountCreate) error
	//Update(ctx context.Context, account *domain.BankAccountUpdate) error
	//ChangeFunds(ctx context.Context, login string, newFunds float32) error
}
