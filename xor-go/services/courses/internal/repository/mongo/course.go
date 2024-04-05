package mongo

import (
	"context"
	"github.com/google/uuid"
	"github.com/juju/zaputil/zapctx"
	"go.mongodb.org/mongo-driver/mongo"
	global "go.opentelemetry.io/otel"
	"go.uber.org/zap"
	"net/http"
	"xor-go/pkg/apperror"
	"xor-go/services/courses/internal/domain"
	"xor-go/services/courses/internal/repository/mongo/collections"
	"xor-go/services/courses/internal/repository/mongo/models"
	"xor-go/services/courses/internal/service/adapters"
)

var _ adapters.CourseRepository = &CourseRepository{}

func NewCourseRepository(database *Database, name collections.CollectionName) *CourseRepository {
	courseCollection := database.database.Collection(name.String())

	return &CourseRepository{
		db:             database,
		course:         courseCollection,
		collectionName: name,
	}
}

type CourseRepository struct {
	db *Database

	collectionName collections.CollectionName
	course         *mongo.Collection
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
		appErr := apperror.New(http.StatusBadRequest, "trip already created",
			"failed to create trip in MongoDB", err)
		return nil, appErr
	}

	insertedID, ok := insertRes.InsertedID.(string)
	if !ok {
		logger.Error("MongoDB id is not a string error", zap.Error(err))
		return nil, apperror.New(http.StatusInternalServerError,
			"internal server error", "MongoDB _id is not a string", err)
	}
	resultID, err := uuid.Parse(insertedID)
	if err != nil {
		logger.Error("MongoDB id is not valid uuid", zap.Error(err))
		return nil, apperror.New(http.StatusInternalServerError,
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
	filter := createIDFilter(courseID)
	err := cr.course.FindOne(newCtx, filter).Decode(&course)
	if mErr := handleMongoError(err, logger); mErr != nil {
		return nil, mErr
	}
	if err != nil {
		appErr := apperror.New(http.StatusNotFound, "can't find course", "error fetching course from MongoDB", err)
		return nil, appErr
	}

	res := course.ToDomain()

	return &res, nil
}

func (cr CourseRepository) GetAll(ctx context.Context, offset, limit int) ([]*domain.Course, error) {
	//TODO implement me
	panic("implement me")
}

func (cr CourseRepository) GetAllByTeacher(ctx context.Context, teacherID uuid.UUID, offset, limit int) ([]*domain.Course, error) {
	//TODO implement me
	panic("implement me")
}

func (cr CourseRepository) Update(ctx context.Context, courseID uuid.UUID, course *domain.Course) error {
	//TODO implement me
	panic("implement me")
}

func (cr CourseRepository) Delete(ctx context.Context, courseID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

//func (cr CourseRepository) CreateWithTx(tx Session, ctx context.Context, course *domain.Course) (*domain.Course, error) {
//	newCr := NewCourseRepository(tx.DB)
//	return newCr.Create(ctx, course)
//}

//func (cr CourseRepository) WithTX(tx Session) *adapters.CourseRepository{
//	return NewCourseRepository(tx.DB)
//}
