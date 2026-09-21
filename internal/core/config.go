package core

import (
	"fmt"
	"os"
)

type Config struct {
	Port  string
	DBURL string
}

func LoadConfig() (*Config, error) {
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		return nil, fmt.Errorf("DB_URL is required")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return &Config{
		Port:  port,
		DBURL: dbURL,
	}, nil
}
