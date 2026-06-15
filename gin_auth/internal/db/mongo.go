package db

import (
	"context"
	"fmt"
	"gin_auth/internal/config"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Mongo struct {
	Client *mongo.Client
	Db *mongo.Database
}

func Connect(ctx context.Context, cfg config.MongoConfig) (*Mongo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5* time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(cfg.MONGO_URI)

	connection, err := mongo.Connect(clientOptions)

	if err != nil {
		return nil, fmt.Errorf("mongo connection failed: %w", err)
	}

	if err := connection.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("Ping faileed for database : %w", err)
	}

	database := connection.Database(cfg.MONGO_DB_NAME)

	return &Mongo{
		Client: connection,
		Db: database,
	}, nil
}