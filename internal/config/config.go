package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/takak2166/scrapbox-mcp/internal/errors"
)

type Config struct {
	ScrapboxSID string
	ProjectName string
	Port        int
}

func LoadConfig() (*Config, error) {
	// Load .env file if exists
	err := godotenv.Load()
	if err != nil {
		log.Printf("failed to load .env file: %v", err)
	}

	sid := os.Getenv("SCRAPBOX_SID")
	if sid == "" {
		return nil, &errors.ScrapboxError{Code: errors.ErrInvalidCredentials, Message: "SCRAPBOX_SID is not set", Err: nil}
	}

	project := os.Getenv("SCRAPBOX_PROJECT")
	if project == "" {
		return nil, &errors.ScrapboxError{Code: errors.ErrInvalidCredentials, Message: "SCRAPBOX_PROJECT is not set", Err: nil}
	}

	// Default port is 8080
	port := 8080
	if portStr := os.Getenv("PORT"); portStr != "" {
		var err error
		port, err = strconv.Atoi(portStr)
		if err != nil {
			return nil, &errors.ScrapboxError{Code: errors.ErrInvalidCredentials, Message: "PORT must be a valid number", Err: err}
		}
	}

	return &Config{
		ScrapboxSID: sid,
		ProjectName: project,
		Port:        port,
	}, nil
}
