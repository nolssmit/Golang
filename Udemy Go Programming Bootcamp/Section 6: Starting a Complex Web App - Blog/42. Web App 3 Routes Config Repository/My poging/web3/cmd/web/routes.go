package main

import (
	"net/http"
	"web3/package/config"
	"web3/package/handlers"

	"github.com/go-chi/chi/v5"
)

func routes(app *config.AppConfig) http.Handler {
	// Mux (HTTP request multiplexer)
	mux := chi.NewRouter()
	mux.Get("/", handlers.Repo.HomeHandler)
	mux.Get("/about", handlers.Repo.AboutHandler)
	return mux
}