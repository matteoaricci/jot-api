package repo

import (
	"gorm.io/gorm"
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

func GetJournalByID(id int) *Journals {
	var journal Journals
	DB.First(&journal, id)
	return &journal
}
