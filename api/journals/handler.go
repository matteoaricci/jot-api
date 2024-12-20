package journals

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/matteoaricci/jot-api/models/journal"
	"github.com/matteoaricci/jot-api/service/journal"
	"net/http"
)

func AddRoutes(e *echo.Echo) {
	g := e.Group("/api/journals")

	g.GET("", func(c echo.Context) error {
		j, err := journal.All()
		if err != nil {
			return c.JSON(err.Code, err)
		}

		return c.JSON(http.StatusOK, j)
	})

	g.POST("", func(c echo.Context) error {
		e.Validator = &models.CreateOrPutJournalValidator{Validator: validator.New()}
		var j models.CreateOrPutJournalVM

		err := c.Bind(&j)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		if err = c.Validate(&j); err != nil {
			return err
		}

		newJ, httpErr := journal.Create(j)

		if httpErr != nil {
			return httpErr
		}

		return c.JSON(http.StatusCreated, newJ)
	})

	g.GET("/:id", func(c echo.Context) error {
		id := c.Param("id")

		j, err := journal.Get(id)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, *j)
	})

	g.DELETE("/:id", func(c echo.Context) error {
		id := c.Param("id")

		_, err := journal.Delete(id)
		if err != nil {
			return err
		}

		return c.NoContent(http.StatusNoContent)
	})

	g.PUT("/:id", func(c echo.Context) error {
		e.Validator = &models.CreateOrPutJournalValidator{Validator: validator.New()}
		id := c.Param("id")

		var j models.CreateOrPutJournalVM

		err := c.Bind(&j)
		if err != nil {
			return err
		}

		if err = c.Validate(&j); err != nil {
			return err
		}

		newJ, httpErr := journal.Put(id, j)
		if httpErr != nil {
			return httpErr
		}

		return c.JSON(http.StatusOK, newJ)
	})
}
