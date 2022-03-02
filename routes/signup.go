package routes

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"github.com/muhduzairmf/jwt-with-fiber/database"
	"github.com/muhduzairmf/jwt-with-fiber/models"
)

func HandleSignup(c *fiber.Ctx) error {
	type user struct {
		Email string `json:"email"`
		Password string `json:"password"`
		ConfirmPassword string `json:"confirmPassword"`
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

	log.Println(newUser)

	if newUser.Email == "" || newUser.Password == "" || newUser.ConfirmPassword == "" {
		return c.Status(400).JSON(fiber.Map{
			"status": "Bad request",
			"message": "Important field is empthy. Please include email (string), password (string) and confirmPassword (string)",
		})
	}

	if newUser.Password != newUser.ConfirmPassword {
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
	hashedPassword := fmt.Sprintf("%x", sha256.Sum256([]byte(newUser.Password+salt)))

	// Bring the salt with the hashed password
	hashWithSalt := fmt.Sprintf("%v:%v", salt, hashedPassword)

	var verifiedUser models.User

	verifiedUser.Email = newUser.Email
	verifiedUser.Password = hashWithSalt

	result := database.Database.Db.Create(&verifiedUser)
	if result.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "Bad request",
			"message": "Error occured when creating data",
			"error": result.Error,
		})
	}

	var userProfile models.Profile

	userProfile.Email = verifiedUser.Email
	userProfile.FirstName = ""
	userProfile.LastName = ""
	userProfile.Birthday = ""
	userProfile.Website = ""
	userProfile.UserId = verifiedUser.ID

	result = database.Database.Db.Create(&userProfile)
	if result.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "Bad request",
			"message": "Error occured when creating data",
			"error": result.Error,
		})
	}

	expires := time.Now().Add(time.Hour * 24).Unix()
	// Make the expiry time is 24 hours

	generateToken := jwt.New(jwt.SigningMethodHS256)
	claims := generateToken.Claims.(jwt.MapClaims)

	claims["id"] = verifiedUser.ID
	claims["email"] = verifiedUser.Email
	claims["password"] = verifiedUser.Password
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

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "Server error",
			"message": "Error occured when generating token. Consider to login",
			"error": err.Error(),
		})
	}

	var response = struct{
		Email string `json:"email"`
		JWT_token string `json:"token"`
	}{
		Email: newUser.Email,
		JWT_token: theToken,
	}

	return c.Status(200).JSON(fiber.Map{
		"status": "Ok",
		"message": "Successfully signed up",
		"data": response,
	})
}