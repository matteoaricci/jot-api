package repo

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type Journal struct {
	ID          uint64         `gorm:"primary_key;auto_increment" json:"id"`
	CreatedAt   time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Title       string         `gorm:"type:text" `
	Description string         `gorm:"type:text"`
}

var db *gorm.DB

func InitJournalRepo(dB *gorm.DB) {
	db = dB
}

func GetAllJournals() ([]Journal, *echo.HTTPError) {
	var journal []Journal

	row := db.Find(&journal)
	if row.Error != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Failed to fetch all journals")
	}

	return journal, nil
}

func GetJournalByID(id string) (*Journal, *echo.HTTPError) {
	var journal Journal

	row := db.First(&journal, id)
	if row.Error != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Failed to fetch journal with id %s", id))
	}

	return &journal, nil
}

func CreateJournal(title string, description string) (*Journal, *echo.HTTPError) {
	journal := Journal{Title: title, Description: description}
	err := db.Create(&journal).Error

	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Failed to create journal")
	}

	return &journal, nil
}

func UpdateJournal(id string, title string, description string) (*Journal, *echo.HTTPError) {
	journal := Journal{Title: title, Description: description}

	if err := db.Save(&journal).Error; err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Failed to update journal with id %s", id))
	}

	return &journal, nil
}

func DeleteJournal(id string) *echo.HTTPError {
	var journal Journal

	if err := db.Delete(&journal, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Failed to delete journal with id %d", id))
	}

	return nil
}
