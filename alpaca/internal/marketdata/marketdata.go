package marketdata

import (
	"time"

	"github.com/alpacahq/alpaca-trade-api-go/v3/marketdata"
)

var (
	client *marketdata.Client
)

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
