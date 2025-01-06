package journal

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/matteoaricci/jot-api/models/journal"
	"github.com/matteoaricci/jot-api/repo"
	"net/http"
)

func Delete(id string) *echo.HTTPError {
	err := repo.DeleteJournal(id)
	if err != nil {
		return err
	}
	return nil
}

func All() ([]models.JournalVM, *echo.HTTPError) {
	jRepos, err := repo.GetAllJournals()
	if err != nil {
		return nil, err
	}

	jVMs := MapRepoSliceToVMSlice(jRepos)

	return jVMs, nil
}

func Get(id string) (*models.JournalVM, *echo.HTTPError) {
	j, err := repo.GetJournalByID(id)
	if err != nil {
		return nil, err
	}
	if j == nil {
		return nil, &echo.HTTPError{
			Code:    http.StatusNotFound,
			Message: fmt.Sprintf("Unable to find journal with id %s", id),
		}
	}

	jVM := MapRepoToVM(*j)

	return &jVM, nil
}

func Create(newJournal models.CreateOrPutJournalVM) (*string, *echo.HTTPError) {
	jRepo, err := repo.CreateJournal(newJournal.Title, newJournal.Description)
	if err != nil {
		return nil, err
	}
	if jRepo == nil {
		return nil, err
	}

	jVM := MapRepoToVM(*jRepo)

	return &jVM.ID, nil
}

func Put(id string, journal models.CreateOrPutJournalVM) (*models.JournalVM, *echo.HTTPError) {
	jRepo, err := repo.UpdateJournal(id, journal.Title, journal.Description)
	if err != nil {
		return nil, err
	}
	jVM := MapRepoToVM(*jRepo)

	return &jVM, nil
}
