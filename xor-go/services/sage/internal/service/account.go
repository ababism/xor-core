package service

import (
	"context"
	"github.com/google/uuid"
	"time"
	"xor-go/pkg/xor_error"
	"xor-go/services/sage/internal/models"
	"xor-go/services/sage/internal/repository"
	"xor-go/services/sage/pkg/auth"
)

const (
	passwordSaltSize = 16
)

type AccountService struct {
	accountRepository repository.AccountRepositoryI
}

func NewAccountService(accountRepository repository.AccountRepositoryI) *AccountService {
	return &AccountService{accountRepository: accountRepository}
}

func (r *AccountService) Create(ctx context.Context, account *models.RegisterAccountEntity) error {
	loginPresent, err := r.accountRepository.LoginPresent(ctx, account.Login)
	if err != nil {
		return err
	}
	if loginPresent {
		return xor_error.NewIllegalArgumentError("login already exists")
	}

	salt, err := hash.CreateSalt(passwordSaltSize)
	if err != nil {
		return err
	}
	passwordHash := hash.CreateHash(account.Password, salt)
	createdAt := time.Now()
	accountEntity := &models.AccountEntity{
		Uuid:         uuid.New(),
		Login:        account.Login,
		PasswordHash: passwordHash,
		CreatedAt:    createdAt,
		UpdatedAt:    createdAt,
		Contacts:     nil,
	}
	return r.accountRepository.Create(ctx, accountEntity)
}

func (r *AccountService) UpdatePassword(ctx context.Context, uuid uuid.UUID, password string) error {
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
		return xor_error.NewIllegalArgumentError("old and new passwords are the same")
	}
	return r.accountRepository.UpdatePassword(ctx, uuid, passwordHash)
}
