package mongo

import (
	"context"
	"github.com/google/uuid"
	"github.com/juju/zaputil/zapctx"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	global "go.opentelemetry.io/otel"
	"net/http"
	"xor-go/pkg/xapperror"
	"xor-go/services/courses/internal/domain"
	"xor-go/services/courses/internal/repository/mongo/collections"
	"xor-go/services/courses/internal/repository/mongo/models"
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

// GetLessonAccess returns a document from the r.lessonAccess collection that matches the userID and lessonID
func (r StudentRepository) GetLessonAccess(ctx context.Context, userID uuid.UUID, lessonID uuid.UUID) (domain.LessonAccess, error) {
	logger := zapctx.Logger(ctx)

	tr := global.Tracer(domain.ServiceName)
	newCtx, span := tr.Start(ctx, "courses/repository/mongo/student.GetLessonAccess")
	defer span.End()

	filter := bson.M{
		"student_id": userID.String(),
		"lesson_id":  lessonID.String(),
	}

	var lessonAccess models.LessonAccess
	err := r.lessonAccess.FindOne(newCtx, filter).Decode(&lessonAccess)
	if mErr := handleMongoError(err, logger); mErr != nil {
		return domain.LessonAccess{}, mErr
	}

	res, err := lessonAccess.ToDomain()
	if err != nil {
		return domain.LessonAccess{}, xapperror.New(http.StatusInternalServerError, "internal server error",
			"failed to convert lesson access to domain", err)
	}
	return res, nil
}

func (r StudentRepository) CreateAccessToLesson(ctx context.Context, lessonAccess domain.LessonAccess) (domain.LessonAccess, error) {
	logger := zapctx.Logger(ctx)

	tr := global.Tracer(domain.ServiceName)
	newCtx, span := tr.Start(ctx, "courses/repository/mongo/student.CreateAccessToLesson")
	defer span.End()

	lessonAccessModel := models.ToMongoModelLessonAccess(lessonAccess)
	_, err := r.lessonAccess.InsertOne(newCtx, lessonAccessModel)
	if mErr := handleMongoError(err, logger); mErr != nil {
		return domain.LessonAccess{}, mErr
	}
	if err != nil {
		appErr := xapperror.New(http.StatusBadRequest, "lesson access already created",
			"failed to create lesson access in MongoDB", err)
		return domain.LessonAccess{}, appErr
	}

	return lessonAccess, nil
}

func (r StudentRepository) Create(ctx context.Context, sProfile domain.Student) error {
	logger := zapctx.Logger(ctx)

	tr := global.Tracer(domain.ServiceName)
	newCtx, span := tr.Start(ctx, "courses/repository/mongo/student.Create")
	defer span.End()

	teacherModel := models.ToMongoModelStudent(sProfile)
	_, err := r.student.InsertOne(newCtx, teacherModel)
	if mErr := handleMongoError(err, logger); mErr != nil {
		return mErr
	}
	if err == nil {
		appErr := xapperror.New(http.StatusBadRequest, "student profile already created",
			"failed to create student profile in MongoDB", err)
		return appErr
	}

	return nil
}
