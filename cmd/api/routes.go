package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (c *serverConfig) routes() {
	router := httprouter.New()
	router.GET(v1("/healthcheck"), func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		HealthCheckHandler(w, r, ps)
	})
	http.ListenAndServe(":"+fmt.Sprint(c.port), router)
}
