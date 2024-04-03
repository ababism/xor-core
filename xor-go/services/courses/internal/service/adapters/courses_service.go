package adapters

import (
	"context"
	"github.com/google/uuid"
	"xor-go/services/courses/internal/domain"
)

type CoursesService interface {
	CourseService
	LessonService
	UserService
	PublicationRequestService
}

type UserService interface {
	BuyCourse(ctx context.Context, studentID uuid.UUID, courseID uuid.UUID) error

	RegisterStudentProfile(ctx context.Context, accountID uuid.UUID, profile domain.Student) error
	RegisterTeacherProfile(ctx context.Context, accountID uuid.UUID, profile domain.Teacher) error
	//isLessonAccessible(ctx context.Context, userID uuid.UUID, lessonID uuid.UUID) (bool, error)
	//IsCourseAccessible(ctx context.Context, userID uuid.UUID, lessonID uuid.UUID) (bool, error)

	ChangeCourseAccess(ctx context.Context, teacherID uuid.UUID, courseID uuid.UUID) error
}

type CourseService interface {
	// Teacher
	CreateCourse(ctx context.Context, teacherID uuid.UUID, course *domain.Course) (uuid.UUID, error)
	GetCourse(ctx context.Context, courseID uuid.UUID) (*domain.Course, error) // hide fields if role == user
	UpdateCourse(ctx context.Context, courseID uuid.UUID, course *domain.Course) error
	DeleteCourse(ctx context.Context, courseID uuid.UUID) error

	// moderator from request
	// TODO
	//publishCourse(ctx context.Context, courseID uuid.UUID) error

	//registerCourseFeedback(ctx context.Context, courseID uuid.UUID) (feedbackID uuid.UUID, err error)
	//registerProducts(ctx context.Context, courseID uuid.UUID) error

	// Finance system
	ConfirmAccess(ctx context.Context, buyerID uuid.UUID, productIDs []uuid.UUID) error
}

// LessonService represents the service interface for managing lessons.
type LessonService interface {
	// User
	ReadLesson(ctx context.Context, lessonID uuid.UUID) (*domain.Lesson, error)

	// Teacher
	CreateLesson(ctx context.Context, lesson *domain.Lesson) error
	GetLesson(ctx context.Context, lessonID uuid.UUID) (*domain.Lesson, error)
	UpdateLesson(ctx context.Context, lesson *domain.Lesson) error
	DeleteLesson(ctx context.Context, lessonID uuid.UUID) error

	// TODO
	//registerLessonFeedback(ctx context.Context, lessonIDs []uuid.UUID) (feedbackIDs []uuid.UUID, err error)
	//registerLessonProduct(ctx context.Context, products []*domain.Lesson) (productIDs []uuid.UUID, err error)
}

type PublicationRequestService interface {
	// Teacher
	RequestCoursePublication(ctx context.Context, courseID uuid.UUID) (domain.PublicationRequest, error)
	// Moderator
	UpdatePublicationRequest(ctx context.Context, requestID domain.PublicationRequest) error
}
