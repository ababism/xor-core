package mongo

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"xor-go/services/courses/internal/domain"
	"xor-go/services/courses/internal/repository/mongo/collections"
	"xor-go/services/courses/internal/service/adapters"
)

var _ adapters.StudentRepository = &StudentRepository{}

func NewStudentRepository(database *Database) *StudentRepository {
	courseCollection := database.database.Collection(collections.StudentCollectionName.String())

	return &StudentRepository{
		db:             database,
		course:         courseCollection,
		collectionName: collections.StudentCollectionName,
	}
}

type StudentRepository struct {
	db *Database

	collectionName collections.CollectionName
	course         *mongo.Collection
}

func (sr StudentRepository) GetLessonAccess(ctx context.Context, userID uuid.UUID, lessonID uuid.UUID) (domain.LessonAccess, error) {
	//TODO implement me
	panic("implement me")
}

func (sr StudentRepository) CreateAccessToLesson(ctx context.Context, lessonAccess domain.LessonAccess) (domain.LessonAccess, error) {
	//TODO implement me
	panic("implement me")
}

func (sr StudentRepository) Create(ctx context.Context, profile domain.Student) error {
	//TODO implement me
	panic("implement me")
}
