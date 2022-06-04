package mongodb

import (
	"context"
	"fmt"
	"html"

	"gitlab.internal.b2w.io/team/a-tech/star-wars/src/server/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	singletonStorage *Storage = nil
)

type Storage struct {
	Error   error
	MongoDB *mongo.Database
}

func New(ctx context.Context) Storage {
	cfg := *config.Instance

	if singletonStorage == nil {
		mongoDB, err := connect(ctx, cfg)

		singletonStorage = &Storage{
			Error:   err,
			MongoDB: mongoDB,
		}
	}

	return *singletonStorage
}

func connect(ctx context.Context, cfg config.Config) (*mongo.Database, error) {
	var mongoUri string = html.UnescapeString(fmt.Sprintf("mongodb://%s:%s@%s/%s", cfg.Mongo.User, cfg.Mongo.Pass, cfg.Mongo.Host, cfg.Mongo.Database))

	if cfg.Mongo.Args != "" {
		mongoUri = html.UnescapeString(fmt.Sprintf("mongodb://%s:%s@%s/%s?%s", cfg.Mongo.User, cfg.Mongo.Pass, cfg.Mongo.Host, cfg.Mongo.Database, cfg.Mongo.Args))
	}

	clientOptions := options.Client().ApplyURI(mongoUri)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return client.Database(cfg.Mongo.Database), nil
}
