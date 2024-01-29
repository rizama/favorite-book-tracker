package controller

import "github.com/gofiber/fiber/v2"

type BookController interface {
	GetBook(ctx *fiber.Ctx) error
	SaveBook(ctx *fiber.Ctx) error
}
