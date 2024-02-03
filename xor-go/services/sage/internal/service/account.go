package service

import (
	"context"
	"github.com/google/uuid"
	"time"
	"xor-go/pkg/xorerror"
	"xor-go/services/sage/internal/domain"
	"xor-go/services/sage/internal/repository"
	"xor-go/services/sage/pkg/auth"
)

const (
	passwordSaltSize = 16
)

var _ AccountService = &accountService{}

type accountService struct {
	accountRepository repository.AccountRepository
}

func NewAccountService(accountRepository repository.AccountRepository) AccountService {
	return &accountService{accountRepository: accountRepository}
}

func (r *accountService) Create(ctx context.Context, registerAccount *domain.RegisterAccount) error {
	loginPresent, err := r.accountRepository.LoginPresent(ctx, registerAccount.Login)
	if err != nil {
		return err
	}
	if loginPresent {
		return xorerror.NewIllegalArgumentError("login already exists")
	}

	salt, err := hash.CreateSalt(passwordSaltSize)
	if err != nil {
		return err
	}
	passwordHash := hash.CreateHash(registerAccount.Password, salt)
	createdAt := time.Now()
	account := &domain.Account{
		Uuid:         uuid.New(),
		Login:        registerAccount.Login,
		PasswordHash: passwordHash,
		CreatedAt:    createdAt,
		UpdatedAt:    createdAt,
		Contacts:     nil,
	}
	return r.accountRepository.Create(ctx, account)
}

func (r *accountService) UpdatePassword(ctx context.Context, uuid uuid.UUID, password string) error {
	salt, err := hash.CreateSalt(passwordSaltSize)
	if err != nil {
		return err
	}
	passwordHash := hash.CreateHash(password, salt)
	previousPasswordHash, err := r.accountRepository.GetPasswordHash(ctx, uuid)
	if err != nil {
		return err
	}
	if passwordHash == previousPasswordHash {
		return xorerror.NewIllegalArgumentError("old and new passwords are the same")
	}
	return r.accountRepository.UpdatePassword(ctx, uuid, passwordHash)
}
