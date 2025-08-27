package api

import (
	"encoding/json"
	"net/http"
)

type HealthCheck struct {
	Status      string `json:"status"`
	Environment string `json:"environment,omitempty"`
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(HealthCheck{Status: "OK"})
}
