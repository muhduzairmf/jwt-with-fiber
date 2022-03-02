package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/joho/godotenv"
	"github.com/muhduzairmf/jwt-with-fiber/database"
	"github.com/muhduzairmf/jwt-with-fiber/routes"
)

func main() {
	database.ConnectDb()

	app := fiber.New()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Unable to load .env")
		os.Exit(2)
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/home")
	})

	app.Get("/home", routes.HomePage)
	// Public route

	app.Post("/signup", routes.HandleSignup)

	app.Post("/login", routes.HandleLogin)

	app.Use("/profile", jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("ACCESS_TOKEN")),
	}))
	// Middleware for /profile route

	app.Get("/profile/:id", routes.ShowProfile)
	// Protected route

	app.Put("/profile/:id", routes.EditProfile)
	// Protected route
	
	log.Println(app.Listen(":3440"))
}