package mongo

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"xor-go/services/courses/internal/domain"
	"xor-go/services/courses/internal/repository/mongo/models"
	"xor-go/services/courses/internal/service/adapters"
)

var _ adapters.PublicationRequestRepository = &PublicationRequestRepository{}

func NewPublicationRepository(database *Database) *PublicationRequestRepository {
	courseCollection := database.database.Collection(models.PublicationRequestsCollectionName)

	return &PublicationRequestRepository{
		db:     database,
		course: courseCollection,
	}
}

type PublicationRequestRepository struct {
	db *Database

	course *mongo.Collection
}

func (pr PublicationRequestRepository) Create(ctx context.Context, req domain.PublicationRequest) error {
	//TODO implement me
	panic("implement me")
}

func (pr PublicationRequestRepository) Get(ctx context.Context, reqID uuid.UUID) (*domain.PublicationRequest, error) {
	//TODO implement me
	panic("implement me")
}

func (pr PublicationRequestRepository) GetAll(ctx context.Context, offset, limit int) ([]domain.PublicationRequest, error) {
	//TODO implement me
	panic("implement me")
}

func (pr PublicationRequestRepository) GetAllFrom(ctx context.Context, teacher int, offset, limit int) ([]domain.PublicationRequest, error) {
	//TODO implement me
	panic("implement me")
}

func (pr PublicationRequestRepository) Update(ctx context.Context, req domain.PublicationRequest) error {
	//TODO implement me
	panic("implement me")
}
