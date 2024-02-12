package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizama/favorite-book-tracker/delivery/http/controller"
	"github.com/rizama/favorite-book-tracker/domain"
)

func NewRouter(app *fiber.App, domain domain.Domain) {
	bookController := controller.NewBookController(domain)

	app.Get("/", bookController.GetBook)
	app.Post("/", bookController.SaveBook)

	app.Get("/htmx", bookController.GetBookHtmx)
	app.Post("/htmx", bookController.SaveBookHtmx)
	app.Delete("/htmx/:id", bookController.DeleteBookHtmx)
}
