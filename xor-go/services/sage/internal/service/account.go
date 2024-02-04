package service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"time"
	"xor-go/pkg/xerror"
	"xor-go/services/sage/internal/domain"
	"xor-go/services/sage/internal/service/adapter"
	"xor-go/services/sage/pkg/auth"
)

var (
	// FIXME move to config
	passwordSalt = []byte("salt")
)

var _ adapter.AccountService = &accountService{}

type accountService struct {
	logger            *zap.Logger
	accountRepository adapter.AccountRepository
}

func NewAccountService(logger *zap.Logger, accountRepository adapter.AccountRepository) adapter.AccountService {
	return &accountService{logger: logger, accountRepository: accountRepository}
}

func (r *accountService) List(ctx context.Context, params *domain.AccountFilter) ([]domain.Account, error) {
	return r.accountRepository.List(ctx, params)
}

func (r *accountService) Create(ctx context.Context, registerAccount *domain.RegisterAccount) error {
	present, err := r.accountRepository.Present(ctx, &domain.AccountFilter{Login: &registerAccount.Login})
	if err != nil {
		return err
	}
	if present {
		return xerror.NewValueError("login already exists")
	}

	passwordHash := hash.CreateHash(registerAccount.Password, passwordSalt)
	createdAt := time.Now()
	account := &domain.Account{
		Uuid:         uuid.New(),
		Login:        registerAccount.Login,
		PasswordHash: passwordHash,
		CreatedAt:    createdAt,
		UpdatedAt:    createdAt,
		Contacts:     nil,
		Active:       true,
	}
	fmt.Println(account.Uuid)
	return r.accountRepository.Create(ctx, account)
}

func (r *accountService) UpdatePassword(ctx context.Context, uuid uuid.UUID, password string) error {
	present, err := r.accountRepository.Present(ctx, &domain.AccountFilter{Uuid: &uuid})
	if err != nil {
		return err
	}
	if !present {
		return xerror.NewValueError("account does not exist")
	}
	passwordHash := hash.CreateHash(password, passwordSalt)
	account, err := r.accountRepository.Get(ctx, &domain.AccountFilter{Uuid: &uuid})
	if err != nil {
		return err
	}
	if account == nil {
		return xerror.NewValueError("account does not exist")
	}
	if passwordHash == account.PasswordHash {
		return xerror.NewValueError("old and new passwords are the same")
	}
	return r.accountRepository.UpdatePassword(ctx, uuid, passwordHash)
}

func (r *accountService) Deactivate(ctx context.Context, uuid uuid.UUID) error {
	present, err := r.accountRepository.Present(ctx, &domain.AccountFilter{Uuid: &uuid})
	if !present {
		return xerror.NewValueError("uuid is not presented")
	}
	account, err := r.accountRepository.Get(ctx, &domain.AccountFilter{Uuid: &uuid})
	if err != nil {
		return err
	}
	if !account.Active {
		return xerror.NewValueError("account is already deactivated")
	}
	return r.accountRepository.Deactivate(ctx, uuid)
}
