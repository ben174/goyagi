package movies

import (
	"testing"

	"github.com/ben174/goyagi/internal/test"
	"github.com/ben174/goyagi/pkg/application"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestRegisterRoutes(t *testing.T) {
	e := echo.New()
	app := application.App{}

	RegisterRoutes(e, app)

	routes := test.FilterRoutes(e.Routes())
	assert.Len(t, routes, 2)
	assert.Contains(t, routes, "GET /movies")
	assert.Contains(t, routes, "GET /movies/:id")
}
