package db

import (
	"context"
	"fmt"
	config "gin_crud/internal"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func Connect(cfg config.Config) (*mongo.Client, *mongo.Database, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) //context timeout for ping and not connection
	defer cancel()

	clientOptions := options.Client().ApplyURI(cfg.MongoURI)

	connection, err := mongo.Connect(clientOptions) //v2 no longer requires context parameter
	if err != nil {
		return nil, nil, fmt.Errorf("mongo connection failed")
	}

	if err := connection.Ping(ctx, nil); err != nil {
		return nil, nil, fmt.Errorf("Mongo ping failed")
	}

	// setting db aname
	database := connection.Database(cfg.MongoDB)

	return connection, database, nil
}

func disConnect(client *mongo.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return client.Disconnect(ctx)
}