package mongo

import (
	"context"
	"fmt"
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

func (r LessonRepository) spanName() string {
	return fmt.Sprintf("courses/repository/mongo/%s.", r.lesson.Name())
}

func (r LessonRepository) Create(ctx context.Context, lesson *domain.Lesson) (*domain.Lesson, error) {
	logger := zapctx.Logger(ctx)

	tr := global.Tracer(domain.ServiceName)
	newCtx, span := tr.Start(ctx, r.spanName()+"Create")
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
	newCtx, span := tr.Start(ctx, r.spanName()+"Get")
	defer span.End()

	var lesson models.Lesson
	filter := createUUIDFilter(lessonID, "lesson_id")
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
		appErr := xapperror.New(http.StatusInternalServerError, "can't convert lesson", "error converting lesson to domain", err)
		return nil, appErr
	}

	return res, nil
}

func (r LessonRepository) GetAllByCourse(ctx context.Context, courseID uuid.UUID) ([]*domain.Lesson, error) {
	logger := zapctx.Logger(ctx)

	tr := global.Tracer(domain.ServiceName)
	newCtx, span := tr.Start(ctx, r.spanName()+"GetAllByCourse")
	defer span.End()

	var lessons []models.Lesson
	filter := createUUIDFilter(courseID, "course_id")
	cursor, err := r.lesson.Find(newCtx, filter)
	if mErr := handleMongoError(err, logger); mErr != nil {
		return nil, mErr
	}
	if err != nil {
		appErr := xapperror.New(http.StatusNotFound, "can't find lessons", "error fetching lessons from MongoDB", err)
		return nil, appErr
	}

	if err := cursor.All(newCtx, &lessons); err != nil {
		appErr := xapperror.New(http.StatusInternalServerError, "can't fetch lessons",
			"error fetching lessons from MongoDB", err)
		return nil, appErr
	}

	var res []*domain.Lesson
	for _, l := range lessons {
		d, err := l.ToDomain()
		if err != nil {
			appErr := xapperror.New(http.StatusInternalServerError, "can't convert lesson",
				"error converting lesson to domain", err)
			return nil, appErr
		}
		res = append(res, d)
	}

	return res, nil
}

func (r LessonRepository) Update(ctx context.Context, lesson *domain.Lesson) error {
	logger := zapctx.Logger(ctx)

	tr := global.Tracer(domain.ServiceName)
	newCtx, span := tr.Start(ctx, r.spanName()+"Update")
	defer span.End()

	mongoLesson := models.ToMongoModelLesson(*lesson)
	filter := createUUIDFilter(lesson.ID, "lesson_id")
	_, err := r.lesson.ReplaceOne(newCtx, filter, mongoLesson)
	if mErr := handleMongoError(err, logger); mErr != nil {
		return mErr
	}
	if err != nil {
		appErr := xapperror.New(http.StatusNotFound, "can't update lesson", "error updating lesson in MongoDB", err)
		return appErr
	}

	return nil
}

func (r LessonRepository) Delete(ctx context.Context, lessonID uuid.UUID) error {
	logger := zapctx.Logger(ctx)

	tr := global.Tracer(domain.ServiceName)
	newCtx, span := tr.Start(ctx, r.spanName()+"Delete")
	defer span.End()

	filter := createUUIDFilter(lessonID, "lesson_id")
	_, err := r.lesson.DeleteOne(newCtx, filter)
	if mErr := handleMongoError(err, logger); mErr != nil {
		return mErr
	}
	if err != nil {
		appErr := xapperror.New(http.StatusNotFound, "can't delete lesson", "error deleting lesson from MongoDB", err)
		return appErr
	}

	return nil
}
