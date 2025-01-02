package journal

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/matteoaricci/jot-api/models/journal"
	"github.com/matteoaricci/jot-api/repo"
	"net/http"
	"strconv"
)

var existingJournals = []models.JournalVM{{
	Title:       "Psychopomp",
	Description: "Japanese Breakfast's first album",
	ID:          "1",
}, {
	Title:       "Soft Sounds from Another Planet",
	Description: "Absolute banger followup",
	ID:          "2",
}, {
	Title:       "Jubilee",
	Description: "Here Michelle Zauner asks: what if joy was as complex as grief",
	ID:          "3",
}}

func Delete(id string) ([]models.JournalVM, *echo.HTTPError) {
	j := findJournal(id)
	if j == nil {
		return nil, &echo.HTTPError{
			Code:    http.StatusNotFound,
			Message: fmt.Sprintf("Unable to find journal with id %s", id),
		}
	}

	filteredJournals := make([]models.JournalVM, 0)
	for _, v := range existingJournals {
		if v.ID != id {
			filteredJournals = append(filteredJournals, v)
		}
	}

	return filteredJournals, nil
}

func All() ([]models.JournalVM, *echo.HTTPError) {
	return existingJournals, nil
}

func Get(id string) (*models.JournalVM, *echo.HTTPError) {
	intId, err := strconv.Atoi(id)
	if err != nil {
		return nil, &echo.HTTPError{}
	}

	j := repo.GetJournalByID(intId)

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
	id := len(existingJournals)
	j := models.JournalVM{
		Title:       newJournal.Title,
		Description: newJournal.Description,
		ID:          strconv.Itoa(id),
	}

	return &j, nil
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

func findJournal(id string) *models.JournalVM {
	for _, v := range existingJournals {
		if v.ID == id {
			return &v
		}
	}

	return nil
}

func convertUintToString(input uint) string {
	return strconv.FormatUint(uint64(input), 10)
}
