package users

import (
	"github.com/labstack/echo/v4"
	jModels "github.com/matteoaricci/jot-api/models/journal"
	"github.com/matteoaricci/jot-api/service/user"
	"net/http"
)

func AddRoutes(e *echo.Echo) {
	e.GET("/api/users/:id/journals", getJournalsByUserID)
	e.DELETE("api/users/:id", deleteUser)
}

func getJournalsByUserID(c echo.Context) error {
	id := c.Param("id")

	var params jModels.JournalQueryParams
	bindErr := c.Bind(&params)
	if bindErr != nil {
		return c.JSON(http.StatusBadRequest, bindErr.Error())
	}

	if params.Size == 0 {
		params.Size = 10
	}
	if params.Page == 0 {
		params.Page = 1
	}

	j, err := user.GetJournals(id, params)
	if err != nil {
		return c.JSON(err.Code, err.Error())
	}

	return c.JSON(http.StatusOK, j)
}

func deleteUser(c echo.Context) error {
	id := c.Param("id")

	err := user.DeleteUser(id)
	if err != nil {
		return err
	}

	return c.NoContent(200)
}
