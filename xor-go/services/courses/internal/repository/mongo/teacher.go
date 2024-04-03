package mongo

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"xor-go/services/courses/internal/domain"
	"xor-go/services/courses/internal/repository/mongo/models"
	"xor-go/services/courses/internal/service/adapters"
)

var _ adapters.TeacherRepository = &TeacherRepository{}

func NewTeacherRepository(database *Database) *TeacherRepository {
	courseCollection := database.db.Collection(models.TeacherCollectionName)

	return &TeacherRepository{
		db:     database,
		course: courseCollection,
	}
}

type TeacherRepository struct {
	db *Database

	course *mongo.Collection
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
