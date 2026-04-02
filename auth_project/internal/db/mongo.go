package db

import (
	"auth_project/internal/config"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	// "go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Mongo struct {
	Client *mongo.Client
	DB     *mongo.Database
}

// the connect func should take the context, config and return a Mongo struct and an error
func Connect(ctx context.Context, cfg config.Config) (*Mongo, error) {

	// pass the context to a child context with a timeout for the database connection
	connectCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	// create client options using the config
	clientOpts := options.Client().ApplyURI(cfg.MongoUri)

	client, err := mongo.Connect(clientOpts)
	if err != nil {
		return nil, fmt.Errorf("mongo connection failed: %w", err)
	}

	// use connectCtx not ctx for ping
	if err := client.Ping(connectCtx, nil); err != nil {
		return nil, fmt.Errorf("mongo ping failed: %w", err)
	}

	database := client.Database(cfg.MongoDB)

	return &Mongo{
		Client: client,
		DB:     database,
	}, nil
}

func Disconnect(client *mongo.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Disconnect(ctx); err != nil {
		return fmt.Errorf("failed to disconnect mongo: %w", err)
	}

	return nil
}
