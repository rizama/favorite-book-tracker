package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/rizama/favorite-book-tracker/database"
	"github.com/rizama/favorite-book-tracker/routes"
)

type Response struct {
	Message string `json:"message"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.ConnectDB()
	app := fiber.New()

	setRoutes(app)

	app.Use(cors.New())

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	log.Fatal(app.Listen(":3000"))
}

func setRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		response := Response{
			Message: "Hello, World!",
		}
		return c.JSON(response)
	})

	app.Post("/addbook", routes.AddBook)
}
