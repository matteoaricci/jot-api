package repo

import (
	models "github.com/matteoaricci/jot-api/models/journal"
	"gorm.io/gorm"
	"time"
)

type Journal struct {
	ID          uint64             `gorm:"primary_key;auto_increment" json:"id"`
	CreatedAt   time.Time          `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time          `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt   gorm.DeletedAt     `gorm:"index" json:"deleted_at"`
	Title       string             `gorm:"type:text" `
	Description string             `gorm:"type:text"`
	Completed   models.IsCompleted `gorm:"type:is_completed, default:'unknown'" json:"completed"`
}

var db *gorm.DB

func InitJournalRepo(dB *gorm.DB) {
	db = dB
}

func GetAllJournals(params models.JournalQueryParams) ([]Journal, error) {
	m := make(map[string]any)
	if params.Completed != "" {
		m["completed"] = params.Completed
	}

	var journal []Journal

	row := db.Where(m).Find(&journal)
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

func CreateJournal(title string, description string, completed models.IsCompleted) (*Journal, error) {
	journal := Journal{Title: title, Description: description, Completed: completed}
	err := db.Create(&journal).Error

	if err != nil {
		return nil, err
	}

	return &journal, nil
}

func UpdateJournal(id uint64, title string, description string, completed models.IsCompleted) (*Journal, error) {
	row := db.First(&Journal{}, id)
	if row.Error != nil {
		return nil, row.Error
	}
	journal := Journal{ID: id, Title: title, Description: description, Completed: completed}

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
		return err
	}

	return nil
}
