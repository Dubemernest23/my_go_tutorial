package user

import (
	"auth_project/internal/auth"
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo      *Repo
	jwtSecret string
}

func NewService(repo *Repo, jwtSecret string) *Service {
	return &Service{
		repo:      repo,
		jwtSecret: jwtSecret,
	}
}

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type AuthResponse struct {
	Token string     `json:"token"`
	User  PublicUser `json:"user"`
}

func (s *Service) Register(ctx context.Context, input RegisterInput) (AuthResponse, error) {
	email := strings.ToLower(strings.TrimSpace(input.Email))
	username := strings.ToLower(strings.TrimSpace(input.Username))
	password := strings.TrimSpace(input.Password)

	if email == "" || password == "" || username == "" {
		return AuthResponse{}, errors.New("provide all required fields")
	}

	if len(password) < 6 {
		return AuthResponse{}, errors.New("password must be at least 6 characters")
	}

	_, err := s.repo.GetUserByEmail(ctx, email)
	if err == nil {
		return AuthResponse{}, errors.New("email already in use")
	}
	if !errors.Is(err, mongo.ErrNoDocuments) {
		return AuthResponse{}, fmt.Errorf("failed to check existing email: %w", err)
	}

	_, err = s.repo.GetUserByUsername(ctx, username)
	if err == nil {
		return AuthResponse{}, errors.New("username already in use")
	}
	if !errors.Is(err, mongo.ErrNoDocuments) {
		return AuthResponse{}, fmt.Errorf("failed to check existing username: %w", err)
	}

	hashbytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return AuthResponse{}, fmt.Errorf("failed to hash password: %w", err)
	}

	now := time.Now().UTC()
	user := User{
		ID:           bson.NewObjectID(), // ✅ calling the function with ()
		Username:     username,
		Email:        email,
		PasswordHash: string(hashbytes),
		Role:         "user",
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	created, err := s.repo.CreateUser(ctx, user)
	if err != nil {
		return AuthResponse{}, fmt.Errorf("failed to create user: %w", err)
	}

	token, err := auth.CreateToken(s.jwtSecret, created.ID.Hex(), created.Role)
	if err != nil {
		return AuthResponse{}, fmt.Errorf("failed to create token: %w", err)
	}

	return AuthResponse{
		Token: token,
		User:  ToPublicUser(created),
	}, nil
}

func (s *Service) Login(ctx context.Context, input LoginInput) (AuthResponse, error) {

	email := strings.ToLower(strings.TrimSpace(input.Email))
	password := strings.ToLower(input.Password)

	if email == "" || password == "" {
		return AuthResponse{}, errors.New("provide all required field(d)")
	}

	// get user by email
	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			return AuthResponse{}, errors.New("Invalid credentials")
		}
		return AuthResponse{}, err
	}
	fmt.Println(user)

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return AuthResponse{}, errors.New("Invalid credentials")
	}

	token, err := auth.CreateToken(s.jwtSecret, user.ID.Hex(), user.Role)
	if err != nil {
		return AuthResponse{}, fmt.Errorf("failed to create token: %w", err)
	}

	return AuthResponse{
		Token: token,
		User:  ToPublicUser(user),
	}, nil

	// return AuthResponse{}, nil
}
