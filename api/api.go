package api

import (
	"encoding/json"
	"net/http"
)

// Coin Balance Parameters
type CoinBalanceParams struct {
	Username string
}

// Coin Balance Response
type CoinBalanceResponse struct {
	Code int

	Balance int64
}

type Error struct {
	Code    int
	Message string
}

func writeError(w http.ResponseWriter, message string, code int) {
	resp := Error{
		Code:    code,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusInternalServerError)
	}

	InternalErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "Invalid Request", http.StatusInternalServerError)
	}
)
