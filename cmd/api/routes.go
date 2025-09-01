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
	router.POST(v1("/quotes"), c.CreateQuoteHandler)
	http.ListenAndServe(":"+fmt.Sprint(c.port), router)
}
