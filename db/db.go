package db

import (
	"embed"
	"fmt"
	"log/slog"
	"os"

	"github.com/matteoaricci/jot-api/repo"
	"github.com/pressly/goose/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

//go:embed migrations/*.sql
var embedMigrationsFS embed.FS

func InitDB(host string, port string, user string, password string, dbName string, sslmode string) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, user, password, dbName, port, sslmode)
	slog.Info("connecting to database",
		slog.String("host", host),
		slog.String("port", port),
		slog.String("dbname", dbName),
		slog.String("user", user),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		slog.Error("failed to connect to database", slog.String("error", err.Error()))
		os.Exit(1)
	}

	slog.Info("database connection established")

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
