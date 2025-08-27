package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Routes() {
	router := httprouter.New()
	router.GET(v1("/healthcheck"), func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		HealthCheckHandler(w, r, ps)
	})
	http.ListenAndServe(":8080", router)
}
