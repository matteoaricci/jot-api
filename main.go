package main

import (
	"fmt"
	"net/http"
	"os"

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

	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = "8080"
	}

	fmt.Println("Listening on port: ", serverPort)


	e.Start(":" + serverPort)
}