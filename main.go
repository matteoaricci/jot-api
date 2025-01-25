package main

import (
	"fmt"
	"github.com/matteoaricci/jot-api/api"
	"github.com/matteoaricci/jot-api/db"
	"log"
	"os"
)

func main() {
	e := api.ConstructServer()

	db.InitDB("localhost", 5432, "matteoaricci", "matteo101", "jot_db")

	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = "8080"
	}

	fmt.Println("Listening on port: ", serverPort)

	err := e.Start(":" + serverPort)
	if err != nil {
		log.Fatal(err)
	}
}
