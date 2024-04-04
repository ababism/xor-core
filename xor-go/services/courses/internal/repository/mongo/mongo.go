package mongo

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
	"net/http"
	"xor-go/pkg/apperror"
	"xor-go/services/courses/internal/repository/mongo/migrate"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

//var _ adapters.CourseRepository = &Database{}

type Database struct {
	client   *mongo.Client
	database *mongo.Database

	//databaseName string

	clientOptions *options.ClientOptions
	logger        *zap.Logger
}

func NewDatabase(logger *zap.Logger) *Database {
	return &Database{logger: logger}
}
func createIDFilter(ID uuid.UUID) bson.M {
	filter := bson.M{"_id": ID}
	return filter
}

func (r *Database) Connect(ctx context.Context, cfg *Config, migrateCfg *ConfigMigrations) (func(ctx context.Context) error, error) {
	// TODO take from context
	// TODO AppError
	r.logger.Info("Connecting to mongo...")
	r.logger.Info(fmt.Sprintf("mongo params: uri: %s; database: %s", cfg.Uri, cfg.Database))
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.Uri))
	if err != nil {
		r.logger.Error("new mongo client create error:", zap.Error(err))
		return nil, fmt.Errorf("new mongo client create error: %w", err)
	}

	r.clientOptions = options.Client().ApplyURI(cfg.Uri)

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		r.logger.Error("new mongo primary node connect error:", zap.Error(err))
		return client.Disconnect, fmt.Errorf("new mongo primary node connect error: %w", err)
	}

	r.client = client
	database := client.Database(cfg.Database)

	//r.databaseName = cfg.Database

	if migrateCfg.Enabled {
		migrationSvc := migrate.NewMigrationsService(r.logger, database)
		err = migrationSvc.RunMigrations(migrateCfg.Path)
		if err != nil {
			r.logger.Fatal("run migrations failed", zap.Error(err))
			return client.Disconnect, fmt.Errorf("run migrations failed")
		}
	}

	return client.Disconnect, nil
}

func handleMongoError(err error, log *zap.Logger) error {
	if err == nil {
		return nil
	}
	switch tErr := err.(type) {
	case mongo.CommandError:
		log.Error("MongoDB command error", zap.Error(tErr))
		return apperror.New(http.StatusInternalServerError, "internal server error", "MongoDB command error", tErr)

	case mongo.WriteException:
		log.Error("MongoDB write exception", zap.Error(tErr))
		return apperror.New(http.StatusInternalServerError, "internal server error", "MongoDB write exception", tErr)

	case mongo.ServerError:
		log.Error("MongoDB server error", zap.Error(tErr))
		return apperror.New(http.StatusInternalServerError, "internal server error", "MongoDB server error", tErr)

	case mongo.BulkWriteError:
		log.Error("MongoDB bulk write error", zap.Error(tErr))
		return apperror.New(http.StatusInternalServerError, "internal server error", "MongoDB bulk write error", tErr)

	case mongo.WriteConcernError:
		log.Error("MongoDB write concern error", zap.Error(tErr))
		return apperror.New(http.StatusInternalServerError, "internal server error", "MongoDB write concern error", tErr)

	case mongo.WriteError:
		log.Error("MongoDB write error", zap.Error(tErr))
		return apperror.New(http.StatusInternalServerError, "internal server error", "MongoDB write error", tErr)

	case mongo.MarshalError:
		log.Error("MongoDB write marshal error", zap.Error(tErr))
		return apperror.New(http.StatusInternalServerError, "internal server error", "MongoDB marshal error", tErr)
	}
	return nil
}
