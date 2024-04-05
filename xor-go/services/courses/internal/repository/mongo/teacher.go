package mongo

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"xor-go/services/courses/internal/domain"
	"xor-go/services/courses/internal/repository/mongo/collections"
	"xor-go/services/courses/internal/service/adapters"
)

var _ adapters.TeacherRepository = &TeacherRepository{}

func NewTeacherRepository(database *Database) *TeacherRepository {
	courseCollection := database.database.Collection(collections.TeacherCollectionName.String())

	return &TeacherRepository{
		db:             database,
		course:         courseCollection,
		collectionName: collections.TeacherCollectionName,
	}
}

type TeacherRepository struct {
	db *Database

	collectionName collections.CollectionName
	course         *mongo.Collection
}

func (tr TeacherRepository) Create(ctx context.Context, profile domain.Teacher) error {
	//TODO implement me
	panic("implement me")
}

func (tr TeacherRepository) IsCourseAccessible(ctx context.Context, teacherID uuid.UUID, courseID uuid.UUID) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (tr TeacherRepository) CreateAccessToCourse(ctx context.Context, teacherID uuid.UUID, courseID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}
