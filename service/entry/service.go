package entry

import (
	"github.com/labstack/echo/v4"
	entryModels "github.com/matteoaricci/jot-api/models/entry"
	"github.com/matteoaricci/jot-api/repo"
	"net/http"
	"strconv"
)

// GetByJournalID retrieves all entries for a given journal
func GetByJournalID(journalID string) ([]entryModels.EntryVM, *echo.HTTPError) {
	es, err := repo.GetEntriesByJournalID(journalID)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	vms := MapRepoSliceToVMSlice(es)
	return vms, nil
}

// Create adds a new entry to a journal
func Create(journalID string, vm entryModels.CreateEntryVM) (*string, *echo.HTTPError) {
	e, err := repo.CreateEntry(vm.Content, journalID)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	id := strconv.FormatUint(e.ID, 10)
	return &id, nil
}
