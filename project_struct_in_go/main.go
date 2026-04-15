package main

import (
	"context"
	"fmt"
)

// 1. Define the dependency as an INTERFACE (this is the key!)
type Database interface {
	GetUser(ctx context.Context, id int) (string, error)
}

// 2. Real implementation (PostgreSQL, for example)
type PostgresDB struct{}

// fakeDB implementation
type FakeDB struct {
	Users map[int]string
}

func NewFakeDB() *FakeDB {
	return &FakeDB{
		Users: map[int]string{
			42: "Alicea string",
			99: "emmy isa string",
		},
	}

}

func (db *FakeDB) GetUser(ctx context.Context, id int) (string, error) {
	// check cancellation
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	default:
	}
	// return data from fake map
	// return empty string if user is not found

	if name, exists := db.Users[id]; exists {
		return name, nil
	}
	return "", fmt.Errorf("User doesn't exist")
}

func (db PostgresDB) GetUser(ctx context.Context, id int) (string, error) {
	// In real life this would query the DB
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	default:
	}
	return fmt.Sprintf("User %d from PostgreSQL", id), nil
}

// 3. Our service that DEPENDS on the Database
type UserService struct {
	db Database // ← This is the dependency we will inject
}

// 4. Constructor (this is where we do Dependency Injection)
func NewUserService(db Database) *UserService {
	return &UserService{db: db}
}

// 5. Method that uses everything we learned
func (s *UserService) GetUserName(ctx context.Context, id int) (string, error) {
	name, err := s.db.GetUser(ctx, id)
	if err != nil {
		return "", fmt.Errorf("failed to get user: %w", err) // wrap error
	}
	return name, nil
}

func main() {
	// Create the real database
	realDB := PostgresDB{}

	// Inject it into the service ← THIS IS DEPENDENCY INJECTION!
	service := NewUserService(realDB)

	// Create a context (as we learned earlier)
	ctx := context.Background()

	name, err := service.GetUserName(ctx, 42)
	// _, err := service.GetUserName(ctx, 42)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Found:", name)
	fmt.Println("\n=== Using Fake Database ===")
	fakeDB := NewFakeDB()                 // Create fake
	fakeService := NewUserService(fakeDB) // ← Inject fake instead!

	name2, err2 := fakeService.GetUserName(ctx, 42)
	if err2 != nil {
		fmt.Println("Error:", err2)
	} else {
		fmt.Println("Found:", name2)
	}
}
