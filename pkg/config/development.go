package config

import (
	"os"
	"strconv"
)

func loadDevelopmentConfig(cfg *Config) {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err == nil {
		cfg.Port = port
	}

	cfg.DatabaseHost = "127.0.0.1"
	cfg.DatabaseName = "goyagi"
	cfg.DatabaseUser = "goyagi_admin"
	cfg.Environment = "development"
}
