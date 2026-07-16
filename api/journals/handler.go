package journals

import (
	"github.com/labstack/echo/v4"
	"github.com/matteoaricci/jot-api/models/journal"
	"github.com/matteoaricci/jot-api/service/journal"
	"net/http"
	"strconv"
)

func AddRoutes(e *echo.Echo) {
	e.GET("/api/journals", func(c echo.Context) error {
		userID := c.Get("userId").(uint64)

		var params models.JournalQueryParams
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

		j, err := journal.All(userID, params)
		if err != nil {
			return c.JSON(err.Code, err)
		}

		return c.JSON(http.StatusOK, j)
	})

	e.POST("/api/journals", func(c echo.Context) error {
		userID := c.Get("userId").(uint64)

		var j models.CreateOrPutJournalVM

		err := c.Bind(&j)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		if err = models.Validate(&j); err != nil {
			return err
		}

		newJID, httpErr := journal.Create(j, userID)
		if httpErr != nil {
			return httpErr
		}

		return c.JSON(http.StatusCreated, *newJID)
	})

	e.GET("/api/journals/:id", func(c echo.Context) error {
		userID := c.Get("userId").(uint64)

		id, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		j, httpErr := journal.Get(id, userID)
		if httpErr != nil {
			return httpErr
		}

		return c.JSON(http.StatusOK, *j)
	})

	e.DELETE("/api/journals/:id", func(c echo.Context) error {
		userID := c.Get("userId").(uint64)

		id, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		httpErr := journal.Delete(id, userID)
		if httpErr != nil {
			return httpErr
		}

		return c.NoContent(http.StatusNoContent)
	})

	e.PUT("/api/journals/:id", func(c echo.Context) error {
		userID := c.Get("userId").(uint64)

		id, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		var j models.CreateOrPutJournalVM

		err = c.Bind(&j)
		if err != nil {
			return err
		}

		if err = models.Validate(&j); err != nil {
			return err
		}

		newJ, httpErr := journal.Put(id, j, userID)
		if httpErr != nil {
			return httpErr
		}

		return c.JSON(http.StatusOK, newJ)
	})
}
