package main

import (
	"fmt"
	"github/lutungp/bookstore_users-api/app"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	fmt.Printf(dbHost)

	app.StartApplication()
}
