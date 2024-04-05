package adapters

import (
	"context"
	"xor-go/services/courses/internal/repository/mongo/collections"
)

type Session interface {
	StartTransaction(ctx context.Context) error
	AbortTransaction(ctx context.Context) error
	CommitTransaction(ctx context.Context) error
	EndSession(ctx context.Context)
	SessionCourses(ctx context.Context, name collections.CollectionName) CourseRepository
	SessionLessons(ctx context.Context, name collections.CollectionName) LessonRepository
}
