package main

import (
	"log"
	"net/http"
	"web3/package/config"
	handlers "web3/package/handlers"
)

func main() {
	var app config.AppConfig

	repo := handlers.NewRepo(&app)  //pass in reference to our config
	handlers.NewHandlers(repo)

	srv := &http.Server{
		Addr: ":8080",
		Handler: routes(&app),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

// To run: $ go run ./cmd/web