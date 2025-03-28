package users

import (
	"github.com/labstack/echo/v4"
	models "github.com/matteoaricci/jot-api/models/user"
	"net/http"
)

func AddRoutes(e *echo.Echo) {
	e.POST("/api/authenticate", authenticate)
}

func authenticate(c echo.Context) error {
	
	var u models.AuthenticateUserVM

	err := c.Bind(&u)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
}
