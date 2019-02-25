package test

import (
    "fmt"
    "strings"

    "github.com/labstack/echo"
)

// FilterRoutes takes in a slice of Echo routes and returns a slice of strings
// with handler routes.
func FilterRoutes(routes []*echo.Route) []string {
    rs := []string{}

    for _, r := range routes {
        if strings.Contains(r.Name, "*handler") {
            rs = append(rs, fmt.Sprintf("%s %s", r.Method, r.Path))
        }
    }

    return rs
}