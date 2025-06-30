package marketdata

import (
	"time"

	"github.com/alpacahq/alpaca-trade-api-go/v3/marketdata"
)

func GetStockQuote(symbols string, quoteLimit int, quoteStartDate string) ([]marketdata.Quote, error) {
	startDate, err := time.Parse(time.UnixDate, quoteStartDate)
	if err != nil {
		panic(err)
	}

	quote, err := marketdata.GetQuotes(symbols, marketdata.GetQuotesRequest{
		Start:      startDate,
		TotalLimit: quoteLimit,
	})
	if err != nil {
		return nil, err
	}

	return quote, nil
}
