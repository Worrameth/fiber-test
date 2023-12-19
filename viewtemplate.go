package main

import (
	"github.com/gofiber/fiber/v2"
)

func testHTML(c *fiber.Ctx) error {
	// Render the template with variable data
	return c.Render("index", fiber.Map{
		"Title": "Hello World!",
		"Name":  "Worrameth",
	})
}
