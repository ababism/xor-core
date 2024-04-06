package adapters

import (
	"context"
	"github.com/google/uuid"
	"xor-go/services/courses/internal/domain"
)

type CoursesService interface {
	CourseService
	LessonService
	PublicationRequestService
	UserService
}

// CourseService represents the part of interface for courses logic.
type CourseService interface {
	// CreateCourse GetCourse UpdateCourse DeleteCourse — teacher courseCRUD
	CreateCourse(ctx context.Context, actor domain.Actor, course *domain.Course) (*domain.Course, error)
	GetCourse(ctx context.Context, actor domain.Actor, courseID uuid.UUID) (*domain.Course, error) // hide fields if role == user
	UpdateCourse(ctx context.Context, actor domain.Actor, courseID uuid.UUID, course *domain.Course) (*domain.Course, error)
	DeleteCourse(ctx context.Context, actor domain.Actor, courseID uuid.UUID) error

	// ReadCourse  — user gets published course (with visibility applied based on roles)
	ReadCourse(ctx context.Context, actor domain.Actor, courseID uuid.UUID) (*domain.Course, error)
}

// LessonService represents the part of interface for lesson logic.
type LessonService interface {
	// ReadLesson  user gets published course (with visibility applied)
	ReadLesson(ctx context.Context, actor domain.Actor, lessonID uuid.UUID) (*domain.Lesson, error)

	// CreateLesson GetLesson UpdateLesson DeleteLesson — teacher lessonCRUD
	CreateLesson(ctx context.Context, actor domain.Actor, lesson *domain.Lesson) (*domain.Lesson, error)
	GetLesson(ctx context.Context, actor domain.Actor, lessonID uuid.UUID) (*domain.Lesson, error)
	UpdateLesson(ctx context.Context, actor domain.Actor, lesson *domain.Lesson) (*domain.Lesson, error)
	DeleteLesson(ctx context.Context, actor domain.Actor, lessonID uuid.UUID) error
}

// PublicationRequestService represents the part of interface for publication of course and moderation.
type PublicationRequestService interface {
	// RequestCoursePublication Teacher requests publication of his course
	RequestCoursePublication(ctx context.Context, actor domain.Actor, request domain.PublicationRequest) (domain.PublicationRequest, error)
	// UpdatePublicationRequest Moderator reviews publication, if
	UpdatePublicationRequest(ctx context.Context, actor domain.Actor, requestID domain.PublicationRequest) (domain.PublicationRequest, error)
}

// UserService represents the part of interface for user data management.
type UserService interface {
	// BuyCourse student buy course
	BuyCourse(ctx context.Context, actor domain.Actor, courseID uuid.UUID) (domain.PaymentRedirect, error)
	// BuyLesson student buy lesson
	BuyLesson(ctx context.Context, actor domain.Actor, courseID uuid.UUID) (domain.PaymentRedirect, domain.LessonAccess, error)

	// RegisterStudentProfile student registers his profile
	RegisterStudentProfile(ctx context.Context, actor domain.Actor, profile domain.Student) error
	// RegisterTeacherProfile admin or moderator  registers teachers profile
	RegisterTeacherProfile(ctx context.Context, actor domain.Actor, profile domain.Teacher) error

	// GetLessonAccess student gets his access to lesson for a user
	GetLessonAccess(ctx context.Context, actor domain.Actor, lessonID uuid.UUID) (domain.LessonAccess, error)
	// ChangeCourseAccess admin or moderator gives access to lesson for a user
	ChangeCourseAccess(ctx context.Context, actor domain.Actor, access domain.LessonAccess) (domain.LessonAccess, error)
	// ConfirmAccess Finances system from webhook confirms payment
	ConfirmAccess(ctx context.Context, buyerID uuid.UUID, products []domain.Product) error
}

// TODO Delete
//registerLessonFeedback(ctx context.Context, lessonIDs []uuid.UUID) (feedbackIDs []uuid.UUID, err error)
//registerLessonProduct(ctx context.Context, products []*domain.Lesson) (productIDs []uuid.UUID, err error)

//isLessonAccessible(ctx context.Context, userID uuid.UUID, lessonID uuid.UUID) (bool, error)
//IsCourseAccessible(ctx context.Context, userID uuid.UUID, lessonID uuid.UUID) (bool, error)

//registerCourseFeedback(ctx context.Context, courseID uuid.UUID) (feedbackID uuid.UUID, err error)
//registerProducts(ctx context.Context, courseID uuid.UUID) error
