package repo

import (
	"context"
	"log/slog"
	"time"

	"gorm.io/gorm"
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
func CreateEntry(ctx context.Context, content string, journalID string) (*Entry, error) {
	slog.InfoContext(ctx, "repo.CreateEntry", slog.String("journalId", journalID))

	e := Entry{Content: content, JournalID: journalID}
	if err := db.WithContext(ctx).Create(&e).Error; err != nil {
		return nil, err
	}
	return &e, nil
}

// GetEntriesByJournalID returns all entries for a journal
func GetEntriesByJournalID(ctx context.Context, journalID string) ([]Entry, error) {
	slog.InfoContext(ctx, "repo.GetEntriesByJournalID", slog.String("journalId", journalID))

	var entries []Entry
	if err := db.WithContext(ctx).Where("journal_id = ?", journalID).Find(&entries).Error; err != nil {
		return nil, err
	}
	return entries, nil
}
