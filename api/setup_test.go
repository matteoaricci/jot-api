package api

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/matteoaricci/jot-api/db"
	"github.com/pressly/goose/v3"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"testing"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	gormPG "gorm.io/driver/postgres"
)

var Server *echo.Echo

var TestDB *gorm.DB

func TestMain(m *testing.M) {
	ctx := context.Background()

	dbUsername := os.Getenv("DB_USERNAME")
	if dbUsername == "" {
		log.Fatal("DB_USER environment variable not set")
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		log.Fatal("DB_PASSWORD environment variable not set")
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		log.Fatal("DB_NAME environment variable not set")
	}

	postgresContainer, err := postgres.Run(ctx,
		"postgres:16-alpine",
		postgres.WithDatabase(dbName),
		postgres.WithUsername(dbUsername),
		postgres.WithPassword(dbPassword),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second)),
	)

	defer func() {
		if err := testcontainers.TerminateContainer(postgresContainer); err != nil {
			log.Printf("failed to terminate container: %s", err)
		}
	}()

	if err != nil {
		log.Printf("failed to start container: %s", err)
		return
	}
	connStr, err := postgresContainer.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		log.Printf("failed to get connection string: %s", err)
		return
	}

	gormDB, err := gorm.Open(gormPG.Open(connStr), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Printf("failed to connect to database: %s", err)
		return
	}
	sqlDB, err := gormDB.DB()
	if err != nil {
		log.Printf("failed to get database connection string: %s", err)
		return
	}

	for {
		if _, err := os.Stat("main.go"); err == nil {
			break
		}
		if err := os.Chdir(".."); err != nil {
			panic(err)
		}
	}

	err = goose.Up(sqlDB, "db/migrations")
	if err != nil {
		log.Printf("failed to up migrations: %s", err)
		return
	}

	db.InitRepo(gormDB)

	s, err := os.ReadFile("testdata/integration_test_init.sql")
	if err != nil {
		log.Printf("failed to read file: %s", err)
		return
	}

	_, err = sqlDB.Exec(string(s))
	if err != nil {
		log.Printf("failed to seed database: %s", err)
		return
	}

	Server = ConstructServer()

	TestDB = gormDB

	c := m.Run()

	if err := testcontainers.TerminateContainer(postgresContainer); err != nil {
		log.Printf("failed to terminate container: %s", err)
	}

	os.Exit(c)
}
