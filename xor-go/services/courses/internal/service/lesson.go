package service

import (
	"context"
	"github.com/google/uuid"
	"xor-go/services/courses/internal/domain"
)

func (s *CoursesService) CreateLesson(ctx context.Context, lesson *domain.Lesson) error {
	//TODO implement me
	panic("implement me")
}

func (s *CoursesService) GetLesson(ctx context.Context, lessonID uuid.UUID) (*domain.Lesson, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CoursesService) UpdateLesson(ctx context.Context, lesson *domain.Lesson) error {
	//TODO implement me
	panic("implement me")
}

func (s *CoursesService) DeleteLesson(ctx context.Context, lessonID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (s *CoursesService) registerLessonFeedback(ctx context.Context, lessonIDs []uuid.UUID) (feedbackIDs []uuid.UUID, err error) {
	//TODO implement me
	panic("implement me")
}

func (s *CoursesService) registerLessonProduct(ctx context.Context, products []*domain.Lesson) (productIDs []uuid.UUID, err error) {
	//TODO implement me
	panic("implement me")
}
