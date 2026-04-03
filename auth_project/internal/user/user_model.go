package user

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson" // ✅ only v2 bson
)

type User struct {
	ID           bson.ObjectID `bson:"_id" json:"id"` // ✅ bson.ObjectID
	Username     string        `bson:"username" json:"username"`
	Email        string        `bson:"email" json:"email"`
	PasswordHash string        `bson:"password_hash" json:"-"`
	Role         string        `bson:"role" json:"role"`
	CreatedAt    time.Time     `bson:"created_at" json:"created_at"`
	UpdatedAt    time.Time     `bson:"updated_at" json:"updated_at"`
}

type PublicUser struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ToPublicUser(u User) PublicUser {
	return PublicUser{
		ID:        u.ID.Hex(),
		Username:  u.Username,
		Email:     u.Email,
		Role:      u.Role,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
