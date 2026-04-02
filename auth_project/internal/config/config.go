package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoUri  string
	MongoDB   string
	JWTSecret string
}

// load env variables and validate them
func Load() (Config, error) {
	// godotenv.Load() reads the .env and set them into the process environment
	// os.getenv() reads those values

	_ = godotenv.Load()
	cfg := Config{
		MongoUri:  strings.TrimSpace(os.Getenv("MONGOURI")),
		MongoDB:   strings.TrimSpace(os.Getenv("MONGO_DB_NAME")),
		JWTSecret: strings.TrimSpace(os.Getenv("JWT_SECRET")),
	}

	if cfg.MongoUri == "" || cfg.MongoDB == "" || cfg.JWTSecret == "" {
		return Config{}, fmt.Errorf("config error %v", os.ErrInvalid)
	}

	return cfg, nil

}
