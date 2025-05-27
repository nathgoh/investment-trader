package utils

import (
	"log"
	"os"

	"github.com/alpacahq/alpaca-trade-api-go/v3/alpaca"
	"github.com/joho/godotenv"
)

func GetAccount() *alpaca.Account {
	// Load .env file
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get API keys from environment variables
	apiKey := os.Getenv("ALPACA_PAPER_API_KEY")
	apiSecret := os.Getenv("ALPACA_PAPER_SECRET_KEY")

	if apiKey == "" || apiSecret == "" {
		log.Fatal("API key or secret not found in .env file")
	}

	client := alpaca.NewClient(alpaca.ClientOpts{
		APIKey:    apiKey,
		APISecret: apiSecret,
		BaseURL:   "https://paper-api.alpaca.markets",
	})
	account, err := client.GetAccount()
	if err != nil {
		panic(err)
	}

	return account
}
