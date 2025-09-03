package main

import (
	"encoding/json"
	"net/http"
	"qotd/cmd/api/types"

	"github.com/julienschmidt/httprouter"
)

type HealthCheck struct {
	Status      string `json:"status"`
	Environment string `json:"environment,omitempty"`
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(HealthCheck{Status: "OK"})
}

func (c *serverConfig) CreateQuoteHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var quote types.Quote
	if err := json.NewDecoder(r.Body).Decode(&quote); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if quote.Author == "" {
		http.Error(w, "Field 'Author' missing", http.StatusBadRequest)
		return
	}

	if quote.Text == "" {
		http.Error(w, "Field 'Text' missing", http.StatusBadRequest)
		return
	}

	if err := c.db.WriteQuote(quote); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(quote)
}

func (c *serverConfig) GetQuotesHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	quotes, err := c.db.GetQuotes()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(quotes) == 0 {
		quotes = []types.Quote{}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(quotes)
}
