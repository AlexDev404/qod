package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (c *serverConfig) routes() {
	c.router.NotFound = http.HandlerFunc(c.notFoundResponse)
	c.router.MethodNotAllowed = http.HandlerFunc(c.methodNotAllowedResponse)

	c.router.GET(v1("/healthcheck"), func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		c.HealthCheckHandler(w, r, ps)
	})
	c.router.POST(v1("/quotes"), c.CreateQuoteHandler)
	c.router.GET(v1("/quotes"), c.GetQuotesHandler)
	c.router.POST(v1("/comments"), c.CreateCommentHandler)

}
