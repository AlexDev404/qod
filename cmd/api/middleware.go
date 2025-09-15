package main

import (
	"net/http"

	"github.com/rs/cors"
)

func (c *serverConfig) middleware(next http.Handler) http.Handler {
	handler := c.RecoverPanic(next)

	// Add CORS
	newCorsWithOptions := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://foo.com", "http://foo.com:8080"},
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	})

	handler = newCorsWithOptions.Handler(handler)
	return handler
}
