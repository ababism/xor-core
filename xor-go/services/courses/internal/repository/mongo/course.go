package mongo

import (
	"context"
	"github.com/google/uuid"
	"github.com/juju/zaputil/zapctx"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	global "go.opentelemetry.io/otel"
	"go.uber.org/zap"
	"net/http"
	"xor-go/pkg/xapperror"
	"xor-go/services/courses/internal/domain"
	"xor-go/services/courses/internal/repository/mongo/collections"
	"xor-go/services/courses/internal/repository/mongo/models"
	"xor-go/services/courses/internal/service/adapters"
)

var _ adapters.CourseRepository = &CourseRepository{}

func NewCourseRepository(database *Database, name collections.CollectionName) *CourseRepository {
	courseCollection := (*database).database.Collection(name.String())

	return &CourseRepository{
		db:     database,
		course: courseCollection,
	}
}

type CourseRepository struct {
	db *Database

	course *mongo.Collection
}

func (cr CourseRepository) Create(ctx context.Context, course *domain.Course) (*domain.Course, error) {
	logger := zapctx.Logger(ctx)

	tr := global.Tracer(domain.ServiceName)
	newCtx, span := tr.Start(ctx, "courses/repository/mongo/course.Create")
	defer span.End()

	mongoCourse := models.ToMongoModelCourse(course)
	insertRes, err := cr.course.InsertOne(newCtx, mongoCourse)
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

func (cr CourseRepository) Get(ctx context.Context, courseID uuid.UUID) (*domain.Course, error) {
	logger := zapctx.Logger(ctx)

	tr := global.Tracer(domain.ServiceName)
	newCtx, span := tr.Start(ctx, "courses/repository/mongo/course.Get")
	defer span.End()

	var course models.Course
	filter := createUUIDFilter(courseID, "course_id")
	err := cr.course.FindOne(newCtx, filter).Decode(&course)
	if mErr := handleMongoError(err, logger); mErr != nil {
		return nil, mErr
	}
	if err != nil {
		appErr := xapperror.New(http.StatusNotFound, "can't find course", "error fetching course from MongoDB", err)
		return nil, appErr
	}

	res, err := course.ToDomain()
	if err != nil {
		appErr := xapperror.New(http.StatusNotFound, "can't parse course", "error converting course from MongoDB to domain", err)
		return nil, appErr
	}

	return res, nil
}

func (cr CourseRepository) GetAll(ctx context.Context, offset, limit int) ([]*domain.Course, error) {
	logger := zapctx.Logger(ctx)

	tr := global.Tracer(domain.ServiceName)
	newCtx, span := tr.Start(ctx, "courses/repository/mongo/course.GetAll")
	defer span.End()

	var courses []models.Course
	opts := options.Find().SetSkip(int64(offset)).SetLimit(int64(limit))
	cursor, err := cr.course.Find(newCtx, bson.D{{}}, opts)
	if mErr := handleMongoError(err, logger); mErr != nil {
		return nil, mErr
	}
	if err != nil {
		appErr := xapperror.New(http.StatusNotFound, "can't find courses", "error fetching courses from MongoDB", err)
		return nil, appErr
	}
	if err := cursor.All(newCtx, &courses); err != nil {
		appErr := xapperror.New(http.StatusInternalServerError, "can't fetch courses", "error fetching courses from MongoDB", err)
		return nil, appErr
	}

	var res []*domain.Course
	for _, course := range courses {
		c, err := course.ToDomain()
		if err != nil {
			appErr := xapperror.New(http.StatusInternalServerError, "can't parse course", "error converting course from MongoDB to domain", err)
			return nil, appErr
		}
		res = append(res, c)
	}

	return res, nil
}

func (cr CourseRepository) GetAllByTeacher(ctx context.Context, teacherID uuid.UUID, offset, limit int) ([]*domain.Course, error) {
	logger := zapctx.Logger(ctx)

	tr := global.Tracer(domain.ServiceName)
	newCtx, span := tr.Start(ctx, "courses/repository/mongo/course.GetAllByTeacher")
	defer span.End()

	var courses []models.Course
	filter := createUUIDFilter(teacherID, "teacher_id")
	opts := options.Find().SetSkip(int64(offset)).SetLimit(int64(limit))
	cursor, err := cr.course.Find(newCtx, filter, opts)
	if mErr := handleMongoError(err, logger); mErr != nil {
		return nil, mErr
	}
	if err != nil {
		appErr := xapperror.New(http.StatusNotFound, "can't find courses", "error fetching courses from MongoDB", err)
		return nil, appErr
	}
	if err := cursor.All(newCtx, &courses); err != nil {
		appErr := xapperror.New(http.StatusInternalServerError, "can't fetch courses", "error fetching courses from MongoDB", err)
		return nil, appErr
	}

	var res []*domain.Course
	for _, course := range courses {
		c, err := course.ToDomain()
		if err != nil {
			appErr := xapperror.New(http.StatusInternalServerError, "can't parse course", "error converting course from MongoDB to domain", err)
			return nil, appErr
		}
		res = append(res, c)
	}

	return res, nil
}

func (cr CourseRepository) Update(ctx context.Context, courseID uuid.UUID, course *domain.Course) error {
	logger := zapctx.Logger(ctx)

	tr := global.Tracer(domain.ServiceName)
	newCtx, span := tr.Start(ctx, "courses/repository/mongo/course.Update")
	defer span.End()

	filter := createUUIDFilter(courseID, "course_id")
	update := bson.D{{Key: "$set", Value: models.ToMongoModelCourse(course)}}
	_, err := cr.course.UpdateOne(newCtx, filter, update)
	if mErr := handleMongoError(err, logger); mErr != nil {
		return mErr
	}
	if err != nil {
		appErr := xapperror.New(http.StatusNotFound, "can't update course", "error updating course in MongoDB", err)
		return appErr
	}

	return nil
}

func (cr CourseRepository) Delete(ctx context.Context, courseID uuid.UUID) error {
	logger := zapctx.Logger(ctx)

	tr := global.Tracer(domain.ServiceName)
	newCtx, span := tr.Start(ctx, "courses/repository/mongo/course.Delete")
	defer span.End()

	filter := createUUIDFilter(courseID, "course_id")
	_, err := cr.course.DeleteOne(newCtx, filter)
	if mErr := handleMongoError(err, logger); mErr != nil {
		return mErr
	}
	if err != nil {
		appErr := xapperror.New(http.StatusNotFound, "can't delete course", "error deleting course in MongoDB", err)
		return appErr
	}

	return nil
}

//func (cr CourseRepository) CreateWithTx(tx Session, ctx context.Context, course *domain.Course) (*domain.Course, error) {
//	newCr := NewCourseRepository(tx.DB)
//	return newCr.Create(ctx, course)
//}

//func (cr CourseRepository) WithTX(tx Session) *adapters.CourseRepository{
//	return NewCourseRepository(tx.DB)
//}
