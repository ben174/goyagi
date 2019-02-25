package movies

import (
	"github.com/ben174/goyagi/pkg/application"
	"github.com/labstack/echo"
)

// RegisterRoutes takes in an Echo router and registers routes onto it.
func RegisterRoutes(e *echo.Echo, app application.App) {
	g := e.Group("/movies")

	// RegiserRoutes takes in an application.App struct and we wrap it around an
	// internal handler struct. This allows us to create a methods over the
	// handler struct that act as our endpoint handler functions. Doing this
	// allows us to access our App struct (which contains references to all
	// objects that need to only be instantiated once) from all of our handlers.
	// If we ever need to add new dependencies to our App struct, we'll
	// automatically make them available to all our handlers through this
	// construct.
	h := handler{app}

	g.GET("", h.listHandler)
	g.GET("/:id", h.retrieveHandler)
}
