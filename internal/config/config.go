package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName     string
	AppVersion  string
	GithubToken string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return &Config{
		AppName:     getEnv("APP_NAME", "rag-with-go"),
		AppVersion:  getEnv("APP_VERSION", "1.0.0"),
		GithubToken: getEnv("GITHUB_TOKEN", ""),
	}
}

func getEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
