package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
	Port        string
	Environment string
}

func Load() *Config {
	if error := godotenv.Load(); error != nil {
		log.Println("no .env file found, cp .env.dist first.")
	}

	return &Config{
		DatabaseURL:   getEnv("DATABASE_URL", ""),
		Port:          getEnv("PORT", "3000"),
		Environment:   getEnv("GO_ENV", "development"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}