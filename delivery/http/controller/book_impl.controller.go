package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rizama/favorite-book-tracker/delivery/http/dto/request"
	"github.com/rizama/favorite-book-tracker/domain"
	"github.com/rizama/favorite-book-tracker/shared/utils"
)

type BookControllerImpl struct {
	domain domain.Domain // domain disini mengandung usecase + repository
}

// provider atau constructor
func NewBookController(domain domain.Domain) BookController {
	return &BookControllerImpl{
		domain: domain,
	}
}

// Method BookController
func (controller *BookControllerImpl) GetBook(ctx *fiber.Ctx) error {
	// panggil book usecase untuk mengambil data
	book, err := controller.domain.BookUsecase.GetBook()

	if err != nil {
		resp, statuCode := utils.ConstructorResponseError(fiber.StatusBadRequest, "Failed ti fetch book")
		return ctx.Status(statuCode).JSON(resp)
	}

	return ctx.Render("resource/views/home", fiber.Map{
		"Books": book,
	})
}

func (controller *BookControllerImpl) SaveBook(ctx *fiber.Ctx) error {

	// definisikan variable book
	var book request.RequestBookDTO

	// parsing dto dari body dan masukan ke variable book
	if err := ctx.BodyParser(&book); err != nil {

		// jika payload tidak sesuai dto maka balikan error
		resp, statusCode := utils.ConstructorResponseError(fiber.StatusBadRequest, "Invalid request body")
		return ctx.Status(statusCode).JSON(resp)
	}

	// transform dari format dto ke entity book
	bookEntity := book.ToBookEntity()

	// send data ke book usecase untuk simpan data
	if err := controller.domain.BookUsecase.SaveBook(bookEntity); err != nil {
		resp, statusCode := utils.ConstructorResponseError(fiber.StatusBadRequest, "Failed to save book")
		return ctx.Status(statusCode).JSON(resp)
	}

	return ctx.Redirect("/")
}
