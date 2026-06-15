package app

import (
	"context"
	"fmt"
	"gin_auth/internal/config"
	"gin_auth/internal/db"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type App struct {
	Config config.MongoConfig
	MongoClient *mongo.Client
	Db *mongo.Database
}

func New(ctx context.Context) (*App, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}

	mongoCli, err := db.Connect(ctx, cfg)
	if err != nil {
		return nil, err
	}

	return &App{
		Config: cfg,
		MongoClient: mongoCli.Client,
		Db: mongoCli.Db,
	}, nil
}

func (a *App) Close(ctx context.Context) error {
	if a.MongoClient == nil {
		return nil
	}

	closeCtx, cancel := context.WithTimeout(ctx, 5 * time.Second)
	defer cancel()

	if err := a.MongoClient.Disconnect(closeCtx); err != nil {
		return fmt.Errorf("mongo disconnection failed: %w", err)
	}

	return nil
}