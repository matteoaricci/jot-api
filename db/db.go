package db

import (
	"embed"
	"fmt"
	"github.com/matteoaricci/jot-api/repo"
	"github.com/pressly/goose/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

//go:embed migrations/*.sql
var embedMigrationsFS embed.FS

func InitDB(host string, port string, user string, password string, dbName string, sslmode string) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, user, password, dbName, port, sslmode)
	log.Printf("Connecting to DB: %s", dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		log.Fatal(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	goose.SetBaseFS(embedMigrationsFS)

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	if err := goose.Up(sqlDB, "migrations"); err != nil {
		panic(err)
	}

	InitRepo(db)

	return db
}

func InitRepo(db *gorm.DB) {
	repo.InitJournalRepo(db)
}
