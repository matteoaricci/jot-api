package db

import (
	"fmt"
	"github.com/matteoaricci/jot-api/repo"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

func InitDB(host string, port string, user string, password string, dbName string, sslmode string) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, user, password, dbName, port, sslmode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		log.Fatal(err)
	}

	InitRepo(db)

	return db
}

func InitRepo(db *gorm.DB) {
	repo.InitJournalRepo(db)
}
