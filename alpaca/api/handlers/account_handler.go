package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/nathgoh/investment-trader/alpaca/internal/utils"
)

func GetPaperAccount(w http.ResponseWriter, r *http.Request) {
	account := utils.GetPaperAccount()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(account)
}

func GetLiveAccount(w http.ResponseWriter, r *http.Request) {
	account := utils.GetLiveAccount()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(account)
}
