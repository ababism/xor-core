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
	studentCollection := database.database.Collection(collections.StudentCollectionName.String())
	lessonAccessCollection := database.database.Collection(collections.LessonAccessCollectionName.String())

	return &StudentRepository{
		db:           database,
		student:      studentCollection,
		lessonAccess: lessonAccessCollection,
	}
}

type StudentRepository struct {
	db *Database

	student      *mongo.Collection
	lessonAccess *mongo.Collection
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
