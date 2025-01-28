package journals

import (
	"github.com/labstack/echo/v4"
	"github.com/matteoaricci/jot-api/models/journal"
	"github.com/matteoaricci/jot-api/service/journal"
	"net/http"
)

func AddRoutes(e *echo.Echo) {
	e.GET("/api/journals", func(c echo.Context) error {
		var params models.JournalQueryParams
		bindErr := c.Bind(&params)
		if bindErr != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}
		if err := models.Validate(&params); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		j, err := journal.All(params)
		if err != nil {
			return c.JSON(err.Code, err)
		}

		return c.JSON(http.StatusOK, j)
	})

	e.POST("/api/journals", func(c echo.Context) error {
		var j models.CreateOrPutJournalVM

		err := c.Bind(&j)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		if err = models.Validate(&j); err != nil {
			return err
		}

		newJID, httpErr := journal.Create(j)

		if httpErr != nil {
			return httpErr
		}

		return c.JSON(http.StatusCreated, *newJID)
	})

	e.GET("/api/journals/:id", func(c echo.Context) error {
		id := c.Param("id")

		j, err := journal.Get(id)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, *j)
	})

	e.DELETE("/api/journals/:id", func(c echo.Context) error {
		id := c.Param("id")

		err := journal.Delete(id)
		if err != nil {
			return err
		}

		return c.NoContent(http.StatusNoContent)
	})

	e.PUT("/api/journals/:id", func(c echo.Context) error {
		id := c.Param("id")

		var j models.CreateOrPutJournalVM

		err := c.Bind(&j)
		if err != nil {
			return err
		}

		if err = models.Validate(&j); err != nil {
			return err
		}

		newJ, httpErr := journal.Put(id, j)
		if httpErr != nil {
			return httpErr
		}

		return c.JSON(http.StatusOK, newJ)
	})
}
