package routes

import (
	"encoding/json"
	"net/http"

	"github.com/nathgoh/investment-trader/alpaca/api/handlers"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "Welcome to the Alpaca Trading API"})
	})

	mux.HandleFunc("/api/v1/account/paper", handlers.GetPaperAccount)
	mux.HandleFunc("/api/v1/account/live", handlers.GetLiveAccount)
	mux.HandleFunc("/api/v1/marketdata/quotes/", handlers.GetStockQuote)

	return mux
}
