package routes

import "github.com/gofiber/fiber/v2"

func HomePage(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"status": "Ok",
		"message": "Welcome to home page",
	})
}