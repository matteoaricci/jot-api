package journal

import (
	"context"
	"errors"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/matteoaricci/jot-api/models/journal"
	"github.com/matteoaricci/jot-api/repo"
	"gorm.io/gorm"
)

func Delete(ctx context.Context, id uint64, userID uint64) *echo.HTTPError {
	slog.InfoContext(ctx, "journal.Delete", slog.Uint64("id", id))

	err := repo.DeleteJournal(ctx, id, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			slog.WarnContext(ctx, "journal.Delete: not found", slog.Uint64("id", id))
			return echo.NewHTTPError(http.StatusNotFound)
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}

func All(ctx context.Context, userID uint64, params models.JournalQueryParams) (*models.PageOfJournalVMs, *echo.HTTPError) {
	slog.InfoContext(ctx, "journal.All")

	jRepos, err := repo.GetAllJournals(ctx, userID, params)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			slog.WarnContext(ctx, "journal.All: not found")
			return nil, echo.NewHTTPError(http.StatusNotFound)
		}
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	pageOfVMs := RepoToPageOfVMs(jRepos, params)

	return &pageOfVMs, nil
}

func Get(ctx context.Context, id uint64, userID uint64) (*models.JournalVM, *echo.HTTPError) {
	slog.InfoContext(ctx, "journal.Get", slog.Uint64("id", id))

	j, err := repo.GetJournalByID(ctx, id, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			slog.WarnContext(ctx, "journal.Get: not found", slog.Uint64("id", id))
			return nil, echo.NewHTTPError(http.StatusNotFound)
		}
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	jVM := MapRepoToVM(*j)

	return &jVM, nil
}

func Create(ctx context.Context, newJournal models.CreateOrPutJournalVM, userID uint64) (*string, *echo.HTTPError) {
	slog.InfoContext(ctx, "journal.Create")

	j, err := repo.CreateJournal(ctx, newJournal.Title, newJournal.Description, newJournal.Completed, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			slog.WarnContext(ctx, "journal.Create: not found")
			return nil, echo.NewHTTPError(http.StatusNotFound)
		}
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if j == nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Unable to create journal")
	}

	jVM := MapRepoToVM(*j)

	return &jVM.ID, nil
}

func Put(ctx context.Context, id uint64, journal models.CreateOrPutJournalVM, userID uint64) (*models.JournalVM, *echo.HTTPError) {
	slog.InfoContext(ctx, "journal.Put", slog.Uint64("id", id))

	jRepo, err := repo.UpdateJournal(ctx, id, journal.Title, journal.Description, journal.Completed, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			slog.WarnContext(ctx, "journal.Put: not found", slog.Uint64("id", id))
			return nil, echo.NewHTTPError(http.StatusNotFound)
		}
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	jVM := MapRepoToVM(*jRepo)

	return &jVM, nil
}
