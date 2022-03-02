package routes

import (
	"crypto/sha256"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"github.com/muhduzairmf/jwt-with-fiber/database"
	"github.com/muhduzairmf/jwt-with-fiber/models"
)

func HandleLogin(c *fiber.Ctx) error {
	type user struct {
		Email string `json:"email"`
		Password string `json:"password"`
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

	if currentUser.Email == "" || currentUser.Password == "" {
		return c.Status(400).JSON(fiber.Map{
			"status": "Bad request",
			"message": "Important field is empthy. Please include email (string) and password (string)",
		})
	}

	var theUser models.User

	database.Database.Db.Find(&theUser, "email = ?", currentUser.Email)
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

	hashedPasswordInput := fmt.Sprintf("%x", sha256.Sum256([]byte(currentUser.Password+salt)))

	if hashedPasswordSaved != hashedPasswordInput {
		return c.Status(400).JSON(fiber.Map{
			"status": "Bad request",
			"message": "Invalid password",
		})
	}

	expires := time.Now().Add(time.Hour * 24).Unix()
	// Make the expiry time is 24 hours

	generateToken := jwt.New(jwt.SigningMethodHS256)
	claims := generateToken.Claims.(jwt.MapClaims)

	claims["id"] = theUser.ID
	claims["email"] = theUser.Email
	claims["password"] = theUser.Password
	claims["expires"] = expires

	err = godotenv.Load()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "Server error",
			"message": "Error occured when generating token. Consider to login",
			"error": err.Error(),
		})
	}

	theToken, err := generateToken.SignedString([]byte(os.Getenv("ACCESS_TOKEN")))
	// The "7061c968eadfba2959d6e12d156590bafc" is like a unique access token
	// It should be saved in .env file, because exposing the access token can make user's password vulnerable

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "Server error",
			"message": "Error occured when processing password",
			"error": err.Error(),
		})
	}

	var response = struct{
		Email string `json:"email"`
		JWT_token string `json:"token"`
	}{
		Email: theUser.Email,
		JWT_token: theToken,
	}

	return c.Status(200).JSON(fiber.Map{
		"status": "Ok",
		"message": "Successfully logged in",
		"data": response,
	})
}