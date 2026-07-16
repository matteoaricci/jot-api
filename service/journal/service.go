package journal

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/matteoaricci/jot-api/models/journal"
	"github.com/matteoaricci/jot-api/repo"
	"gorm.io/gorm"
	"net/http"
)

func Delete(id uint64, userID uint64) *echo.HTTPError {
	err := repo.DeleteJournal(id, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound)
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}

func All(userID uint64, params models.JournalQueryParams) (*models.PageOfJournalVMs, *echo.HTTPError) {
	jRepos, err := repo.GetAllJournals(userID, params)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, echo.NewHTTPError(http.StatusNotFound)
		}
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	pageOfVMs := RepoToPageOfVMs(jRepos, params)

	return &pageOfVMs, nil
}

func Get(id uint64, userID uint64) (*models.JournalVM, *echo.HTTPError) {
	j, err := repo.GetJournalByID(id, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, echo.NewHTTPError(http.StatusNotFound)
		}
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	jVM := MapRepoToVM(*j)

	return &jVM, nil
}

func Create(newJournal models.CreateOrPutJournalVM, userID uint64) (*string, *echo.HTTPError) {
	j, err := repo.CreateJournal(newJournal.Title, newJournal.Description, newJournal.Completed, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
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

func Put(id uint64, journal models.CreateOrPutJournalVM, userID uint64) (*models.JournalVM, *echo.HTTPError) {
	jRepo, err := repo.UpdateJournal(id, journal.Title, journal.Description, journal.Completed, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, echo.NewHTTPError(http.StatusNotFound)
		}
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	jVM := MapRepoToVM(*jRepo)

	return &jVM, nil
}
