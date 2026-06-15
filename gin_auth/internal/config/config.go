package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)


type MongoConfig struct {
	MONGO_URI string
	MONGO_DB_NAME string
	JWT_SECRET string
}

func Load() (MongoConfig, error) {
	_ = godotenv.Load()

	cfg := MongoConfig{
		MONGO_URI: strings.TrimSpace(os.Getenv("MONGO_UR")),
		MONGO_DB_NAME: strings.TrimSpace(os.Getenv("MONGO_DB_NAME")),
		JWT_SECRET: strings.TrimSpace(os.Getenv("JWT_SECRET")),
	}

	if cfg.MONGO_URI == "" {
		return MongoConfig{}, fmt.Errorf("missing momgo uri")
	}
	if cfg.MONGO_DB_NAME == "" {
		return MongoConfig{}, fmt.Errorf("missing momgo db name")
	}
	if cfg.JWT_SECRET == "" {
		return MongoConfig{}, fmt.Errorf("missing jwt credentials")
	}

	return cfg, nil
}



// extracts env values from key. --initial method one.
func extractEnv(key string) (string, error) {
	value := os.Getenv(key)

	if value == "" {
		return "", fmt.Errorf("could not find value of env")
	}
	return value, nil
}