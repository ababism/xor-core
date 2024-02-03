package repository

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"xor-go/services/sage/internal/model"
)

const (
	accountCollection = "account"
)

type AccountMongoRepository struct {
	logger            *zap.Logger
	accountCollection *mongo.Collection
}

func NewAccountMongoRepository(logger *zap.Logger, db *mongo.Database) *AccountMongoRepository {
	accountCollection := db.Collection(accountCollection)
	return &AccountMongoRepository{logger: logger, accountCollection: accountCollection}
}

func (r *AccountMongoRepository) LoginPresent(ctx context.Context, login string) (bool, error) {
	count, err := r.accountCollection.CountDocuments(ctx, bson.M{"login": login})
	if err != nil {
		r.logger.Error("failed to check if login is presented", zap.Error(err))
		return false, err
	}
	return count > 0, nil
}

func (r *AccountMongoRepository) GetPasswordHash(ctx context.Context, uuid uuid.UUID) (string, error) {
	var passwordHash string
	err := r.accountCollection.FindOne(ctx, bson.M{"uuid": uuid}).Decode(&passwordHash)
	if err != nil {
		r.logger.Error("failed to get password_hash")
		return "", err
	}
	return passwordHash, nil
}

func (r *AccountMongoRepository) Create(ctx context.Context, account *model.AccountEntity) error {
	_, err := r.accountCollection.InsertOne(ctx, account)
	if err != nil {
		r.logger.Error("failed to create account", zap.Error(err))
		return err
	}
	return nil
}

func (r *AccountMongoRepository) UpdatePassword(ctx context.Context, uuid uuid.UUID, passwordHash string) error {
	filter := bson.M{"uuid": uuid}
	update := bson.M{"$set": bson.M{"password_hash": passwordHash}}
	_, err := r.accountCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		r.logger.Error("failed to update account password", zap.Error(err))
		return err
	}
	return nil
}
