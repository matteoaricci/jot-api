package journal

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/matteoaricci/jot-api/models/journal"
	"github.com/matteoaricci/jot-api/repo"
	"net/http"
	"strconv"
)

func Delete(id string) *echo.HTTPError {
	err := repo.DeleteJournal(id)
	if err != nil {
		return err
	}

	return nil
}

func All() ([]models.JournalVM, *echo.HTTPError) {
	journals, err := repo.GetAllJournals()

	if err != nil {
		return nil, err
	}

	return journals
}

func Get(id string) (*models.JournalVM, *echo.HTTPError) {

	j, jErr := repo.GetJournalByID(id)
	if jErr != nil {
		return nil, jErr
	}
	if j == nil {
		return nil, &echo.HTTPError{
			Code:    http.StatusNotFound,
			Message: fmt.Sprintf("Unable to find journal with id %s", id),
		}
	}

	jVM := models.JournalVM{
		Title:       j.Title,
		Description: j.Description,
		ID:          convertUintToString(j.ID),
	}

	return &jVM, nil
}

func Create(newJournal models.CreateOrPutJournalVM) (*models.JournalVM, *echo.HTTPError) {
	journal, err := repo.CreateJournal(newJournal.Title, newJournal.Description)

	if err != nil {
		return nil, err
	}

	return &journal, nil
}

func Put(id string, journal models.CreateOrPutJournalVM) (*models.JournalVM, *echo.HTTPError) {
	j := findJournal(id)

	if j == nil {
		return nil, &echo.HTTPError{
			Code:    http.StatusNotFound,
			Message: fmt.Sprintf("Unable to find journal with id %s", id),
		}
	}

	j.Title = journal.Title
	j.Description = journal.Description

	return j, nil
}

func convertUintToString(input uint) string {
	return strconv.FormatUint(uint64(input), 10)
}
