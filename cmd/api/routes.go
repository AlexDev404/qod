package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (c *serverConfig) routes() http.Handler {
	c.router.NotFound = http.HandlerFunc(c.notFoundResponse)
	c.router.MethodNotAllowed = http.HandlerFunc(c.methodNotAllowedResponse)

	c.router.GET(v1("/healthcheck"), func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		c.HealthCheckHandler(w, r, ps)
	})

	// Quotes
	c.router.POST(v1("/quotes"), c.CreateQuoteHandler)              // C
	c.router.GET(v1("/quotes"), c.GetQuotesHandler)                 // R
	c.router.PUT(v1("/quotes/:id"), c.UpdateQuoteHandler)           // U
	c.router.DELETE(v1("/quotes/:id"), c.DeleteQuoteHandler)        // D

	// Comments
	c.router.POST(v1("/comments"), c.CreateCommentHandler)          // C
	c.router.GET(v1("/comments"), c.GetCommentsHandler)             // R
	c.router.PUT(v1("/comments/:id"), c.UpdateCommentHandler)       // U
	c.router.DELETE(v1("/comments/:id"), c.DeleteCommentHandler)    // D
	return c.RecoverPanic(c.router)
}
