package main

import (
	"os"

	"github.com/ben174/goyagi/pkg/config"
	"github.com/ben174/goyagi/pkg/database"
	"github.com/lob/logger-go"
	migrations "github.com/robinjoseph08/go-pg-migrations"
)

const directory = "cmd/migrations"

func main() {
	log := logger.New()

	cfg := config.New()

	db, err := database.New(cfg)
	if err != nil {
		log.Err(err).Fatal("failed to connect to database")
	}

	err = migrations.Run(db, directory, os.Args)
	if err != nil {
		log.Err(err).Fatal("failed migration")
	}
}
