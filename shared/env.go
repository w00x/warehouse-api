package shared

import (
	"github.com/joho/godotenv"
	"os"
)

func GetEnv(key string) string {
	env := os.Getenv("ENV_APP")
	if env == "" {
		godotenv.Load()
	}
	return os.Getenv(key)
}
