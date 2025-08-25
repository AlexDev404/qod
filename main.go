package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"net/http"
)

type Quote struct {
	Text   string `json:"text"`
	Author string `json:"author"`
}

var quotes = []Quote{
	{"The best way to get started is to quit talking and begin doing.", "Walt Disney"},
	{"Don't let yesterday take up too much of today.", "Will Rogers"},
	{"It's not whether you get knocked down, it's whether you get up.", "Vince Lombardi"},
	{"If you are working on something exciting, it will keep you motivated.", "Steve Jobs"},
	{"Success is not in what you have, but who you are.", "Bo Bennett"},
}

func quoteHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	b := make([]byte, 1)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	idx := int(b[0]) % len(quotes)
	json.NewEncoder(w).Encode(quotes[idx])
}

func main() {
	port := "8080"
	fmt.Print("Listening on port " + port)
	http.HandleFunc("/quote", quoteHandler)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Print("\n")
		fmt.Print(err)
	}
}
