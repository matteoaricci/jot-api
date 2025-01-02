package db

import (
	"github.com/matteoaricci/jot-api/repo"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func InitDB() {
	dsn := "host=localhost user=matteoaricci password=matteo101 dbname=jot_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	initRepo(db)
}

func initRepo(db *gorm.DB) {
	repo.InitJournalRepo(db)
}
