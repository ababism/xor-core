package service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"time"
	xorerror "xor-go/pkg/error"
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

func (r *accountService) Create(ctx context.Context, registerAccount *domain.RegisterAccount) error {
	loginPresent, err := r.accountRepository.LoginPresent(ctx, registerAccount.Login)
	if err != nil {
		return err
	}
	if loginPresent {
		return xorerror.NewIllegalArgumentError("login already exists")
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
	}
	fmt.Println(account.Uuid)
	return r.accountRepository.Create(ctx, account)
}

func (r *accountService) UpdatePassword(ctx context.Context, uuid uuid.UUID, password string) error {
	passwordHash := hash.CreateHash(password, passwordSalt)
	account, err := r.accountRepository.Get(ctx, uuid)
	if err != nil {
		return err
	}
	if passwordHash == account.PasswordHash {
		return xorerror.NewIllegalArgumentError("old and new passwords are the same")
	}
	return r.accountRepository.UpdatePassword(ctx, uuid, passwordHash)
}
