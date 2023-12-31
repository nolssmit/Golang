package handlers

import (
	"github.com/divrhino/divrhino-trivia/database"
	"github.com/divrhino/divrhino-trivia/models"
	"github.com/gofiber/fiber/v2"
)

func ListFacts(c *fiber.Ctx) error {
	facts := []models.Fact{}
	database.DB.Db.Find(&facts)

	return c.Render("index", fiber.Map{
		"Title": "Div Rhino Trivia Time",
		"Subtitle": "Facts for funtimes with friends!",
		"Facts":    facts,
	})
}

func NewFactView(c *fiber.Ctx) error {
	return c.Render("new", fiber.Map{
		"Title":    "New Fact",
		"Subtitle": "Add a cool fact!",
	})
}

func CreateFact(c *fiber.Ctx) error {
    fact := new(models.Fact)
    // Parse request body
    if err := c.BodyParser(fact); err != nil {
        return NewFactView(c)
    }

    // Create fact in database
    result := database.DB.Db.Create(&fact)
    if result.Error != nil {
        return NewFactView(c)
    }

    return ListFacts(c)
}

func ShowFact(c *fiber.Ctx) error {
    fact := models.Fact{}
    id := c.Params("id")

    database.DB.Db.Where("id = ?", id).First(&fact)

	return c.Render("show", fiber.Map {
		"Title": "Single Fact",
		"Fact": fact,
	})
}