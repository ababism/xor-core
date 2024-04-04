package service

import (
	"context"
	"github.com/google/uuid"
	"xor-go/services/courses/internal/domain"
)

func (c CoursesService) BuyCourse(ctx context.Context, actor domain.Actor, courseID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (c CoursesService) BuyLesson(ctx context.Context, actor domain.Actor, courseID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (c CoursesService) RegisterStudentProfile(ctx context.Context, actor domain.Actor, profile domain.Student) error {
	//TODO implement me
	panic("implement me")
}

func (c CoursesService) RegisterTeacherProfile(ctx context.Context, actor domain.Actor, teacherID uuid.UUID, profile domain.Teacher) error {
	//TODO implement me
	panic("implement me")
}

func (c CoursesService) ChangeCourseAccess(ctx context.Context, actor domain.Actor, userID uuid.UUID, lessonID uuid.UUID) (domain.LessonAccess, error) {
	//TODO implement me
	panic("implement me")
}

func (c CoursesService) GetCourseAccess(ctx context.Context, actor domain.Actor, lessonID uuid.UUID) (domain.LessonAccess, error) {
	//TODO implement me
	panic("implement me")
}
