package server

import (
    "context"
    "fmt"
    "net/http"

    "github.com/ben174/goyagi/pkg/signals"
    "github.com/ben174/goyagi/pkg/health"
    "github.com/labstack/echo"
    "github.com/lob/logger-go"
)

// New returns a new HTTP server with the registered routes.
func New() *http.Server {
    log := logger.New()

    e := echo.New()

    health.RegisterRoutes(e)

    srv := &http.Server{
        Addr:    fmt.Sprintf(":%d", 3000),
        Handler: e,
    }

    // signals.Setup() returns a channel we can wait until it's closed before we
    // shutdown our server
    graceful := signals.Setup()

    // start a goroutine that will wait for the graceful channel to close.
    // Becase this happens in a goroutine it will run concurrently with our
    // server but will not block the execution of this function.
    go func() {
        <-graceful
        err := srv.Shutdown(context.Background())
        if err != nil {
            log.Err(err).Error("server shutdown")
        }
    }()

    return srv
}

