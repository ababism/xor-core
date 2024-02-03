package mongo

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"xor-go/services/sage/internal/domain"
	"xor-go/services/sage/internal/repository"
	"xor-go/services/sage/internal/repository/mongo/entity"
)

const (
	accountCollection = "account"
)

var _ repository.AccountRepository = &accountRepository{}

type accountRepository struct {
	logger            *zap.Logger
	accountCollection *mongo.Collection
}

func NewAccountMongoRepository(logger *zap.Logger, db *mongo.Database) repository.AccountRepository {
	accountCollection := db.Collection(accountCollection)
	return &accountRepository{logger: logger, accountCollection: accountCollection}
}

func (r *accountRepository) LoginPresent(ctx context.Context, login string) (bool, error) {
	count, err := r.accountCollection.CountDocuments(ctx, bson.M{"login": login})
	if err != nil {
		r.logger.Error("failed to check if login is presented", zap.Error(err))
		return false, err
	}
	return count > 0, nil
}

func (r *accountRepository) Get(ctx context.Context, uuid uuid.UUID) (*domain.Account, error) {
	var accountMongo entity.AccountMongo

	filter := bson.M{"uuid": uuid}
	err := r.accountCollection.FindOne(ctx, filter).Decode(&accountMongo)
	if err != nil {
		r.logger.Error("failed to get password_hash", zap.Error(err))
		return nil, err
	}
	return entity.ToAccount(&accountMongo), nil
}

func (r *accountRepository) Create(ctx context.Context, account *domain.Account) error {
	accountMongo := entity.ToAccountMongo(account)
	_, err := r.accountCollection.InsertOne(ctx, accountMongo)
	if err != nil {
		r.logger.Error("failed to create account", zap.Error(err))
		return err
	}
	return nil
}

func (r *accountRepository) UpdatePassword(ctx context.Context, uuid uuid.UUID, passwordHash string) error {
	filter := bson.M{"uuid": uuid}
	update := bson.M{"$set": bson.M{"password_hash": passwordHash}}
	_, err := r.accountCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		r.logger.Error("failed to update account password", zap.Error(err))
		return err
	}
	return nil
}
