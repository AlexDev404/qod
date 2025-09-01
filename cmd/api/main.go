package main

import (
	"fmt"
	"net/http"
	"os"
	"qotd/cmd/api/database"
)

type serverConfig struct {
	port int
	env  string
	db   *database.Database
}

func main() {
	config := serverConfig{
		port: 8080,
		env:  "development",
		db:   database.NewDatabase(database.InMemory, nil),
	}

	fmt.Print("Listening on port " + fmt.Sprint(config.port))
	config.routes()
	err := http.ListenAndServe(":"+fmt.Sprint(config.port), nil)
	if err != nil {
		fmt.Print("\n")
		fmt.Print(err)
	}
	os.Exit(1)
}
