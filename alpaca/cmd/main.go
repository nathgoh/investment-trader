package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"

	"github.com/nathgoh/investment-trader/alpaca/internal/marketdata"
)

func main() {
	// Load .env file
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get API keys from environment variables
	// apiKey := os.Getenv("ALPACA_PAPER_API_KEY")
	// apiSecret := os.Getenv("ALPACA_PAPER_SECRET_KEY")

	// if apiKey == "" || apiSecret == "" {
	// 	log.Fatal("API key or secret not found in .env file")
	// }

	// client := alpaca.NewClient(alpaca.ClientOpts{
	// 	APIKey:    apiKey,
	// 	APISecret: apiSecret,
	// 	BaseURL:   "https://paper-api.alpaca.markets",
	// })
	// acct, err := client.GetAccount()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%+v\n", *acct)

	quotes, err := marketdata.GetStockQuote("AAPL", 1, "7/25/2025")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", quotes)
}
