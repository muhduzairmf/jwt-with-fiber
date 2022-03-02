package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhduzairmf/jwt-with-fiber/database"
	"github.com/muhduzairmf/jwt-with-fiber/models"
)

func ShowProfile(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "Bad request",
			"message": "The id params must be an integer",
			"error": err.Error(),
		})
	}

	if id < 0 {
		return c.Status(400).JSON(fiber.Map{
			"status": "Bad request",
			"message": "The id params must be a positive integer",
			"error": "The id params has value less than zero (0).",
		})
	}

	var userProfile models.Profile

	database.Database.Db.Find(&userProfile, "id = ?", id)
	if userProfile.ID == 0 {
		return c.Status(400).JSON(fiber.Map{
			"status": "Bad request",
			"message": "The user profile does not exist",
			"error": "Cannot find user with the given id",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status": "Ok",
		"message": "Your profile",
		"data": userProfile,
	})
}

func EditProfile(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "Bad request",
			"message": "The id params must be an integer",
			"error": err.Error(),
		})
	}

	if id < 0 {
		return c.Status(400).JSON(fiber.Map{
			"status": "Bad request",
			"message": "The id params must be a positive integer",
			"error": "The id params has value less than zero (0).",
		})
	}

	var userProfile models.User

	database.Database.Db.Find(&userProfile, "id = ?", id)
	if userProfile.ID == 0 {
		return c.Status(400).JSON(fiber.Map{
			"status": "Bad request",
			"message": "The user profile does not exist",
			"error": "Cannot find user with the given id",
		})
	}

	type profile struct {
		Email string `json:"email"`
		FirstName string `json:"firstName"`
		LastName string `json:"lastName"`
		Birthday string `json:"birthday"`
		Website string `json:"website"`
	}

	var toEdit profile

	err = c.BodyParser(&toEdit)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "Bad request",
			"message": "Error occured when parsing the request body",
			"error": err.Error(),
		})
	}

	var updatedProfile models.Profile

	updatedProfile.Email = toEdit.Email
	updatedProfile.FirstName = toEdit.FirstName
	updatedProfile.LastName = toEdit.LastName
	updatedProfile.Birthday = toEdit.Birthday
	updatedProfile.Website = toEdit.Website

	database.Database.Db.Save(&updatedProfile)

	return c.Status(200).JSON(fiber.Map{
		"status": "Ok",
		"message": "Profile successfully edited!",
		"data": updatedProfile,
	})
}