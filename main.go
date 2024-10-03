package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/api/healthz", func(c echo.Context) error  {
		res := struct {
			Status string `json:"status"`
		} {
			Status: "OK",
		}

		return c.JSON(http.StatusOK, res)
	})

	e.Start(":8080")
}