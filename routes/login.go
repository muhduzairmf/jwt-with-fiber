package routes

import (
	"crypto/sha256"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/muhduzairmf/jwt-with-fiber/database"
	"github.com/muhduzairmf/jwt-with-fiber/models"
)

func HandleLogin(c *fiber.Ctx) error {
	type user struct {
		email string
		password string
	}

	var currentUser user

	err := c.BodyParser(&currentUser)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "Bad request",
			"message": "Error occured when parsing the request body",
			"error": err.Error(),
		})
	}

	if currentUser.email == "" || currentUser.password == "" {
		return c.Status(400).JSON(fiber.Map{
			"status": "Bad request",
			"message": "Important field is empthy. Please include email (string) and password (string)",
		})
	}

	var theUser models.User

	database.Database.Db.Find(&theUser, "email = ?", currentUser.email)
	if theUser.ID == 0 {
		return c.Status(400).JSON(fiber.Map{
			"status": "Bad request",
			"message": "The user does not exist",
			"error": "Cannot find user with the given id",
		})
	}

	saltAndHash := strings.Split(theUser.Password, ":")
	salt := saltAndHash[0]
	hashedPasswordSaved := saltAndHash[1]

	hashedPasswordInput := fmt.Sprintf("%x", sha256.Sum256([]byte(currentUser.password+salt)))

	if hashedPasswordSaved != hashedPasswordInput {
		return c.Status(400).JSON(fiber.Map{
			"status": "Bad request",
			"message": "Invalid password",
		})
	}

	var response = struct{
		email string
		jwt_token string
	}{
		email: theUser.Email,
		jwt_token: "",
	}

	return c.Status(200).JSON(fiber.Map{
		"status": "Ok",
		"message": "Successfully signed up",
		"data": response,
	})
}