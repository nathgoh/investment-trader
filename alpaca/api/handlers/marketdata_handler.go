package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/nathgoh/investment-trader/alpaca/internal/marketdata"
)

func GetStockQuote(w http.ResponseWriter, r *http.Request) {
	// Extract symbol from URL path, e.g. /api/v1/marketdata/quotes/AAPL
	parts := strings.Split(r.URL.Path, "/")
	symbol := parts[len(parts)-1]

	limitStr := r.URL.Query().Get("limit")
	if limitStr == "" {
		limitStr = "1"
	}

	startDateStr := r.URL.Query().Get("startDate")
	if startDateStr == "" {
		startDateStr = time.Now().Format("1/2/2006")
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		http.Error(w, "Invalid limit parameter", http.StatusBadRequest)
		return
	}

	quotes, err := marketdata.GetStockQuote(symbol, limit, startDateStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(quotes)
}
