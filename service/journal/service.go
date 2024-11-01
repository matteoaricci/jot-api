package journal

import (
	"errors"
	"strconv"
)

type Journal struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ID          string `json:"id"`
}

var existingJournals = []Journal{{
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

func DeleteJournal(id string) (*[]Journal, error) {
	j := findJournal(id)
	if j == nil {
		return nil, errors.New("journal not found")
	}

	filteredJournals := make([]Journal, 0)
	for _, v := range existingJournals {
		if v.ID != id {
			filteredJournals = append(filteredJournals, v)
		}
	}

	return &filteredJournals, nil
}

func GetAllJournals() (*[]Journal, error) {
	return &existingJournals, nil
}

func GetJournal(id string) (*Journal, error) {
	j := findJournal(id)

	if j == nil {
		return nil, errors.New("journal not found")
	}

	return j, nil
}

func CreateJournal(newJournal Journal) (*[]Journal, error) {
	id := len(existingJournals)
	newJournal.ID = strconv.Itoa(id)

	existingJournals = append(existingJournals, newJournal)

	return &existingJournals, nil
}

func PatchJournal(id string, journal Journal) (*Journal, error) {
	j := findJournal(id)

	if j == nil {
		return nil, errors.New("journal not found")
	}

	j.Title = journal.Title
	j.Description = journal.Description

	return j, nil
}

func findJournal(id string) *Journal {
	var j *Journal
	for _, v := range existingJournals {
		if v.ID == id {
			j = &v
		}
	}

	return j
}
