package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/muhduzairmf/jwt-with-fiber/database"
	"github.com/muhduzairmf/jwt-with-fiber/routes"
)

func main() {
	database.ConnectDb()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/home")
	})

	app.Get("/home", routes.HomePage)
	// Public route

	app.Post("/signup", routes.HandleSignup)

	app.Post("/login", routes.HandleLogin)

	app.Get("/profile", routes.ShowProfile)
	// Protected route

	app.Put("/profile", routes.EditProfile)
	// Protected route
	
	log.Println(app.Listen(":3440"))
}