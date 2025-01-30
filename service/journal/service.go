package journal

import (
	"github.com/labstack/echo/v4"
	"github.com/matteoaricci/jot-api/models/journal"
	"github.com/matteoaricci/jot-api/repo"
	"net/http"
	"strconv"
)

func Delete(id string) *echo.HTTPError {
	err := repo.DeleteJournal(id)
	if err != nil {
		if err.Error() == "record not found" {
			return echo.NewHTTPError(http.StatusNotFound)
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}

func All(params models.JournalQueryParams) (*models.PageOfJournalVMs, *echo.HTTPError) {
	jRepos, err := repo.GetAllJournals(params)
	if err != nil {
		if err.Error() == "record not found" {
			return nil, echo.NewHTTPError(http.StatusNotFound)
		}
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	pageOfVMs := RepoToPageOfVMs(jRepos, params)

	return &pageOfVMs, nil
}

func Get(id string) (*models.JournalVM, *echo.HTTPError) {
	j, err := repo.GetJournalByID(id)
	if err != nil {
		if err.Error() == "record not found" {
			return nil, echo.NewHTTPError(http.StatusNotFound)
		}
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	jVM := MapRepoToVM(*j)

	return &jVM, nil
}

func Create(newJournal models.CreateOrPutJournalVM) (*string, *echo.HTTPError) {
	j, err := repo.CreateJournal(newJournal.Title, newJournal.Description, newJournal.Completed)
	if err != nil {
		if err.Error() == "record not found" {
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

func Put(id string, journal models.CreateOrPutJournalVM) (*models.JournalVM, *echo.HTTPError) {
	id64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	jRepo, err := repo.UpdateJournal(id64, journal.Title, journal.Description, journal.Completed)
	if err != nil {
		if err.Error() == "record not found" {
			return nil, echo.NewHTTPError(http.StatusNotFound)
		}
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	jVM := MapRepoToVM(*jRepo)

	return &jVM, nil
}
