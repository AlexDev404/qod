package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
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

func (a *serverConfig) readRequestJSON(w http.ResponseWriter, r *http.Request, destination any) error {
	err := json.NewDecoder(r.Body).Decode(destination)
	if err != nil {
		// check for the different errors
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var invalidUnmarshalError *json.InvalidUnmarshalError

		switch {
		case errors.As(err, &syntaxError):
			return fmt.Errorf("The body contains badly-formed JSON (at character %d)", syntaxError.Offset)
			// Decode can also send back an io error message
		case errors.Is(err, io.ErrUnexpectedEOF):
			return errors.New("The body contains badly-formed JSON")

		case errors.As(err, &unmarshalTypeError):
			if unmarshalTypeError.Field != "" {
				return fmt.Errorf("The body contains the incorrect JSON type for field %q",
					unmarshalTypeError.Field)
			}
			return fmt.Errorf("The body contains the incorrect JSON type (at character %d)",
				unmarshalTypeError.Offset)
		case errors.Is(err, io.EOF):
			return errors.New("The body must not be empty")

		// the programmer messed up
		case errors.As(err, &invalidUnmarshalError):
			panic(err)
			// some other type of error
		default:
			return err
		}
	}

	return nil
}
