package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DSN string
	MINENTROPY int
	JWTsecret string
}

func LoadConfig() *Config {
	err := godotenv.Load("env/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DATABASE_URL")
	jwtSecret := os.Getenv("JWT_SECRET")
	minEntropy, err := strconv.Atoi(os.Getenv("MINENTROPY"))
	if err != nil {
		log.Fatalf("Invalid MINENTROPY: %v", err)
	}

	return &Config{
		DSN:       dsn,
		MINENTROPY: minEntropy,
		JWTsecret: jwtSecret,
	}
} 