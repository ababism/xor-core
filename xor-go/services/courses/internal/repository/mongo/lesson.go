package mongo

import (
	"context"
	"github.com/google/uuid"
	"github.com/juju/zaputil/zapctx"
	"go.mongodb.org/mongo-driver/mongo"
	global "go.opentelemetry.io/otel"
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
	newCtx, span := tr.Start(ctx, "courses/repository/mongo/lesson.Create")
	defer span.End()

	mongoLesson := models.ToMongoModelLesson(*lesson)
	_, err := r.lesson.InsertOne(newCtx, mongoLesson)
	if mErr := handleMongoError(err, logger); mErr != nil {
		return nil, mErr
	}
	if err != nil {
		appErr := xapperror.New(http.StatusBadRequest, "lesson already created",
			"failed to create lesson in MongoDB", err)
		return nil, appErr
	}

	return lesson, nil
}

func (r LessonRepository) Get(ctx context.Context, lessonID uuid.UUID) (*domain.Lesson, error) {
	logger := zapctx.Logger(ctx)

	tr := global.Tracer(domain.ServiceName)
	newCtx, span := tr.Start(ctx, "courses/repository/mongo/lesson.Get")
	defer span.End()

	var lesson models.Lesson
	filter := createIDFilter(lessonID, "lesson_id")
	err := r.lesson.FindOne(newCtx, filter).Decode(&lesson)
	if mErr := handleMongoError(err, logger); mErr != nil {
		return nil, mErr
	}
	if err != nil {
		appErr := xapperror.New(http.StatusNotFound, "can't find lesson", "error fetching lesson from MongoDB", err)
		return nil, appErr
	}

	res, err := lesson.ToDomain()
	if err != nil {
		return nil, err
	}

	return res, nil
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
