package main

import (
	"os"
	"strconv"
	"strings"
)

func v(str string) string {
	return "/" + getEnvAsString("API_VERSION", "v1") + str
}

func getEnvAsString(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

func getEnvAsBool(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		// Convert to lowercase for case-insensitive comparison
		lowerValue := strings.ToLower(value)
		if lowerValue == "true" || lowerValue == "1" || lowerValue == "yes" {
			return true
		} else if lowerValue == "false" || lowerValue == "0" || lowerValue == "no" {
			return false
		}
	}
	return defaultValue
}
