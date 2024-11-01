package journals

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/matteoaricci/jot-api/service/journal"
	"net/http"
)

func Create(e *echo.Echo) {
	g := e.Group("/api/journals")

	g.GET("/", func(c echo.Context) error {
		j, err := journal.GetAllJournals()
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, j)
	})

	g.POST("/", func(c echo.Context) error {
		var j journal.Journal
		err := c.Bind(&j)

		newJs, err := journal.CreateJournal(j)

		if err != nil {
			return err
		}

		return c.JSON(http.StatusCreated, newJs)
	})

	g.GET("/:id", func(c echo.Context) error {
		id := c.Param("id")

		j, err := journal.GetJournal(id)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, *j)
	})

	g.DELETE("/:id", func(c echo.Context) error {
		id := c.Param("id")

		j, err := journal.DeleteJournal(id)
		if err != nil {
			return err
		}

		res := struct {
			Message  string            `json:"message"`
			Journals []journal.Journal `json:"journals"`
		}{
			Message:  fmt.Sprintf("Hey diva! We did delete journal with id: %s.", id),
			Journals: *j,
		}

		return c.JSON(http.StatusOK, res)
	})

	g.PATCH("/:id", func(c echo.Context) error {
		id := c.Param("id")

		var j journal.Journal

		err := c.Bind(&j)
		if err != nil {
			return err
		}

		newJ, err := journal.PatchJournal(id, j)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, newJ)
	})
}
