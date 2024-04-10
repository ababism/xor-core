package service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	_ "github.com/google/uuid"
	"github.com/juju/zaputil/zapctx"
	global "go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"net/http"
	"time"
	"xor-go/pkg/xapperror"
	"xor-go/services/courses/internal/domain"
	"xor-go/services/courses/internal/domain/keys"
	"xor-go/services/courses/internal/repository/mongo/collections"
)

func (c CoursesService) RequestCoursePublication(initialCtx context.Context, actor domain.Actor, request domain.PublicationRequest) (domain.PublicationRequest, error) {
	_ = zapctx.Logger(initialCtx)

	tr := global.Tracer(domain.ServiceName)
	ctx, span := tr.Start(initialCtx, "courses/service.RequestCoursePublication")
	defer span.End()

	if !actor.HasOneOfRoles(domain.TeacherRole, domain.AdminRole) {
		return domain.PublicationRequest{}, xapperror.New(http.StatusForbidden, "user does not have teacher rights to publish course",
			fmt.Sprintf("user do not have %s or %s roles", domain.TeacherRole, domain.AdminRole), nil)
	}
	// teacher opt scenario
	if !actor.HasRole(domain.AdminRole) {
		ok, err := c.teacher.IsCourseAccessible(ctx, actor.ID, request.CourseID)
		if err != nil {
			return domain.PublicationRequest{}, err
		}
		if !ok {
			return domain.PublicationRequest{}, xapperror.New(http.StatusForbidden, "teacher can't publish someone else's course",
				fmt.Sprintf("teacher do not own course %s", request.CourseID.String()), nil)
		}
	}
	request.UpdatedAt = time.Now()
	request.RequestStatus = domain.Unwatched
	request.AssigneeID = actor.ID

	if request.ID == uuid.Nil || (request.ID == uuid.UUID{}) {
		request.ID = uuid.New()
	}

	err := c.publication.Create(ctx, request)
	if err != nil {
		return domain.PublicationRequest{}, err
	}

	span.AddEvent("publication request created", trace.WithAttributes(attribute.String(keys.LessonIDAttributeKey, request.ID.String())))
	return request, nil
}

func (c CoursesService) UpdatePublicationRequest(initialCtx context.Context, actor domain.Actor, incomingPR domain.PublicationRequest) (domain.PublicationRequest, error) {
	log := zapctx.Logger(initialCtx)

	tr := global.Tracer(domain.ServiceName)
	ctx, span := tr.Start(initialCtx, "courses/service.UpdatePublicationRequest")
	defer span.End()

	if !actor.HasOneOfRoles(domain.ModeratorRole, domain.AdminRole) {
		return domain.PublicationRequest{}, xapperror.New(http.StatusForbidden, "user does not have moderator rights to publish course",
			fmt.Sprintf("user do not have %s or %s roles", domain.ModeratorRole, domain.AdminRole), nil)
	}

	currentPR, err := c.publication.Get(ctx, incomingPR.ID)
	if err != nil {
		return domain.PublicationRequest{}, err
	}
	if currentPR == nil {
		log.Error("nil publication request from repository with no errors")
		return domain.PublicationRequest{}, xapperror.New(http.StatusInternalServerError,
			"nil publication request", "nil publication request from repository with no errors", nil)
	}

	if currentPR.RequestStatus != domain.Unwatched {
		return *currentPR, xapperror.New(http.StatusForbidden, "publication request already checked",
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
		txSessionP, err := c.publication.NewSession(ctx)
		if err != nil {
			return domain.PublicationRequest{}, err
		}
		if txSessionP == nil {
			log.Error("nil publication session from repository with no errors")
			return domain.PublicationRequest{}, xapperror.New(http.StatusInternalServerError,
				"internal transaction error", "nil publication session from repository with no error", nil)
		}
		txSession := *txSessionP

		defer txSession.EndSession(ctx)

		publicationTxSession := txSession.SessionPublications(ctx)

		err = publicationTxSession.Update(ctx, incomingPR)
		if err != nil {
			errTx := txSession.AbortTransaction(ctx)
			if errTx != nil {
				log.Error("err aborting publication session transaction")
				return domain.PublicationRequest{}, xapperror.New(http.StatusInternalServerError,
					"internal transaction error", "err aborting publication session transaction", errTx)
			}
			return domain.PublicationRequest{}, err
		}
		// COURSE PREPARE
		courseEditTxSession := txSession.SessionCourses(ctx, collections.CourseEditorCollectionName)

		courseTemplate, err := courseEditTxSession.Get(ctx, currentPR.CourseID)
		if err != nil {
			errTx := txSession.AbortTransaction(ctx)
			if errTx != nil {
				log.Error("err aborting publication session transaction")
				return domain.PublicationRequest{}, xapperror.New(http.StatusInternalServerError,
					"internal transaction error", "err aborting publication session transaction", errTx)
			}
			return domain.PublicationRequest{}, err
		}

		err = courseTemplate.Validate()
		if err != nil {
			errTx := txSession.AbortTransaction(ctx)
			if errTx != nil {
				log.Error("err aborting publication session transaction")
				return domain.PublicationRequest{}, xapperror.New(http.StatusInternalServerError,
					"internal transaction error", "err aborting publication session transaction", errTx)
			}
			return domain.PublicationRequest{}, err
		}

		if courseTemplate.TeacherID != currentPR.AssigneeID {
			log.Error("err teacher made publication request on someone else's course")
			return domain.PublicationRequest{}, xapperror.New(http.StatusInternalServerError,
				"course owner and request assignee do not match", "course owner and request assignee do not match", nil)
		}

		courseTxSession := txSession.SessionCourses(ctx, collections.CourseCollectionName)

		_, err = courseTxSession.Create(ctx, courseTemplate)
		if err != nil {
			errTx := txSession.AbortTransaction(ctx)
			if errTx != nil {
				log.Error("err aborting publication session transaction")
				return domain.PublicationRequest{}, xapperror.New(http.StatusInternalServerError,
					"internal transaction error", "err aborting publication session transaction", errTx)
			}
			return domain.PublicationRequest{}, err
		}
		// LESSON PREPARE
		lessonEditTxSession := txSession.SessionLessons(ctx, collections.LessonEditorCollectionName)

		lessonTemplates, err := lessonEditTxSession.GetAllByCourse(ctx, currentPR.CourseID)
		if err != nil {
			errTx := txSession.AbortTransaction(ctx)
			if errTx != nil {
				log.Error("err aborting publication session transaction")
				return domain.PublicationRequest{}, xapperror.New(http.StatusInternalServerError,
					"internal transaction error", "err aborting publication session transaction", errTx)
			}
			return domain.PublicationRequest{}, err
		}

		lessonTxSession := txSession.SessionLessons(ctx, collections.LessonCollectionName)

		products := make([]domain.Product, 0, 30)

		for _, l := range lessonTemplates {
			errLV := l.Validate()
			// TODO check again
			if errLV != nil || l.TeacherID != courseTemplate.TeacherID {
				errTx := txSession.AbortTransaction(ctx)
				if errTx != nil {
					log.Error("err aborting publication session transaction")
					return domain.PublicationRequest{}, xapperror.New(http.StatusInternalServerError,
						"internal transaction error", "err aborting publication session transaction", errTx)
				}
				return domain.PublicationRequest{}, errLV
			}

			errPV := l.Product.Validate()
			if errPV != nil {
				errTx := txSession.AbortTransaction(ctx)
				if errTx != nil {
					log.Error("err aborting publication session transaction")
					return domain.PublicationRequest{}, xapperror.New(http.StatusInternalServerError,
						"internal transaction error", "err aborting publication session transaction", errTx)
				}
				return domain.PublicationRequest{}, errPV
			}

			products = append(products, l.Product)

			_, errLC := lessonTxSession.Create(ctx, l)
			if errLC != nil {
				errTx := txSession.AbortTransaction(ctx)
				if errTx != nil {
					log.Error("err aborting publication session transaction")
					return domain.PublicationRequest{}, xapperror.New(http.StatusInternalServerError,
						"internal transaction error", "err aborting publication session transaction", errTx)
				}
				return domain.PublicationRequest{}, errLC
			}
		}

		// REGISTER PRODUCTS
		_, errRegister := c.financesClient.RegisterProducts(ctx, products)
		if errRegister != nil {
			errTx := txSession.AbortTransaction(ctx)
			if errTx != nil {
				log.Error("err aborting publication session transaction")
				return domain.PublicationRequest{}, xapperror.New(http.StatusInternalServerError,
					"internal transaction error", "err aborting publication session transaction", errTx)
			}
			return domain.PublicationRequest{}, errRegister
		}

		// NO ERRORS
		errTx := txSession.CommitTransaction(ctx)
		if errTx != nil {
			log.Debug("err commit publication session transaction")
			return domain.PublicationRequest{}, xapperror.New(http.StatusInternalServerError,
				"internal transaction error", "err commit publication session transaction", errTx)
		}

		// RETURN CASE
		return incomingPR, nil

	// REJECT
	case domain.Rejected:
		err = c.publication.Update(ctx, incomingPR)
		if err != nil {
			return domain.PublicationRequest{}, err
		}
		// RETURN CASE
		return incomingPR, nil

	case domain.Unwatched:
		return domain.PublicationRequest{}, xapperror.New(http.StatusForbidden,
			"cannot unwatch request", "no option to unwatch publication request", nil)
	default:
		return domain.PublicationRequest{}, xapperror.New(http.StatusBadRequest,
			"unsupported request status",
			fmt.Sprintf("unsupported request status can be only %d %d %d",
				domain.Unwatched, domain.Approved, domain.Rejected), nil)
	}
}

func (c CoursesService) UpdatePublicationRequestWithoutTx(initialCtx context.Context, actor domain.Actor, incomingPR domain.PublicationRequest) (domain.PublicationRequest, error) {
	log := zapctx.Logger(initialCtx)

	tr := global.Tracer(domain.ServiceName)
	ctx, span := tr.Start(initialCtx, "courses/service.UpdatePublicationRequest")
	defer span.End()

	if !actor.HasOneOfRoles(domain.ModeratorRole, domain.AdminRole) {
		return domain.PublicationRequest{}, xapperror.New(http.StatusForbidden, "user does not have moderator rights to publish course",
			fmt.Sprintf("user do not have %s or %s roles", domain.ModeratorRole, domain.AdminRole), nil)
	}

	currentPR, err := c.publication.Get(ctx, incomingPR.ID)
	if err != nil {
		return domain.PublicationRequest{}, err
	}
	if currentPR == nil {
		log.Error("got nil publication request from repository with no errors")
		return domain.PublicationRequest{}, xapperror.New(http.StatusInternalServerError,
			"nil publication request", "got nil publication request from repository with no errors", nil)
	}

	if currentPR.RequestStatus != domain.Unwatched {
		return *currentPR, xapperror.New(http.StatusForbidden, "publication request already checked",
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
		return domain.PublicationRequest{}, xapperror.New(http.StatusForbidden,
			"cannot unwatch request", "no option to unwatch publication request", nil)
	default:
		return domain.PublicationRequest{}, xapperror.New(http.StatusBadRequest,
			"unsupported request status",
			fmt.Sprintf("unsupported request status can be only %d %d %d",
				domain.Unwatched, domain.Approved, domain.Rejected), nil)
	}
}
