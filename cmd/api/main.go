package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"qotd/cmd/api/database"

	"github.com/julienschmidt/httprouter"
)

type serverConfig struct {
	port    int
	env     string
	logger  *slog.Logger
	version string
	db      *database.Database
	router  *httprouter.Router
}

func main() {
	config := serverConfig{
		port:    8080,
		env:     "development",
		version: "1.0.0",
		db:      database.NewDatabase(database.InMemory, nil),
		router:  httprouter.New(),
	}

	// Create a logger
	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
	config.logger = logger

	fmt.Print("Listening on port " + fmt.Sprint(config.port))
	config.routes()
	err := http.ListenAndServe(":"+fmt.Sprint(config.port), config.router)
	if err != nil {
		fmt.Print("\n")
		fmt.Print(err)
	}
	os.Exit(1)
}
