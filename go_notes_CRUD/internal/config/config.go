package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// this file keeps all the config values for the project

type Config struct {
	MongoUri   string
	MongoDB    string
	ServerPort string
}

// load function
// this function validates the env variables for the app

func Load() (Config, error) {

	// godotenv.Load() reads the .env and set them into the process environment
	// os.getenv() reads those values
	if err := godotenv.Load(); err != nil {
		return Config{}, fmt.Errorf("Failed to load env variables")
	}

	mongoUri, err := extractEnv("MONGOURI")
	if err != nil {
		return Config{}, err
	}

	mongo_DB_Name, err := extractEnv("MONGO_DB_NAME")
	if err != nil {
		return Config{}, err
	}

	serverPort, err := extractEnv("PORT")
	if err != nil {
		return Config{}, err
	}

	return Config{
			MongoUri:   mongoUri,
			MongoDB:    mongo_DB_Name,
			ServerPort: serverPort,
		},
		nil

}

// helper function to extract env
func extractEnv(key string) (string, error) {
	val := os.Getenv(key)

	if val == "" {
		return "", fmt.Errorf("Missing required env")
	}

	return val, nil
}
