package controller

import "github.com/gofiber/fiber/v2"

type BookController interface {
	// render server side
	GetBook(ctx *fiber.Ctx) error
	SaveBook(ctx *fiber.Ctx) error

	// htmx
	GetBookHtmx(ctx *fiber.Ctx) error
	SaveBookHtmx(ctx *fiber.Ctx) error
	DeleteBookHtmx(ctx *fiber.Ctx) error

	// restful api
	GetBookApi(ctx *fiber.Ctx) error
	GetBookByIdApi(ctx *fiber.Ctx) error
	SaveBookApi(ctx *fiber.Ctx) error
	UpdateBookApi(ctx *fiber.Ctx) error
}
