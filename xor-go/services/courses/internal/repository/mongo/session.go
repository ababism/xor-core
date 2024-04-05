package mongo

import (
	"context"
	"github.com/juju/zaputil/zapctx"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"net/http"
	"xor-go/pkg/apperror"
	"xor-go/services/courses/internal/repository/mongo/collections"
	"xor-go/services/courses/internal/service/adapters"
)

var _ adapters.Session = &Session{}

type Session struct {
	DB *Database

	session *mongo.Session
	//
	//txClient *mongo.Client
	//db       *mongo.Database
	//clientOpt options.ClientOptions
}

func (tx Session) SessionPublications(ctx context.Context) adapters.PublicationRequestRepository {
	return NewPublicationRepository(tx.DB)
}

func (tx Session) SessionLessons(ctx context.Context, name collections.CollectionName) adapters.LessonRepository {
	return NewLessonRepository(tx.DB, name)
}

func (tx Session) SessionCourses(ctx context.Context, name collections.CollectionName) adapters.CourseRepository {
	return NewCourseRepository(tx.DB, name)
}

func (tx Session) IsPossibleTransaction(dbName string, clientOptions options.ClientOptions) bool {
	return clientOptions.GetURI() == tx.DB.clientOptions.GetURI() && dbName == tx.DB.database.Name()
}

func (tx Session) StartTransaction(ctx context.Context) error {
	logger := zapctx.Logger(ctx)
	if tx.session == nil {
		logger.Error("MongoDB nil session error")
		return apperror.New(http.StatusInternalServerError, "internal server error", "MongoDB nil session error", nil)
	}
	s := *tx.session

	if err := s.StartTransaction(); err != nil {
		logger.Error("MongoDB start transaction error")
		return apperror.New(http.StatusInternalServerError, "internal server error", "MongoDB start transaction error", nil)
	}

	return nil
}

func (tx Session) AbortTransaction(ctx context.Context) error {
	logger := zapctx.Logger(ctx)
	if tx.session == nil {
		logger.Error("MongoDB nil session error")
		return apperror.New(http.StatusInternalServerError, "internal server error", "MongoDB nil session error", nil)
	}
	s := *tx.session

	if err := s.AbortTransaction(ctx); err != nil {
		logger.Error("MongoDB abort transaction error")
		return apperror.New(http.StatusInternalServerError, "internal server error", "MongoDB abort transaction error", nil)
	}

	return nil
}

func (tx Session) CommitTransaction(ctx context.Context) error {
	logger := zapctx.Logger(ctx)
	if tx.session == nil {
		logger.Error("MongoDB nil session error")
		return apperror.New(http.StatusInternalServerError, "internal server error", "MongoDB nil session error", nil)
	}
	s := *tx.session

	if err := s.CommitTransaction(ctx); err != nil {
		logger.Error("MongoDB commit transaction error")
		return apperror.New(http.StatusInternalServerError, "internal server error", "MongoDB commit transaction error", nil)
	}

	return nil
}

func (tx Session) EndSession(ctx context.Context) {
	logger := zapctx.Logger(ctx)
	if tx.session == nil {
		logger.Error("MongoDB nil session")
	}
	s := *tx.session
	s.EndSession(ctx)
	return
}

func (cr CourseRepository) NewSession(ctx context.Context) (*Session, error) {
	logger := zapctx.Logger(ctx)

	txClient, err := mongo.Connect(ctx, cr.db.clientOptions)
	if err != nil {
		logger.Error("MongoDB create new client for session error", zap.Error(err))
		return nil, apperror.New(http.StatusInternalServerError, "internal server error", "MongoDB create new client for session error", err)
	}

	txDatabase := txClient.Database(cr.db.database.Name())

	session, err := txClient.StartSession()
	if err != nil {
		logger.Error("MongoDB start session error ", zap.Error(err))
		return nil, apperror.New(http.StatusInternalServerError, "internal server error", "MongoDB start session error", err)

	}
	newTxDB := Database{
		client:        txClient,
		database:      txDatabase,
		clientOptions: cr.db.clientOptions,
	}

	return &Session{
		DB:      &newTxDB,
		session: &session,
	}, nil
}

func (lr LessonRepository) NewSession(ctx context.Context) (*Session, error) {
	logger := zapctx.Logger(ctx)

	txClient, err := mongo.Connect(ctx, lr.db.clientOptions)
	if err != nil {
		logger.Error("MongoDB create new client for session error", zap.Error(err))
		return nil, apperror.New(http.StatusInternalServerError, "internal server error", "MongoDB create new client for session error", err)
	}

	txDatabase := txClient.Database(lr.db.database.Name())

	session, err := txClient.StartSession()
	if err != nil {
		logger.Error("MongoDB start session error ", zap.Error(err))
		return nil, apperror.New(http.StatusInternalServerError, "internal server error", "MongoDB start session error", err)

	}
	newTxDB := Database{
		client:        txClient,
		database:      txDatabase,
		clientOptions: lr.db.clientOptions,
	}

	return &Session{
		DB:      &newTxDB,
		session: &session,
	}, nil
}

func (sr StudentRepository) NewSession(ctx context.Context) (*Session, error) {
	logger := zapctx.Logger(ctx)

	txClient, err := mongo.Connect(ctx, sr.db.clientOptions)
	if err != nil {
		logger.Error("MongoDB create new txClient for session error", zap.Error(err))
		return nil, apperror.New(http.StatusInternalServerError, "internal server error", "MongoDB create new txClient for session error", err)
	}

	txDatabase := txClient.Database(sr.db.database.Name())

	session, err := txClient.StartSession()
	if err != nil {
		logger.Error("MongoDB start session error ", zap.Error(err))
		return nil, apperror.New(http.StatusInternalServerError, "internal server error", "MongoDB start session error", err)

	}
	newTxDB := Database{
		client:        txClient,
		database:      txDatabase,
		clientOptions: sr.db.clientOptions,
	}

	return &Session{
		DB:      &newTxDB,
		session: &session,
	}, nil
}

func (tr TeacherRepository) NewSession(ctx context.Context) (*Session, error) {
	logger := zapctx.Logger(ctx)

	txClient, err := mongo.Connect(ctx, tr.db.clientOptions)
	if err != nil {
		logger.Error("MongoDB create new client for session error", zap.Error(err))
		return nil, apperror.New(http.StatusInternalServerError, "internal server error", "MongoDB create new client for session error", err)
	}

	txDatabase := txClient.Database(tr.db.database.Name())

	session, err := txClient.StartSession()
	if err != nil {
		logger.Error("MongoDB start session error ", zap.Error(err))
		return nil, apperror.New(http.StatusInternalServerError, "internal server error", "MongoDB start session error", err)

	}
	newTxDB := Database{
		client:        txClient,
		database:      txDatabase,
		clientOptions: tr.db.clientOptions,
	}

	return &Session{
		DB:      &newTxDB,
		session: &session,
	}, nil
}

func (pr PublicationRequestRepository) NewSession(ctx context.Context) (*adapters.Session, error) {
	logger := zapctx.Logger(ctx)

	txClient, err := mongo.Connect(ctx, pr.db.clientOptions)
	if err != nil {
		logger.Error("MongoDB create new client for session error", zap.Error(err))
		return nil, apperror.New(http.StatusInternalServerError, "internal server error", "MongoDB create new client for session error", err)
	}

	txDatabase := txClient.Database(pr.db.database.Name())

	session, err := txClient.StartSession()
	if err != nil {
		logger.Error("MongoDB start session error ", zap.Error(err))
		return nil, apperror.New(http.StatusInternalServerError, "internal server error", "MongoDB start session error", err)

	}
	newTxDB := Database{
		client:        txClient,
		database:      txDatabase,
		clientOptions: pr.db.clientOptions,
	}

	s := Session{
		DB:      &newTxDB,
		session: &session,
	}
	as := adapters.Session(s)
	return &as, nil
}
