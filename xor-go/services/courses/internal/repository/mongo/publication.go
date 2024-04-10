package mongo

import (
	"context"
	"github.com/google/uuid"
	"github.com/juju/zaputil/zapctx"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	global "go.opentelemetry.io/otel"
	"net/http"
	"xor-go/pkg/xapperror"
	"xor-go/services/courses/internal/domain"
	"xor-go/services/courses/internal/repository/mongo/collections"
	"xor-go/services/courses/internal/repository/mongo/models"
	"xor-go/services/courses/internal/service/adapters"
)

var _ adapters.PublicationRequestRepository = &PublicationRequestRepository{}

func NewPublicationRepository(database *Database) *PublicationRequestRepository {
	publicationCollection := database.database.Collection(collections.PublicationRequestsCollectionName.String())

	return &PublicationRequestRepository{
		db:             database,
		publications:   publicationCollection,
		collectionName: collections.PublicationRequestsCollectionName,
	}
}

type PublicationRequestRepository struct {
	db *Database

	collectionName collections.CollectionName
	publications   *mongo.Collection
}

func (r PublicationRequestRepository) Create(ctx context.Context, req domain.PublicationRequest) error {
	logger := zapctx.Logger(ctx)

	tr := global.Tracer(domain.ServiceName)
	newCtx, span := tr.Start(ctx, "courses/repository/mongo/publication.Create")
	defer span.End()

	publicationRequest := models.ToMongoModelPublicationRequest(req)
	_, err := r.publications.InsertOne(newCtx, publicationRequest)
	if mErr := handleMongoError(err, logger); mErr != nil {
		return mErr
	}
	if err != nil {
		appErr := xapperror.New(http.StatusBadRequest, "publication request already created",
			"failed to create publication request in MongoDB", err)
		return appErr
	}

	return nil
}

func (r PublicationRequestRepository) Get(ctx context.Context, reqID uuid.UUID) (*domain.PublicationRequest, error) {
	logger := zapctx.Logger(ctx)

	tr := global.Tracer(domain.ServiceName)
	newCtx, span := tr.Start(ctx, "courses/repository/mongo/publication.Get")
	defer span.End()

	var publicationRequest models.PublicationRequest
	err := r.publications.FindOne(newCtx, models.PublicationRequest{ID: reqID.String()}).Decode(&publicationRequest)
	if err != nil {
		if mErr := handleMongoError(err, logger); mErr != nil {
			return nil, mErr
		}
		if err == mongo.ErrNoDocuments {
			return nil, xapperror.New(http.StatusNotFound, "publication request not found",
				"publication request not found", err)
		}
		return nil, xapperror.New(http.StatusInternalServerError, "internal server error",
			"failed to get publication request from MongoDB", err)
	}
	res, err := publicationRequest.ToDomain()
	if err != nil {
		appErr := xapperror.New(http.StatusNotFound, "can't parse publication request", "error converting publication request from MongoDB to domain", err)
		return nil, appErr
	}
	return res, nil
}

func (r PublicationRequestRepository) GetAll(ctx context.Context, offset, limit int) ([]domain.PublicationRequest, error) {
	logger := zapctx.Logger(ctx)

	tr := global.Tracer(domain.ServiceName)
	newCtx, span := tr.Start(ctx, "courses/repository/mongo/publication.GetAll")
	defer span.End()

	cursor, err := r.publications.Find(newCtx, nil)
	if mErr := handleMongoError(err, logger); mErr != nil {
		return nil, mErr
	}
	defer cursor.Close(newCtx)

	var publicationRequests []models.PublicationRequest
	if err = cursor.All(newCtx, &publicationRequests); err != nil {
		if mErr := handleMongoError(err, logger); mErr != nil {
			return nil, mErr
		}
		return nil, xapperror.New(http.StatusInternalServerError, "internal server error",
			"failed to get publication requests from MongoDB", err)
	}

	var res []domain.PublicationRequest
	for _, publicationRequest := range publicationRequests {
		dR, err := publicationRequest.ToDomain()
		if err != nil {
			appErr := xapperror.New(http.StatusInternalServerError, "can't parse publication request",
				"error converting publication request from MongoDB to domain", err)
			return nil, appErr
		}
		res = append(res, *dR)
	}

	return res, nil
}

func (r PublicationRequestRepository) GetAllFromTeacher(ctx context.Context, teacherID uuid.UUID, offset, limit int) ([]domain.PublicationRequest, error) {
	logger := zapctx.Logger(ctx)

	tr := global.Tracer(domain.ServiceName)
	newCtx, span := tr.Start(ctx, "courses/repository/mongo/publication.GetAllFromTeacher")
	defer span.End()

	filter := createUUIDFilter(teacherID, "teacher_id")
	opts := options.Find().SetSkip(int64(offset)).SetLimit(int64(limit))
	cursor, err := r.publications.Find(newCtx, filter, opts)
	if mErr := handleMongoError(err, logger); mErr != nil {
		return nil, mErr
	}
	defer cursor.Close(newCtx)

	var publicationRequests []models.PublicationRequest
	if err = cursor.All(newCtx, &publicationRequests); err != nil {
		if mErr := handleMongoError(err, logger); mErr != nil {
			return nil, mErr
		}
		return nil, xapperror.New(http.StatusInternalServerError, "internal server error",
			"failed to get publication requests from MongoDB", err)
	}

	var res []domain.PublicationRequest
	for _, publicationRequest := range publicationRequests {
		dR, err := publicationRequest.ToDomain()
		if err != nil {
			appErr := xapperror.New(http.StatusInternalServerError, "can't parse publication request",
				"error converting publication request from MongoDB to domain", err)
			return nil, appErr
		}
		res = append(res, *dR)
	}

	return res, nil
}

func (r PublicationRequestRepository) Update(ctx context.Context, req domain.PublicationRequest) error {
	logger := zapctx.Logger(ctx)

	tr := global.Tracer(domain.ServiceName)
	newCtx, span := tr.Start(ctx, "courses/repository/mongo/publication.Update")
	defer span.End()

	publicationRequest := models.ToMongoModelPublicationRequest(req)
	_, err := r.publications.ReplaceOne(newCtx, models.PublicationRequest{ID: req.ID.String()}, publicationRequest)
	if mErr := handleMongoError(err, logger); mErr != nil {
		return mErr
	}
	if err != nil {
		appErr := xapperror.New(http.StatusInternalServerError, "failed to update publication request",
			"failed to update publication request in MongoDB", err)
		return appErr
	}

	return nil
}
