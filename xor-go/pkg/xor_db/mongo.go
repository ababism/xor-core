package xor_db

import (
	"context"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

type MongoConfig struct {
	Uri      string        `yaml:"uri"`
	Database string        `yaml:"database"`
	Timeout  time.Duration `yaml:"timeout"`
}

func NewMongoClient(ctx context.Context, cfg *MongoConfig) (*mongo.Client, error) {
	opts := options.Client()
	opts.ApplyURI(cfg.Uri)
	opts.SetTimeout(cfg.Timeout)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to create mongo client")
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, errors.WithMessage(err, "failed to ping mongo")
	}

	return client, nil
}
