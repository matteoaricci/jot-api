package users

import (
	"github.com/labstack/echo/v4"
	models "github.com/matteoaricci/jot-api/models/user"
	"github.com/matteoaricci/jot-api/service/user"
	"net/http"
)

func AddRoutes(e *echo.Echo) {
	e.POST("/api/authenticate", authenticate)
	e.POST("api/sign-up", signUp)
}

func signUp(c echo.Context) error {
	var params models.SignUpUserVM
	bindErr := c.Bind(&params)
	if bindErr != nil {
		return c.JSON(http.StatusBadRequest, bindErr.Error())
	}

}

func authenticate(c echo.Context) error {
	var params models.AuthenticateUserVM
	bindErr := c.Bind(&params)
	if bindErr != nil {
		return c.JSON(http.StatusBadRequest, bindErr.Error())
	}

	err := user.AuthenticateUser(params.Email, params.Password)
	if err != nil {
		return c.JSON(err.Code, err.Error())
	}

	return c.NoContent(http.StatusOK)
}
