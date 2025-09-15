package main

import (
	"net/http"

	"github.com/rs/cors"
)

func (c *serverConfig) middleware(next http.Handler) http.Handler {
	handler := c.RecoverPanic(c.router)
	handler = cors.Default().Handler(handler)
	return handler
}
