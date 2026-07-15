package auth

import (
	"github.com/labstack/echo/v4"
	models "github.com/matteoaricci/jot-api/models/user"
	"github.com/matteoaricci/jot-api/service/auth"
	"net/http"
)

type Handler struct {
	jwtSecret string
}

func AddRoutes(e *echo.Echo, jwtSecret string) {
	h := &Handler{jwtSecret: jwtSecret}
	e.POST("/api/public/authenticate", h.authenticate)
	e.POST("/api/public/sign-up", h.signUp)
}

func (h *Handler) signUp(c echo.Context) error {
	var params models.SignUpUserVM
	bindErr := c.Bind(&params)
	if bindErr != nil {
		return c.JSON(http.StatusBadRequest, bindErr.Error())
	}

	err := auth.SignUpUser(params.FirstName, params.LastName, params.Email, params.Password, params.Role)
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}

func (h *Handler) authenticate(c echo.Context) error {
	var params models.AuthenticateUserVM
	bindErr := c.Bind(&params)
	if bindErr != nil {
		return c.JSON(http.StatusBadRequest, bindErr.Error())
	}

	t, err := auth.AuthenticateUser(params.Email, params.Password, h.jwtSecret)
	if err != nil {
		return c.JSON(err.Code, err.Error())
	}

	cookie := auth.CreateAuthCookie(*t)

	c.SetCookie(cookie)

	return c.NoContent(http.StatusOK)
}
