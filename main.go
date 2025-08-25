package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HealthCheck struct {
	Status      string `json:"status"`
	Environment string `json:"environment,omitempty"`
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(HealthCheck{Status: "OK"})
}

func main() {
	port := "8080"
	fmt.Print("Listening on port " + port)
	http.HandleFunc("/healthcheck", healthCheckHandler)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Print("\n")
		fmt.Print(err)
	}
}
