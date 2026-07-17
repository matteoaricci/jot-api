package users

import (
	"github.com/labstack/echo/v4"
	"github.com/matteoaricci/jot-api/service/user"
)

func AddRoutes(e *echo.Echo) {
	e.DELETE("api/users/:id", deleteUser)
}

func deleteUser(c echo.Context) error {
	id := c.Param("id")

	err := user.DeleteUser(c.Request().Context(), id)
	if err != nil {
		return err
	}

	return c.NoContent(200)
}
