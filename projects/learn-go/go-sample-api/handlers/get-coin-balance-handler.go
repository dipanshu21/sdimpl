package handlers

import (
	"api/api"
	"api/internal/tools"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"
)

func GetCoinBalance(w http.ResponseWriter, r *http.Request) {
	params := api.CoinBalanceParams{}
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error

	err = decoder.Decode(&params, r.URL.Query())
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var database tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil {
		log.Error("Failed to connect to the database:", err)
		api.InternalErrorHandler(w)
		return
	}

	coinDetails := database.GetUserCoins(params.Username)
	if coinDetails == nil {
		log.Error("User not found:", params.Username)
		api.RequestErrorHandler(w, errors.New("user not found"))
		return
	}

	response := api.CoinBalanceResponse{
		Code:    http.StatusOK,
		Balance: coinDetails.Coins,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Error("Failed to encode response:", err)
		api.InternalErrorHandler(w)
		return
	}
}
