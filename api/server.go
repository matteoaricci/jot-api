package api

import (
	"github.com/labstack/echo/v4"
	"github.com/matteoaricci/jot-api/api/auth"
	"github.com/matteoaricci/jot-api/api/entries"
	"github.com/matteoaricci/jot-api/api/journals"
	"github.com/matteoaricci/jot-api/api/users"
	"github.com/matteoaricci/jot-api/middleware"
	"github.com/matteoaricci/jot-api/version"
	"net/http"
)

func AddRouteHandlers(e *echo.Echo, jwtSecret string) {

	journals.AddRoutes(e)
	users.AddRoutes(e)
	entries.AddRoutes(e)
	auth.AddRoutes(e, jwtSecret)

	e.RouteNotFound("/*", func(c echo.Context) error {
		return c.NoContent(http.StatusNotFound)
	})

	e.GET("/api/public/healthz", func(c echo.Context) error {
		res := struct {
			Status string `json:"status"`
		}{
			Status: "OK",
		}

		return c.JSON(http.StatusOK, res)
	})

	e.GET("/api/public/version", func(c echo.Context) error {
		return c.JSON(http.StatusOK, version.GetInfo())
	})
}

func ConstructServer(jwtSecret string) *echo.Echo {
	e := echo.New()

	middleware.AddMiddleware(e, jwtSecret)

	AddRouteHandlers(e, jwtSecret)

	return e
}
