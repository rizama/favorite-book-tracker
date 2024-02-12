package controller

import "github.com/gofiber/fiber/v2"

type BookController interface {
	GetBook(ctx *fiber.Ctx) error
	SaveBook(ctx *fiber.Ctx) error

	GetBookHtmx(ctx *fiber.Ctx) error
	SaveBookHtmx(ctx *fiber.Ctx) error
	DeleteBookHtmx(ctx *fiber.Ctx) error
}
