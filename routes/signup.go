package routes

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/muhduzairmf/jwt-with-fiber/database"
	"github.com/muhduzairmf/jwt-with-fiber/models"
)

func HandleSignup(c *fiber.Ctx) error {
	type user struct {
		email string
		password string
		confirmPassword string
	}

	var newUser user

	err := c.BodyParser(&newUser)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "Bad request",
			"message": "Error occured when parsing the request body",
			"error": err.Error(),
		})
	}

	if newUser.email == "" || newUser.password == "" || newUser.confirmPassword == "" {
		return c.Status(400).JSON(fiber.Map{
			"status": "Bad request",
			"message": "Important field is empthy. Please include email (string), password (string) and confirmPassword (string)",
		})
	}

	if newUser.password != newUser.confirmPassword {
		return c.Status(400).JSON(fiber.Map{
			"status": "Bad request",
			"message": "The password field and confirmPassword field is not match!",
		})
	}

	// Get random number
	randomNum, err := rand.Prime(rand.Reader, 64)
	if err != nil {
		log.Fatal(err)
		return c.Status(500).JSON(fiber.Map{
			"status": "Server error",
			"message": "Error occured when processing password",
			"error": err.Error(),
		})
	}

	// Create salt by converting randomNum to hex
	salt := fmt.Sprintf("%x", randomNum)

	// Hash password with the salt
	hashedPassword := fmt.Sprintf("%x", sha256.Sum256([]byte(newUser.password+salt)))

	// Bring the salt with the hashed password
	hashWithSalt := fmt.Sprintf("%v:%v", salt, hashedPassword)

	var verifiedUser models.User

	verifiedUser.Email = newUser.email
	verifiedUser.Password = hashWithSalt

	result := database.Database.Db.Create(&verifiedUser)
	if result.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "Bad request",
			"message": "Error occured when creating data",
			"error": result.Error,
		})
	}

	var response = struct{
		email string
		jwt_token string
	}{
		email: newUser.email,
		jwt_token: "",
	}

	return c.Status(200).JSON(fiber.Map{
		"status": "Ok",
		"message": "Successfully signed up",
		"data": response,
	})
}