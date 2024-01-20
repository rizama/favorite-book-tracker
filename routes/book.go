package routes

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/rizama/favorite-book-tracker/database"
	"github.com/rizama/favorite-book-tracker/models"
)

func AddBook(c *fiber.Ctx) error {
	book := new(models.Book)
	err := c.BodyParser(book)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	fmt.Println(book)
	database.DB.DB.Create(&book)

	return c.Status(http.StatusCreated).JSON(book)
}
