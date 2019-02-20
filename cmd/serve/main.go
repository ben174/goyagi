package main

import (
    "net/http"

    "github.com/ben174/goyagi/pkg/server"
    "github.com/lob/logger-go"
)

func main() {
    log := logger.New()

    srv := server.New()

    log.Info("server started")

    err := srv.ListenAndServe()
    if err != nil && err != http.ErrServerClosed {
        log.Err(err).Fatal("server stopped")
    }

    log.Info("server stopped")
}
