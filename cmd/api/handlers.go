package api

import (
	"net/http"
)

func Routes() {
	http.HandleFunc(v1("/healthcheck"), HealthCheckHandler)
}
