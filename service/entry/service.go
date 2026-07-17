package entry

import (
	"context"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	entryModels "github.com/matteoaricci/jot-api/models/entry"
	"github.com/matteoaricci/jot-api/repo"
)

// GetByJournalID retrieves all entries for a given journal
func GetByJournalID(ctx context.Context, journalID string) ([]entryModels.EntryVM, *echo.HTTPError) {
	slog.InfoContext(ctx, "entry.GetByJournalID", slog.String("journalId", journalID))

	es, err := repo.GetEntriesByJournalID(ctx, journalID)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	vms := MapRepoSliceToVMSlice(es)
	return vms, nil
}

// Create adds a new entry to a journal
func Create(ctx context.Context, journalID string, vm entryModels.CreateEntryVM) (*string, *echo.HTTPError) {
	slog.InfoContext(ctx, "entry.Create", slog.String("journalId", journalID))

	e, err := repo.CreateEntry(ctx, vm.Content, journalID)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	id := strconv.FormatUint(e.ID, 10)
	return &id, nil
}
