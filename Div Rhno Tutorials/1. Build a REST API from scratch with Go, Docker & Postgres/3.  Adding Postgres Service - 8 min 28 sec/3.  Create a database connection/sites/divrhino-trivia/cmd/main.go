package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/divrhino/divrhino-trivia/database"
)

func main() {
	database.ConnectDb()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Nols Smit!")
	})

	app.Listen(":3000")
}