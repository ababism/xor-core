package service

import (
	"context"
	"github.com/google/uuid"
	"xor-go/services/courses/internal/domain"
)

func (c CoursesService) RequestCoursePublication(ctx context.Context, courseID uuid.UUID) (domain.PublicationRequest, error) {
	//TODO implement me
	panic("implement me")
}

func (c CoursesService) UpdatePublicationRequest(ctx context.Context, requestID domain.PublicationRequest) (domain.PublicationRequest, error) {
	//TODO implement me
	panic("implement me")
}
