package users

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	models "github.com/matteoaricci/jot-api/models/user"
	"github.com/matteoaricci/jot-api/service/user"
)

func AddRoutes(e *echo.Echo) {
	e.DELETE("api/users/:id", deleteUser)
	e.PUT("/api/admin/users/:id/role", updateUserRole)
}

func updateUserRole(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var body models.UpdateUserRoleVM
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if body.Role == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "role is required")
	}

	if httpErr := user.UpdateRole(c.Request().Context(), id, body.Role); httpErr != nil {
		return httpErr
	}

	return c.NoContent(http.StatusNoContent)
}

func deleteUser(c echo.Context) error {
	id := c.Param("id")

	err := user.DeleteUser(c.Request().Context(), id)
	if err != nil {
		return err
	}

	return c.NoContent(200)
}
