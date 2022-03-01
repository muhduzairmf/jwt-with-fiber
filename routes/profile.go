package routes

import "github.com/gofiber/fiber/v2"

func ShowProfile(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"status": "Ok",
		"message": "Your profile",
	})
}

func EditProfile(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"status": "Ok",
		"message": "Profile successfully edited!",
	})
}