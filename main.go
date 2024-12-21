package main

import (
	"fmt"
	"github.com/matteoaricci/jot-api/api/journals"
	"github.com/matteoaricci/jot-api/middleware"
	"github.com/matteoaricci/jot-api/repo"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

var DB *gorm.DB

func main() {
	e := echo.New()

	middleware.AddMiddleware(e)

	AddRouteHandlers(e)

	repo.InitJournalRepo()

	//DB = db.ConnectToDB()

	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = "8080"
	}

	fmt.Println("Listening on port: ", serverPort)

	err := e.Start(":" + serverPort)
	if err != nil {
		log.Fatal(err)
	}
}

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
