package marketdata

import (
	"log"
	"os"
	"time"

	"github.com/alpacahq/alpaca-trade-api-go/v3/marketdata"
	"github.com/joho/godotenv"
	"github.com/nathgoh/investment-trader/alpaca/internal/utils"
)

var (
	client *marketdata.Client
)

func init() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	client = marketdata.NewClient(marketdata.ClientOpts{
		APIKey:    os.Getenv("ALPACA_PAPER_API_KEY"),
		APISecret: os.Getenv("ALPACA_PAPER_SECRET_KEY"),
		BaseURL:   utils.MARKETDATA_BASE_URL,
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
