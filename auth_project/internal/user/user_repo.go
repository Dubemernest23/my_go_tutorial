package user

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson" // ✅ only v2, removed v1
	"go.mongodb.org/mongo-driver/v2/mongo"
	// "go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Repo struct {
	coll *mongo.Collection
}

func NewRepo(db *mongo.Database) *Repo {
	return &Repo{coll: db.Collection("users")}
}

func (r *Repo) CreateUser(ctx context.Context, u User) (User, error) {
	opctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	res, err := r.coll.InsertOne(opctx, u)
	if err != nil {
		return User{}, fmt.Errorf("failed to create user: %w", err)
	}

	// ✅ correct type assertion for v2
	id, ok := res.InsertedID.(bson.ObjectID)
	if !ok {
		return User{}, fmt.Errorf("failed to extract inserted ID")
	}
	u.ID = id

	return u, nil
}

func (r *Repo) GetUserByEmail(ctx context.Context, email string) (User, error) {
	opctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	email = strings.ToLower(strings.TrimSpace(email))
	filter := bson.M{"email": email}
	var user User

	err := r.coll.FindOne(opctx, filter).Decode(&user) // ✅ removed unnecessary options.FindOne()
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return User{}, mongo.ErrNoDocuments
		}
		return User{}, fmt.Errorf("failed to get user by email: %w", err)
	}

	return user, nil
}

func (r *Repo) GetUserByUsername(ctx context.Context, username string) (User, error) {
	opctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	username = strings.ToLower(strings.TrimSpace(username))
	filter := bson.M{"username": username}
	var user User

	err := r.coll.FindOne(opctx, filter).Decode(&user) // ✅ removed unnecessary options.FindOne()
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return User{}, mongo.ErrNoDocuments
		}
		return User{}, fmt.Errorf("failed to get user by username: %w", err)
	}

	return user, nil
}
