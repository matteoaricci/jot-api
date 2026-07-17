package repo

import (
	"context"
	"log/slog"
	"time"

	models "github.com/matteoaricci/jot-api/models/journal"
	"gorm.io/gorm"
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

func GetAllJournals(ctx context.Context, userID uint64, params models.JournalQueryParams) ([]Journal, error) {
	slog.InfoContext(ctx, "repo.GetAllJournals", slog.Uint64("userId", userID))

	m := make(map[string]any)
	m["user_id"] = userID

	if params.Completed != "" {
		m["completed"] = params.Completed
	}

	var journal []Journal

	row := db.WithContext(ctx).Scopes(paginate(params)).Where(m).Find(&journal)
	if row.Error != nil {
		return nil, row.Error
	}

	return journal, nil
}

func GetJournalByID(ctx context.Context, id uint64, userID uint64) (*Journal, error) {
	slog.InfoContext(ctx, "repo.GetJournalByID", slog.Uint64("id", id), slog.Uint64("userId", userID))

	var journal Journal

	row := db.WithContext(ctx).Where(&Journal{ID: id, UserID: &userID}).First(&journal)
	if row.Error != nil {
		return nil, row.Error
	}

	return &journal, nil
}

func CreateJournal(ctx context.Context, title string, description string, completed models.IsCompleted, userID uint64) (*Journal, error) {
	slog.InfoContext(ctx, "repo.CreateJournal", slog.Uint64("userId", userID))

	journal := Journal{Title: title, Description: description, Completed: completed, UserID: &userID}
	err := db.WithContext(ctx).Create(&journal).Error

	if err != nil {
		return nil, err
	}

	return &journal, nil
}

func UpdateJournal(ctx context.Context, id uint64, title string, description string, completed models.IsCompleted, userID uint64) (*Journal, error) {
	slog.InfoContext(ctx, "repo.UpdateJournal", slog.Uint64("id", id), slog.Uint64("userId", userID))

	var journal Journal
	row := db.WithContext(ctx).Where(&Journal{ID: id, UserID: &userID}).First(&journal)
	if row.Error != nil {
		return nil, row.Error
	}

	journal.Title = title
	journal.Description = description
	journal.Completed = completed

	if err := db.WithContext(ctx).Save(&journal).Error; err != nil {
		return nil, err
	}

	return &journal, nil
}

func DeleteJournal(ctx context.Context, id uint64, userID uint64) error {
	slog.InfoContext(ctx, "repo.DeleteJournal", slog.Uint64("id", id), slog.Uint64("userId", userID))

	var journal Journal
	row := db.WithContext(ctx).Where(&Journal{ID: id, UserID: &userID}).First(&journal)
	if row.Error != nil {
		return row.Error
	}

	err := db.WithContext(ctx).Delete(&journal).Error
	if err != nil {
		return err
	}

	return nil
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
