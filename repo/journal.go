package repo

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type Journals struct {
	gorm.Model
	Title       string `gorm:"type:text"`
	Description string `gorm:"type:text"`
}
type JournalRepo struct {
}

var DB *gorm.DB

func InitJournalRepo(db *gorm.DB) {
	DB = db
}

func GetAllJournals() ([]Journals, *echo.HTTPError) {
	var journals []Journals

	row := DB.Find(&journals)
	if row.Error != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Failed to fetch all journals")
	}

	return journals, nil
}

func GetJournalByID(id string) (*Journals, *echo.HTTPError) {
	var journal Journals

	row := DB.First(&journal, id)
	if row.Error != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Failed to fetch journal with id %s", id))
	}

	return &journal, nil
}

func CreateJournal(title string, description string) (*Journals, *echo.HTTPError) {
	journal := Journals{Title: title, Description: description}
	if err := DB.Create(&journal).Error; err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Failed to create journal")
	}

	return &journal, nil
}

func UpdateJournal(id string, title string, description string) (*Journals, *echo.HTTPError) {
	idUint, _ := strconv.ParseUint(id, 10, 32)
	journal := Journals{Title: title, Description: description, Model: gorm.Model{ID: uint(idUint)}}

	if err := DB.Save(&journal).Error; err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Failed to update journal with id %s", id))
	}

	return &journal, nil
}

func DeleteJournal(id string) *echo.HTTPError {
	var journal Journals

	if err := DB.Delete(&journal, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Failed to delete journal with id %d", id))
	}

	return nil
}
