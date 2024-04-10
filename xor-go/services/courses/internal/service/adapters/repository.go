package adapters

import (
	"context"
	"github.com/google/uuid"
	"xor-go/services/courses/internal/domain"
)

type CourseRepository interface {
	Create(ctx context.Context, course *domain.Course) (*domain.Course, error)
	Get(ctx context.Context, courseID uuid.UUID) (*domain.Course, error)
	GetAll(ctx context.Context, offset, limit int) ([]*domain.Course, error)
	GetAllByTeacher(ctx context.Context, teacherID uuid.UUID, offset, limit int) ([]*domain.Course, error)
	Update(ctx context.Context, courseID uuid.UUID, course *domain.Course) error
	Delete(ctx context.Context, courseID uuid.UUID) error
}

type LessonRepository interface {
	Create(ctx context.Context, lesson *domain.Lesson) (*domain.Lesson, error)
	Get(ctx context.Context, lessonID uuid.UUID) (*domain.Lesson, error)
	GetAllByCourse(ctx context.Context, courseID uuid.UUID) ([]*domain.Lesson, error)
	Update(ctx context.Context, lesson *domain.Lesson) error
	Delete(ctx context.Context, lessonID uuid.UUID) error
}

type StudentRepository interface {
	Create(ctx context.Context, profile domain.Student) error
	GetLessonAccess(ctx context.Context, userID uuid.UUID, lessonID uuid.UUID) (domain.LessonAccess, error)
	CreateAccessToLesson(ctx context.Context, lessonAccess domain.LessonAccess) (domain.LessonAccess, error)
}

type TeacherRepository interface {
	Create(ctx context.Context, profile domain.Teacher) error
	IsCourseAccessible(ctx context.Context, teacherID uuid.UUID, courseID uuid.UUID) (bool, error)
	//CreateAccessToCourse(ctx context.Context, teacherID uuid.UUID, courseID uuid.UUID) error
}

type PublicationRequestRepository interface {
	Create(ctx context.Context, req domain.PublicationRequest) error
	Get(ctx context.Context, reqID uuid.UUID) (*domain.PublicationRequest, error)
	GetAll(ctx context.Context, offset, limit int) ([]domain.PublicationRequest, error)
	GetAllFromTeacher(ctx context.Context, teacher uuid.UUID, offset, limit int) ([]domain.PublicationRequest, error)
	Update(ctx context.Context, req domain.PublicationRequest) error
	//Delete(ctx context.Context, lessonID uuid.UUID) error
	NewSession(ctx context.Context) (*Session, error)
}
