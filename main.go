package main

import (
	"flag"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/matteoaricci/jot-api/api"
	"github.com/matteoaricci/jot-api/db"
	"log"
	"os"
)

func main() {
	runLocally := flag.Bool("local", false, "Run in local mode")

	e := api.ConstructServer()

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
		log.Fatal("DB_USERNAME environment variable not set")
	}

	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		log.Fatal("DB_PASSWORD environment variable not set")
	}

	dbname := os.Getenv("DB_NAME")
	if dbname == "" {
		log.Fatal("DB_NAME environment variable not set")
	}

	db.InitDB(host, port, username, password, dbname, sslmode)

	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = "8080"
	}

	fmt.Println("Listening on port: ", serverPort)

	if *runLocally {
		err := e.Start(":" + serverPort)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		lambda.Start(LambdaEchoProxy(e))
	}
}
