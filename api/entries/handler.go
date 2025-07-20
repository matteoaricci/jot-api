package entries

import (
	"net/http"

	"github.com/labstack/echo/v4"
	entryModels "github.com/matteoaricci/jot-api/models/entry"
	entryService "github.com/matteoaricci/jot-api/service/entry"
)

// AddRoutes registers entry routes on the Echo router
func AddRoutes(e *echo.Echo) {
	e.GET("/api/journals/:id/entries", getEntriesByJournalID)
	e.POST("/api/journals/:id/entries", createEntry)
}

func getEntriesByJournalID(c echo.Context) error {
	journalID := c.Param("id")
	entries, httpErr := entryService.GetByJournalID(journalID)
	if httpErr != nil {
		return c.JSON(httpErr.Code, httpErr.Message)
	}
	return c.JSON(http.StatusOK, entries)
}

func createEntry(c echo.Context) error {
	journalID := c.Param("id")
	var vm entryModels.CreateEntryVM
	if err := c.Bind(&vm); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := entryModels.ValidateEntry(&vm); err != nil {
		return err
	}
	newID, httpErr := entryService.Create(journalID, vm)
	if httpErr != nil {
		return c.JSON(httpErr.Code, httpErr.Message)
	}
	return c.JSON(http.StatusCreated, *newID)
}
