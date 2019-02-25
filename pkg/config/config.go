package config

import (
	"os"
)

// Config contains the environment specific configuration values needed by the
// application.
type Config struct {
	DatabaseHost     string
	DatabasePort     int
	DatabaseName     string
	DatabaseUser     string
	DatabasePassword string
	Environment      string
	Port             int
}

const environmentENV = "ENVIRONMENT"

// New returns an instance of Config based on the "ENVIRONMENT" environment
// variable.
func New() Config {
	cfg := Config{
		Port:         3000,
		DatabasePort: 5432,
	}

	switch os.Getenv(environmentENV) {
	case "development", "":
		loadDevelopmentConfig(&cfg)
	case "test":
		loadTestConfig(&cfg)
	}

	return cfg
}
