package mongo

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"xor-go/services/courses/internal/domain"
	"xor-go/services/courses/internal/repository/mongo/collections"
	"xor-go/services/courses/internal/service/adapters"
)

var _ adapters.LessonRepository = &LessonRepository{}

func NewLessonRepository(database *Database, name collections.CollectionName) *LessonRepository {
	courseCollection := database.database.Collection(name.String())

	return &LessonRepository{
		db:     database,
		lesson: courseCollection,
	}
}

type LessonRepository struct {
	db *Database

	lesson *mongo.Collection
}

func (lr LessonRepository) Create(ctx context.Context, lesson *domain.Lesson) (*domain.Lesson, error) {
	//TODO implement me
	panic("implement me")
}

func (lr LessonRepository) Get(ctx context.Context, lessonID uuid.UUID) (*domain.Lesson, error) {
	//TODO implement me
	panic("implement me")
}

func (lr LessonRepository) GetAllByCourse(ctx context.Context, courseID uuid.UUID) ([]*domain.Lesson, error) {
	//TODO implement me
	panic("implement me")
}

func (lr LessonRepository) Update(ctx context.Context, lesson *domain.Lesson) error {
	//TODO implement me
	panic("implement me")
}

func (lr LessonRepository) Delete(ctx context.Context, lessonID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}
