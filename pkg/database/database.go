package database

import (
    "fmt"

    "github.com/ben174/goyagi/pkg/config"
    "github.com/go-pg/pg"
)

// New initializes a new database connection.
func New(cfg config.Config) (*pg.DB, error) {
    addr := fmt.Sprintf("%s:%d", cfg.DatabaseHost, cfg.DatabasePort)

    db := pg.Connect(&pg.Options{
        Addr:     addr,
        User:     cfg.DatabaseUser,
        Password: cfg.DatabasePassword,
        Database: cfg.DatabaseName,
    })

    // Ensure the database can connect
    _, err := db.Exec("SELECT 1")
    if err != nil {
        return nil, err
    }

    return db, nil
}
