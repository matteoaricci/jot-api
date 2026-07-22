package admin

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/matteoaricci/jot-api/service/stats"
)

func AddRoutes(e *echo.Echo) {
	e.GET("/api/admin/stats", getSystemStats)
	e.GET("/api/admin/users/stats", getUserStats)
}

func getSystemStats(c echo.Context) error {
	s, httpErr := stats.GetSystemStats(c.Request().Context())
	if httpErr != nil {
		return httpErr
	}
	return c.JSON(http.StatusOK, s)
}

func getUserStats(c echo.Context) error {
	s, httpErr := stats.GetUserStats(c.Request().Context())
	if httpErr != nil {
		return httpErr
	}
	return c.JSON(http.StatusOK, s)
}
