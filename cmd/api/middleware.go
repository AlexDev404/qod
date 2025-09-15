package main

import (
	"net/http"
	"strings"

	"github.com/rs/cors"
)

func (c *serverConfig) middleware(next http.Handler) http.Handler {
	handler := c.RecoverPanic(next)

	// Configure CORS based on environment
	corsOptions := cors.Options{
		AllowedOrigins: c.getAllowedOrigins(),
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
			http.MethodOptions,
		},
		AllowedHeaders: []string{
			"Content-Type",
			"Authorization",
			"X-Requested-With",
			"Accept",
			"Origin",
		},
		ExposedHeaders: []string{
			"Content-Length",
			"Content-Type",
		},
		AllowCredentials: getEnvAsBool("CORS_ALLOW_CREDENTIALS", false),
		MaxAge:           300, // 5 minutes
		Debug:            c.env == "development",
	}

	corsHandler := cors.New(corsOptions)
	handler = corsHandler.Handler(handler)
	return handler
}

// getAllowedOrigins returns the list of allowed origins based on environment
func (c *serverConfig) getAllowedOrigins() []string {
	// Get allowed origins from environment variable
	allowedOrigins := getEnvAsString("CORS_ALLOWED_ORIGINS", "")

	if allowedOrigins != "" {
		// Split comma-separated origins and trim whitespace
		origins := strings.Split(allowedOrigins, ",")
		for i, origin := range origins {
			origins[i] = strings.TrimSpace(origin)
		}
		return origins
	}

	// Default origins based on environment
	if c.env == "development" {
		return []string{
			"http://localhost:3000",
			"http://localhost:3001",
			"http://localhost:8080",
			"http://127.0.0.1:3000",
			"http://127.0.0.1:3001",
			"http://127.0.0.1:8080",
		}
	}

	// Production: require explicit configuration
	return []string{}
}
