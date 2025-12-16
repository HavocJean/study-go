package mongodb

import (
	"context"
	"os"

	"github.com/HavocJean/study-go/internal/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MONGODB_URL = "MONGODB_URL"
	MONGO_DB    = "MONGO_DB"
)

func NewMongoDBConnection(ctx context.Context) (*mongo.Database, error) {
	mongodb_uri := os.Getenv(MONGODB_URL)
	mongo_database := os.Getenv(MONGO_DB)

	client, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI(mongodb_uri),
	)

	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	logger.Info("Connect successfully")

	return client.Database(mongo_database), nil
}
