package main

import (
	"github.com/divrhino/divrhino-trivia/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.ListFacts)

	app.Get("/fact", handlers.NewFactView) // Add new route for new view
	app.Post("/fact", handlers.CreateFact)

    // Add new route to show single Fact, given `:id`
    app.Get("/fact/:id", handlers.ShowFact)	
}
