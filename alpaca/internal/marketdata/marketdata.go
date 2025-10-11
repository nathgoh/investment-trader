package marketdata

import (
	"log"
	"os"
	"time"

	"github.com/alpacahq/alpaca-trade-api-go/v3/marketdata"
	"github.com/joho/godotenv"
)

var (
	client *marketdata.Client
)

// Initialize the market data client
func init() {
	// Load .env file
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println("Warning: Error loading .env file")
	}

	// Get API keys from environment variables (use paper keys for market data)
	apiKey := os.Getenv("ALPACA_PAPER_API_KEY")
	apiSecret := os.Getenv("ALPACA_PAPER_SECRET_KEY")

	if apiKey == "" || apiSecret == "" {
		log.Println("Warning: API key or secret not found in .env file")
		return
	}

	client = marketdata.NewClient(marketdata.ClientOpts{
		APIKey:    apiKey,
		APISecret: apiSecret,
	})
}

func GetStockQuote(symbols string, quoteLimit int, quoteStartDate string) ([]marketdata.Quote, error) {
	startDate, err := time.Parse("1/2/2006", quoteStartDate) // Layout: M/D/YYYY
	if err != nil {
		return nil, err
	}

	quotes, err := client.GetQuotes(symbols, marketdata.GetQuotesRequest{
		Start:      startDate,
		TotalLimit: quoteLimit,
	})
	if err != nil {
		return nil, err
	}

	return quotes, nil
}
