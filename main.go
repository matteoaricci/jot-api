package main

import (
	"flag"
	"log/slog"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/matteoaricci/jot-api/api"
	"github.com/matteoaricci/jot-api/db"
	"github.com/matteoaricci/jot-api/logger"
)

func main() {
	logger.Init()

	runLocally := flag.Bool("local", true, "Run in local mode")

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		slog.Error("JWT_SECRET environment variable not set")
		os.Exit(1)
	}

	e := api.ConstructServer(jwtSecret)

	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost"
	}

	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "5432"
	}

	sslmode := os.Getenv("DB_SSLMODE")
	if sslmode == "" {
		sslmode = "disable"
	}

	username := os.Getenv("DB_USERNAME")
	if username == "" {
		slog.Error("DB_USERNAME environment variable not set")
		os.Exit(1)
	}

	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		slog.Error("DB_PASSWORD environment variable not set")
		os.Exit(1)
	}

	dbname := os.Getenv("DB_NAME")
	if dbname == "" {
		slog.Error("DB_NAME environment variable not set")
		os.Exit(1)
	}

	db.InitDB(host, port, username, password, dbname, sslmode)

	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = "8080"
	}

	slog.Info("starting server", slog.String("port", serverPort))

	if *runLocally {
		err := e.Start(":" + serverPort)
		if err != nil {
			slog.Error("server stopped", slog.String("error", err.Error()))
			os.Exit(1)
		}
	} else {
		lambda.Start(LambdaEchoProxy(e))
	}
}
