package repo

import (
	"context"
	"log/slog"
)

type SystemStats struct {
	TotalUsers        int64
	TotalJournals     int64
	TotalEntries      int64
	CompletedJournals int64
}

type UserStats struct {
	UserID       uint64 `gorm:"column:user_id"`
	Email        string `gorm:"column:email"`
	JournalCount int64  `gorm:"column:journal_count"`
	EntryCount   int64  `gorm:"column:entry_count"`
}

func GetSystemStats(ctx context.Context) (*SystemStats, error) {
	slog.InfoContext(ctx, "repo.GetSystemStats")

	var s SystemStats
	d := db.WithContext(ctx)

	if err := d.Model(&Usr{}).Count(&s.TotalUsers).Error; err != nil {
		return nil, err
	}
	if err := d.Model(&Journal{}).Count(&s.TotalJournals).Error; err != nil {
		return nil, err
	}
	if err := d.Model(&Entry{}).Count(&s.TotalEntries).Error; err != nil {
		return nil, err
	}
	if err := d.Model(&Journal{}).Where("completed = ?", "true").Count(&s.CompletedJournals).Error; err != nil {
		return nil, err
	}

	return &s, nil
}

func GetUserStats(ctx context.Context) ([]UserStats, error) {
	slog.InfoContext(ctx, "repo.GetUserStats")

	var stats []UserStats
	err := db.WithContext(ctx).
		Table("usr").
		Select(`usr.id AS user_id,
			usr.email AS email,
			COUNT(DISTINCT journal.id) AS journal_count,
			COUNT(DISTINCT entry.id) AS entry_count`).
		Joins("LEFT JOIN journal ON journal.user_id = usr.id AND journal.deleted_at IS NULL").
		Joins("LEFT JOIN entry ON CAST(entry.journal_id AS INTEGER) = journal.id AND entry.deleted_at IS NULL").
		Where("usr.deleted_at IS NULL").
		Group("usr.id, usr.email").
		Order("usr.id").
		Scan(&stats).Error
	if err != nil {
		return nil, err
	}

	return stats, nil
}
