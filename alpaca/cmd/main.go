package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/nathgoh/investment-trader/alpaca/api/routes"
)

func main() {
	// Load .env file from the project root
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file. Make sure it's in the project root.")
	}

	mux := routes.SetupRoutes()

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
