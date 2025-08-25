package main

import (
	"fmt"
	"net/http"
	"qotd/cmd/api"
)

func main() {
	port := "8080"
	fmt.Print("Listening on port " + port)
	api.Routes()
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Print("\n")
		fmt.Print(err)
	}
}
