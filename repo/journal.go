package repo

import (
	"fmt"
	"gorm.io/gorm"
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

func GetAllJournals() ([]Journal, error) {
	var journal []Journal

	row := db.Find(&journal)
	if row.Error != nil {
		return nil, row.Error
	}

	return journal, nil
}

func GetJournalByID(id string) (*Journal, error) {
	var journal Journal

	row := db.First(&journal, id)
	if row.Error != nil {
		return nil, row.Error
	}

	return &journal, nil
}

func CreateJournal(title string, description string) (*Journal, error) {
	journal := Journal{Title: title, Description: description}
	err := db.Create(&journal).Error

	if err != nil {
		return nil, err
	}

	return &journal, nil
}

func UpdateJournal(id uint64, title string, description string) (*Journal, error) {
	row := db.First(&Journal{}, id)
	if row.Error != nil {
		return nil, row.Error
	}
	journal := Journal{ID: id, Title: title, Description: description}

	if err := db.Save(&journal).Error; err != nil {
		return nil, err
	}

	return &journal, nil
}

func DeleteJournal(id string) error {
	row := db.First(&Journal{}, id)
	if row.Error != nil {
		return row.Error
	}

	var journal Journal

	err := db.Delete(&journal, id).Error
	if err != nil {
		fmt.Sprintf("hello")
		return err
	}

	return nil
}
