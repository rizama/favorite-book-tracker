package router

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/rizama/favorite-book-tracker/delivery/http/controller"
	"github.com/rizama/favorite-book-tracker/delivery/http/middleware"
	"github.com/rizama/favorite-book-tracker/domain"
)

func BookRouter(app *fiber.App, domain domain.Domain) {
	validate := validator.New()
	bookController := controller.NewBookController(domain, validate)
	htmx := app.Group("/htmx")                          // group /htmx
	api := app.Group("/api", middleware.AuthMiddleware) // group /api

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
