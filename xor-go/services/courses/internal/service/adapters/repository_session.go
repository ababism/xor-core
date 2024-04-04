package adapters

import (
	"context"
)

type Session interface {
	StartTransaction(ctx context.Context) error
	AbortTransaction(ctx context.Context) error
	CommitTransaction(ctx context.Context) error
	EndSession(ctx context.Context)
	SessionCourses(ctx context.Context) CourseRepository
	SessionLessons(ctx context.Context) LessonRepository
}
