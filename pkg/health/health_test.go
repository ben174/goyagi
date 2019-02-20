package health

import (
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/labstack/echo"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
)

func TestHealthRoute(t *testing.T) {
    // Create a new echo router
    e := echo.New()

    // Regier our echo routes onto our new router
    RegisterRoutes(e)

    // Create a new test request that will be given to our server
    req, err := http.NewRequest("GET", "/health", nil)
    require.NoError(t, err)

    // Create a new test HTTP Recoder. The httptest.ResponseRecorder is an
    // implementation of http.ResponseWriter that records its mutations for
    // later inspection in tests.
    w := httptest.NewRecorder()

    e.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code, "incorrect status code")
    assert.Equal(t, `{"healthy":true}`, w.Body.String(), "incorrect response")
}
