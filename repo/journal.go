package repo

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Journals struct {
	gorm.Model
}
type JournalRepo struct {
}

var DB *gorm.DB

func InitJournalRepo() {
	dsn := "host=localhost user=matteoaricci password=matteo101 dbname=jot_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	DB = db
}

func (j JournalRepo) GetJournalByID(id int) *gorm.DB {
	row := DB.First(&Journals{}, id)
	fmt.Println(row)
	return row
}
