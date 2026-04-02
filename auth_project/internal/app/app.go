package app

import (
	"auth_project/internal/config"
	"auth_project/internal/db"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type App struct {
	Config config.Config

	MongoClient *mongo.Client
	DB          *mongo.Database
}

func New(ctx context.Context) (*App, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	mongoClient, err := db.Connect(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return &App{
		Config:      cfg,
		MongoClient: mongoClient.Client,
		DB:          mongoClient.DB,
	}, nil

}

func (a *App) Close(ctx context.Context) error {
	if a.MongoClient == nil {
		return nil // nothing to close
	}

	closeCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := db.Disconnect(a.MongoClient); err != nil {
		return fmt.Errorf("failed to disconnect from database: %w", err)
	}

	if err := a.MongoClient.Disconnect(closeCtx); err != nil {
		return fmt.Errorf("failed to disconnect mongo client: %w", err)
	}
	return nil
}
