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
	CreatedBy   string             `gorm:"type:text" json:"created_by"`
	UpdatedBy   string             `gorm:"type:text" json:"updated_by"`
	Title       string             `gorm:"type:text" `
	Description string             `gorm:"type:text"`
	Completed   models.IsCompleted `gorm:"type:is_completed, default:'unknown'" json:"completed"`
	// UserID links this journal to its owner (nullable)
	UserID *uint64 `gorm:"column:user_id" json:"userId"`
}

var db *gorm.DB

func InitJournalRepo(dB *gorm.DB) {
	db = dB
}

func GetAllJournals(userID uint64, params models.JournalQueryParams) ([]Journal, error) {
	m := make(map[string]any)
	m["user_id"] = userID

	if params.Completed != "" {
		m["completed"] = params.Completed
	}

	var journal []Journal

	row := db.Scopes(paginate(params)).Where(m).Find(&journal)
	if row.Error != nil {
		return nil, row.Error
	}

	return journal, nil
}

func GetJournalByID(id uint64, userID uint64) (*Journal, error) {
	var journal Journal

	row := db.Where(&Journal{ID: id, UserID: &userID}).First(&journal)
	if row.Error != nil {
		return nil, row.Error
	}

	return &journal, nil
}

func CreateJournal(title string, description string, completed models.IsCompleted, userID uint64) (*Journal, error) {
	journal := Journal{Title: title, Description: description, Completed: completed, UserID: &userID}
	err := db.Create(&journal).Error

	if err != nil {
		return nil, err
	}

	return &journal, nil
}

func UpdateJournal(id uint64, title string, description string, completed models.IsCompleted, userID uint64) (*Journal, error) {
	var journal Journal
	row := db.Where(&Journal{ID: id, UserID: &userID}).First(&journal)
	if row.Error != nil {
		return nil, row.Error
	}

	journal.Title = title
	journal.Description = description
	journal.Completed = completed

	if err := db.Save(&journal).Error; err != nil {
		return nil, err
	}

	return &journal, nil
}

func DeleteJournal(id uint64, userID uint64) error {
	var journal Journal
	row := db.Where(&Journal{ID: id, UserID: &userID}).First(&journal)
	if row.Error != nil {
		return row.Error
	}

	err := db.Delete(&journal).Error
	if err != nil {
		return err
	}

	return nil
}

func GetJournalsByUserID(id string, params models.JournalQueryParams) ([]Journal, error) {
	m := make(map[string]any)

	m["user_id"] = id

	var journal []Journal

	row := db.Scopes(paginate(params)).Where(m).Find(&journal)
	if row.Error != nil {
		return nil, row.Error
	}

	return journal, nil
}

func paginate(params models.JournalQueryParams) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		pageSize := params.Size
		switch {
		case pageSize < 1:
			pageSize = 10
		}

		offset := (params.Page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
