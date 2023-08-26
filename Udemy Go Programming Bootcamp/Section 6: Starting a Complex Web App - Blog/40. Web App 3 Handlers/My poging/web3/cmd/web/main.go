package main

import (
	"log"
	"net/http"
	handlers "web3/package/handlers"
)

func main() {
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/about", handlers.AboutHandler)

	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}

// To run: $ go run ./cmd/web