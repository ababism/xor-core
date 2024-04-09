package mongo

import (
	"context"
	"github.com/google/uuid"
	"github.com/juju/zaputil/zapctx"
	"go.mongodb.org/mongo-driver/mongo"
	global "go.opentelemetry.io/otel"
	"go.uber.org/zap"
	"net/http"
	"xor-go/pkg/xapperror"
	"xor-go/services/courses/internal/domain"
	"xor-go/services/courses/internal/repository/mongo/collections"
	"xor-go/services/courses/internal/repository/mongo/models"
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

func (r LessonRepository) Create(ctx context.Context, lesson *domain.Lesson) (*domain.Lesson, error) {
	logger := zapctx.Logger(ctx)

	tr := global.Tracer(domain.ServiceName)
	newCtx, span := tr.Start(ctx, "courses/repository/mongo/course.Create")
	defer span.End()

	mongoLesson := models.ToMongoModelLesson(lesson)
	insertRes, err := r.lesson.InsertOne(newCtx, mongoLesson)
	if mErr := handleMongoError(err, logger); mErr != nil {
		return nil, mErr
	}
	if err != nil {
		appErr := xapperror.New(http.StatusBadRequest, "trip already created",
			"failed to create trip in MongoDB", err)
		return nil, appErr
	}

	insertedID, ok := insertRes.InsertedID.(string)
	if !ok {
		logger.Error("MongoDB id is not a string error", zap.Error(err))
		return nil, xapperror.New(http.StatusInternalServerError,
			"internal server error", "MongoDB _id is not a string", err)
	}
	resultID, err := uuid.Parse(insertedID)
	if err != nil {
		logger.Error("MongoDB id is not valid uuid", zap.Error(err))
		return nil, xapperror.New(http.StatusInternalServerError,
			"internal server error", "MongoDB _id is not a uuid", err)
	}
	course.ID = resultID
	return course, nil
}

func (r LessonRepository) Get(ctx context.Context, lessonID uuid.UUID) (*domain.Lesson, error) {
	//TODO implement me
	panic("implement me")
}

func (r LessonRepository) GetAllByCourse(ctx context.Context, courseID uuid.UUID) ([]*domain.Lesson, error) {
	//TODO implement me
	panic("implement me")
}

func (r LessonRepository) Update(ctx context.Context, lesson *domain.Lesson) error {
	//TODO implement me
	panic("implement me")
}

func (r LessonRepository) Delete(ctx context.Context, lessonID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}
