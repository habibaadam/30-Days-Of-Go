package config

// import needed packages
import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
)

// config struct
type Config struct {
	MongoURI string
	MongoDB string
	ServerPort string
}

func Load() (Config, error) {

	// loads env file and sets in to process environment
	if err := godotenv.Load(); err != nil {
		return Config{}, fmt.Errorf("Failed to load .env")
	}

	databaseURI, err := extractEnv("MONGO_URI")
	if err != nil {
		return Config{}, fmt.Errorf("Failed to get db uri")
	}

	dbName, err := extractEnv("MONOD_DB_NAME")
	if err != nil {
		return Config{}, fmt.Errorf("Failed to get db name")
	}

	port, err := extractEnv("PORT")
	if err != nil {
		return Config{}, fmt.Errorf("Failed to get server port")
	}

	return Config{
		MongoURI: databaseURI,
		MongoDB: dbName,
		ServerPort: port,
	}, nil
}

// helper function for extracting env values
func extractEnv(key string) (string, error) {
	value := os.Getenv(key)

	if value == "" {
		return "", fmt.Errorf("missing env value")
	}
	return value, nil
}