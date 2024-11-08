package journal

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/matteoaricci/jot-api/models/journal"
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
	j := findJournal(id)

	if j == nil {
		return nil, &echo.HTTPError{
			Code:    http.StatusNotFound,
			Message: fmt.Sprintf("Unable to find journal with id %s", id),
		}
	}

	return j, nil
}

func Create(newJournal models.CreateJournalVM) (*models.JournalVM, *echo.HTTPError) {
	id := len(existingJournals)
	j := models.JournalVM{
		Title:       newJournal.Title,
		Description: newJournal.Description,
		ID:          strconv.Itoa(id),
	}

	return &j, nil
}

func Patch(id string, journal models.JournalVM) (*models.JournalVM, *echo.HTTPError) {
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
