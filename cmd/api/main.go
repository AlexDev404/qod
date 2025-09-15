package main

import (
	"flag"
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
		router:  httprouter.New(),
	}

	dbDsn := os.Getenv("DB_DSN")
	dbType := os.Getenv("DB_TYPE")

	// Read in the database type
	flag.StringVar(&dbType, "db-type", dbType, "Database type (IN_MEMORY or POSTGRES)")

	if dbType == "POSTGRES" {
		flag.StringVar(&dbDsn, "db-dsn", dbDsn, "PostgreSQL DSN")
	}

	fmt.Println(dbDsn)
	fmt.Println(dbType)

	flag.IntVar(&config.port, "port", 8080, "API server port")

	if dbType != "IN_MEMORY" && dbType != "POSTGRES" {
		fmt.Print("Error: Unsupported database type. Use IN_MEMORY or POSTGRES.\n")
		os.Exit(1)
	}

	if dbType == "POSTGRES" && dbDsn == "" {
		fmt.Print("Error: DSN must be provided for PostgreSQL database type.\n")
		os.Exit(1)
	}

	flag.Parse()
	// Initialize the database connection
	if dbType == "POSTGRES" {
		config.db = database.NewDatabase(database.Postgres, &dbDsn)
	} else {
		config.db = database.NewDatabase(database.InMemory, nil)
	}

	// Create a logger
	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
	config.logger = logger

	fmt.Println("Listening on port " + fmt.Sprint(config.port))
	router := config.routes()
	config.db.Connect()
	err := http.ListenAndServe(":"+fmt.Sprint(config.port), router)
	// release the database resources before exiting
	defer config.db.Disconnect()
	if err != nil {
		fmt.Print("\n")
		fmt.Print(err)
	}
	os.Exit(1)
}
