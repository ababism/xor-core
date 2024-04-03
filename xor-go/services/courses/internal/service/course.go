package service

import (
	"context"
	"github.com/google/uuid"
	"xor-go/services/courses/internal/domain"
)

func (s *CoursesService) CreateCourse(ctx context.Context, teacherID uuid.UUID, course *domain.Course) (uuid.UUID, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CoursesService) GetCourse(ctx context.Context, courseID uuid.UUID) (*domain.Course, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CoursesService) UpdateCourse(ctx context.Context, courseID uuid.UUID, course *domain.Course) error {
	//TODO implement me
	panic("implement me")
}

func (s *CoursesService) DeleteCourse(ctx context.Context, courseID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (s *CoursesService) publishCourse(ctx context.Context, courseID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (s *CoursesService) registerCourseFeedback(ctx context.Context, courseID uuid.UUID) (feedbackID uuid.UUID, err error) {
	//TODO implement me
	panic("implement me")
}

func (s *CoursesService) registerProducts(ctx context.Context, courseID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (s *CoursesService) ConfirmAccess(ctx context.Context, buyerID uuid.UUID, productIDs []uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}
