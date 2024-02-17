package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizama/favorite-book-tracker/delivery/http/controller"
	"github.com/rizama/favorite-book-tracker/domain"
)

func NewRouter(app *fiber.App, domain domain.Domain) {
	bookController := controller.NewBookController(domain)

	api := app.Group("/api")   // group /api
	htmx := app.Group("/htmx") // group /api

	app.Get("/", bookController.GetBook)
	app.Post("/", bookController.SaveBook)

	htmx.Get("/", bookController.GetBookHtmx)
	htmx.Post("/", bookController.SaveBookHtmx)
	htmx.Delete("/:id", bookController.DeleteBookHtmx)

	api.Get("/", bookController.GetBookApi)
	api.Get("/:id", bookController.GetBookByIdApi)
	api.Post("/", bookController.SaveBookApi)
	api.Put("/:id", bookController.UpdateBookApi)
}
