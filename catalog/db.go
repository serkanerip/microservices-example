package catalog

import (
	"context"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type GetMongoDBConnectionParams struct {
	URI string
}

func GetMongoDBConnection(ctx context.Context, params GetMongoDBConnectionParams) (*mongo.Client, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(params.URI))

	if err != nil {
		return nil, errors.Wrap(err, "cannot connect mongodb")
	}

	return client, nil
}
