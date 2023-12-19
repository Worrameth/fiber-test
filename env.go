package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func getEnv(c *fiber.Ctx) error {
	secret := os.Getenv("SECRET")

	if secret == "" {
		secret = "default secret"
	}

	return c.JSON(fiber.Map{
		"SECRET": secret,
	})
}
