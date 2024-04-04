package service

import (
	"context"
	"github.com/google/uuid"
	"xor-go/services/courses/internal/domain"
)

func (c CoursesService) ReadLesson(ctx context.Context, lessonID uuid.UUID) (*domain.Lesson, error) {
	//TODO implement me
	panic("implement me")
}

func (c CoursesService) CreateLesson(ctx context.Context, lesson *domain.Lesson) error {
	//TODO implement me
	panic("implement me")
}

func (c CoursesService) GetLesson(ctx context.Context, lessonID uuid.UUID) (*domain.Lesson, error) {
	//TODO implement me
	panic("implement me")
}

func (c CoursesService) UpdateLesson(ctx context.Context, lesson *domain.Lesson) error {
	//TODO implement me
	panic("implement me")
}

func (c CoursesService) DeleteLesson(ctx context.Context, lessonID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}
