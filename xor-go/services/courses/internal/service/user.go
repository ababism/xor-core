package service

import (
	"context"
	"github.com/google/uuid"
	"xor-go/services/courses/internal/domain"
)

func (s *CoursesService) ReadLesson(ctx context.Context, lessonID uuid.UUID) (*domain.Lesson, error) {
	// If teacher show all, else censor
	//TODO implement me
	panic("implement me")
}

func (s *CoursesService) BuyCourse(ctx context.Context, studentID uuid.UUID, courseID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (s *CoursesService) RegisterStudentProfile(ctx context.Context, accountID uuid.UUID, profile domain.Student) error {
	//TODO implement me
	panic("implement me")
}

func (s *CoursesService) RegisterTeacherProfile(ctx context.Context, accountID uuid.UUID, profile domain.Teacher) error {
	//TODO implement me
	panic("implement me")
}

func (s *CoursesService) ChangeCourseAccess(ctx context.Context, teacherID uuid.UUID, courseID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}
