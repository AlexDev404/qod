package main

import (
	"encoding/json"
	"net/http"
)

type envelope map[string]any

func (a *serverConfig) writeResponseJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	jsonResponse, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	jsonResponse = append(jsonResponse, '\n')
	// additional headers to be set
	for key, value := range headers {
		w.Header()[key] = value
	}
	// set content type header
	w.Header().Set("Content-Type", "application/json")
	// explicitly set the response status code
	w.WriteHeader(status)
	_, err = w.Write(jsonResponse)
	if err != nil {
		return err
	}

	return nil

}
