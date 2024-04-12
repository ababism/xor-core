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

var _ adapters.TeacherRepository = &TeacherRepository{}

func NewTeacherRepository(database *Database) *TeacherRepository {
	teacherCollection := database.database.Collection(collections.TeacherCollectionName.String())
	courseCollection := database.database.Collection(collections.CourseEditorCollectionName.String())

	return &TeacherRepository{
		db:      database,
		teacher: teacherCollection,
		course:  courseCollection,
	}
}

type TeacherRepository struct {
	db *Database

	teacher *mongo.Collection
	course  *mongo.Collection
}

func (r TeacherRepository) Get(ctx context.Context, teacherID uuid.UUID) (domain.Teacher, error) {
	logger := zapctx.Logger(ctx)

	tr := global.Tracer(domain.ServiceName)
	newCtx, span := tr.Start(ctx, "courses/repository/mongo/teacher.Get")
	defer span.End()

	filter := bson.M{"account_id": teacherID.String()}
	var teacher models.Teacher
	err := r.teacher.FindOne(newCtx, filter).Decode(&teacher)
	if mErr := handleMongoError(err, logger); mErr != nil {
		return domain.Teacher{}, mErr
	}

	res, err := teacher.ToDomain()
	if err != nil {
		return domain.Teacher{}, xapperror.New(http.StatusInternalServerError, "internal server error",
			"failed to convert teacher to domain", err)
	}
	return *res, nil
}

func (r TeacherRepository) Create(ctx context.Context, tProfile domain.Teacher) error {
	logger := zapctx.Logger(ctx)

	tr := global.Tracer(domain.ServiceName)
	newCtx, span := tr.Start(ctx, "courses/repository/mongo/publication.Create")
	defer span.End()

	teacherModel := models.ToMongoModelTeacher(tProfile)
	_, err := r.teacher.InsertOne(newCtx, teacherModel)
	if mErr := handleMongoError(err, logger); mErr != nil {
		return mErr
	}
	if err == nil {
		appErr := xapperror.New(http.StatusBadRequest, "teacher profile already created",
			"failed to create teacher profile in MongoDB", err)
		return appErr
	}

	return nil
}

func teacherCourseAccessFilter(teacherID uuid.UUID, courseID uuid.UUID) bson.M {
	return bson.M{
		"teacher_id": teacherID.String(),
		"course_id":  courseID.String(),
	}
}

// IsCourseAccessible	checks if r.course collection has a document with teacherID and courseID
func (r TeacherRepository) IsCourseAccessible(ctx context.Context, teacherID uuid.UUID, courseID uuid.UUID) (bool, error) {
	logger := zapctx.Logger(ctx)

	tr := global.Tracer(domain.ServiceName)
	newCtx, span := tr.Start(ctx, "courses/repository/mongo/teacher.IsCourseAccessible")
	defer span.End()

	filter := teacherCourseAccessFilter(teacherID, courseID)
	count, err := r.course.CountDocuments(newCtx, filter)
	if err != nil {
		mErr := handleMongoError(err, logger)
		if mErr != nil {
			return false, mErr
		}
		return false, xapperror.New(http.StatusInternalServerError, "internal server error",
			"failed to count documents in course collection", err)
	}

	return count > 0, nil
}

//func (r TeacherRepository) CreateAccessToCourse(ctx context.Context, teacherID uuid.UUID, courseID uuid.UUID) error {
//	//TODO implement me
//	panic("implement me")
//}
