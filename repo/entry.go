package repo

import (
	"gorm.io/gorm"
	"time"
)

// Entry represents the DB model for a journal entry
type Entry struct {
	ID        uint64         `gorm:"primary_key;auto_increment" json:"id"`
	CreatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Content   string         `gorm:"type:text" json:"content"`
	JournalID string         `gorm:"type:text" json:"journal_id"`
	CreatedBy string         `gorm:"type:text" json:"created_by"`
	UpdatedBy string         `gorm:"type:text" json:"updated_by"`
}

// CreateEntry inserts a new entry for the given journal
func CreateEntry(content string, journalID string) (*Entry, error) {
	e := Entry{Content: content, JournalID: journalID}
	if err := db.Create(&e).Error; err != nil {
		return nil, err
	}
	return &e, nil
}

// GetEntriesByJournalID returns all entries for a journal
func GetEntriesByJournalID(journalID string) ([]Entry, error) {
	var entries []Entry
	if err := db.Where("journal_id = ?", journalID).Find(&entries).Error; err != nil {
		return nil, err
	}
	return entries, nil
}
