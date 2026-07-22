package models

type SystemStatsVM struct {
	TotalUsers        int64 `json:"totalUsers"`
	TotalJournals     int64 `json:"totalJournals"`
	TotalEntries      int64 `json:"totalEntries"`
	CompletedJournals int64 `json:"completedJournals"`
}

type UserStatsVM struct {
	UserID       uint64 `json:"userId"`
	Email        string `json:"email"`
	JournalCount int64  `json:"journalCount"`
	EntryCount   int64  `json:"entryCount"`
}
