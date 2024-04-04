package mongo

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"xor-go/services/courses/internal/domain"
	"xor-go/services/courses/internal/repository/mongo/models"
	"xor-go/services/courses/internal/service/adapters"
)

var _ adapters.StudentRepository = &StudentRepository{}

func NewStudentRepository(database *Database) *StudentRepository {
	courseCollection := database.database.Collection(models.StudentCollectionName)

	return &StudentRepository{
		db:     database,
		course: courseCollection,
	}
}

type StudentRepository struct {
	db *Database

	course *mongo.Collection
}

func (sr StudentRepository) Create(ctx context.Context, profile domain.Student) error {
	//TODO implement me
	panic("implement me")
}

func (sr StudentRepository) GetLessonAccessStatus(ctx context.Context, userID uuid.UUID, lessonID uuid.UUID) (domain.AccessStatus, error) {
	//TODO implement me
	panic("implement me")
}

func (sr StudentRepository) CreateAccessToLesson(ctx context.Context, userID uuid.UUID, lessonID uuid.UUID, status domain.AccessStatus) error {
	//TODO implement me
	panic("implement me")
}
