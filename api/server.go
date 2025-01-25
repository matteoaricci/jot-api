package api

import (
	"github.com/labstack/echo/v4"
	"github.com/matteoaricci/jot-api/api/journals"
	"github.com/matteoaricci/jot-api/middleware"
	"net/http"
)

func AddRouteHandlers(e *echo.Echo) {
	journals.AddRoutes(e)

	e.RouteNotFound("/*", func(c echo.Context) error {
		return c.NoContent(http.StatusNotFound)
	})

	e.GET("/api/healthz", func(c echo.Context) error {
		res := struct {
			Status string `json:"status"`
		}{
			Status: "OK",
		}

		return c.JSON(http.StatusOK, res)
	})
}

func ConstructServer() *echo.Echo {
	e := echo.New()

	middleware.AddMiddleware(e)

	AddRouteHandlers(e)

	return e
}
