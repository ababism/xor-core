package service

import (
	"context"
	"fmt"
	_ "github.com/google/uuid"
	"github.com/juju/zaputil/zapctx"
	global "go.opentelemetry.io/otel"
	"net/http"
	"time"
	"xor-go/pkg/apperror"
	"xor-go/services/courses/internal/domain"
)

func (c CoursesService) RequestCoursePublication(initialCtx context.Context, actor domain.Actor, request domain.PublicationRequest) (domain.PublicationRequest, error) {
	_ = zapctx.Logger(initialCtx)

	tr := global.Tracer(domain.ServiceName)
	ctx, span := tr.Start(initialCtx, "courses/service.RequestCoursePublication")
	defer span.End()

	if !actor.HasOneOfRoles(domain.TeacherRole, domain.AdminRole) {
		return domain.PublicationRequest{}, apperror.New(http.StatusForbidden, "user does not have teacher rights to publish course",
			fmt.Sprintf("user do not have %s or %s roles", domain.TeacherRole, domain.AdminRole), nil)
	}
	// TODO Check teacher access
	// teacher opt scenario
	if !actor.HasRole(domain.AdminRole) {
		ok, err := c.teacher.IsCourseAccessible(ctx, actor.ID, request.CourseID)
		if err != nil {
			return domain.PublicationRequest{}, err
		}
		if !ok {
			return domain.PublicationRequest{}, apperror.New(http.StatusForbidden, "teacher can't publish someone else's course",
				fmt.Sprintf("teacher do not own course %s", request.CourseID.String()), nil)
		}
	}

	err := c.publication.Create(ctx, request)
	if err != nil {
		return domain.PublicationRequest{}, err
	}

	request.UpdatedAt = time.Now()
	return request, nil
}

func (c CoursesService) UpdatePublicationRequest(initialCtx context.Context, actor domain.Actor, incomingPR domain.PublicationRequest) (domain.PublicationRequest, error) {
	log := zapctx.Logger(initialCtx)

	tr := global.Tracer(domain.ServiceName)
	ctx, span := tr.Start(initialCtx, "courses/service.UpdatePublicationRequest")
	defer span.End()

	if !actor.HasOneOfRoles(domain.ModeratorRole, domain.AdminRole) {
		return domain.PublicationRequest{}, apperror.New(http.StatusForbidden, "user does not have moderator rights to publish course",
			fmt.Sprintf("user do not have %s or %s roles", domain.ModeratorRole, domain.AdminRole), nil)
	}

	currentPR, err := c.publication.Get(ctx, incomingPR.ID)
	if err != nil {
		return domain.PublicationRequest{}, err
	}
	if currentPR == nil {
		log.Error("got nil publication request from repository with no errors")
		return domain.PublicationRequest{}, apperror.New(http.StatusInternalServerError,
			"nil publication request", "got nil publication request from repository with no errors", nil)
	}

	if currentPR.RequestStatus != domain.Unwatched {
		return *currentPR, apperror.New(http.StatusForbidden, "publication request already checked",
			"publication request already is not unwatched", nil)
	}

	if incomingPR.RequestStatus == domain.Rejected {
		err = c.publication.Update(ctx, incomingPR)
		if err != nil {
			return domain.PublicationRequest{}, err
		}
		// RETURN
		return incomingPR, nil
	}

	switch incomingPR.RequestStatus {
	// ACCEPT
	case domain.Approved:
		// TODO Tx
		err = c.publication.Update(ctx, incomingPR)
		if err != nil {
			return domain.PublicationRequest{}, err
		}

		courseTemplate, err := c.courseEdit.Get(ctx, currentPR.CourseID)
		if err != nil {
			return domain.PublicationRequest{}, err
		}

		err = courseTemplate.Validate()
		if err != nil {
			return domain.PublicationRequest{}, err
		}

		_, err = c.course.Create(ctx, courseTemplate)
		if err != nil {
			return domain.PublicationRequest{}, err
		}
		// RETURN
		return incomingPR, nil

	// REJECT
	case domain.Rejected:
		err = c.publication.Update(ctx, incomingPR)
		if err != nil {
			return domain.PublicationRequest{}, err
		}
		// RETURN
		return incomingPR, nil

	case domain.Unwatched:
		return domain.PublicationRequest{}, apperror.New(http.StatusForbidden,
			"cannot unwatch request", "no option to unwatch publication request", nil)
	default:
		return domain.PublicationRequest{}, apperror.New(http.StatusBadRequest,
			"unsupported request status",
			fmt.Sprintf("unsupported request status can be only %d %d %d",
				domain.Unwatched, domain.Approved, domain.Rejected), nil)
	}
}
