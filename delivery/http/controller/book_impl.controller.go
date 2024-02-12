package controller

import (
	"fmt"
	"strconv"

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

func (controller *BookControllerImpl) GetBookHtmx(ctx *fiber.Ctx) error {
	// panggil book usecase untuk mengambil data
	book, err := controller.domain.BookUsecase.GetBook()

	if err != nil {
		resp, statuCode := utils.ConstructorResponseError(fiber.StatusBadRequest, "Failed to fetch book")
		return ctx.Status(statuCode).JSON(resp)
	}

	var result string
	for _, b := range book {
		result += fmt.Sprintf("<tr id='items-%d'><td>%s</td><td>%s</td><td>%d</td><td><button hx-delete='http://localhost:3000/htmx/%d' hx-trigger='click' hx-swap='outerHTML' hx-target='#items-%d' hx-confirm='apakah yakin menghapus data ini?'>Delete</button></td></tr>",
			b.Id, b.Title, b.Author, b.Rating, b.Id, b.Id)
	}

	return ctx.SendString(result)
}

func (controller *BookControllerImpl) SaveBookHtmx(ctx *fiber.Ctx) error {

	// definisikan variable book
	var book request.RequestBookDTO

	// parsing dto dari body dan masukan ke variable book
	if err := ctx.BodyParser(&book); err != nil {

		// jika payload tidak sesuai dto maka balikan error
		resp, statusCode := utils.ConstructorResponseError(fiber.StatusBadRequest, "Invalid request body")
		return ctx.Status(statusCode).JSON(resp)
	}

	if book.Title == "" || book.Author == "" {
		return ctx.SendString("Tolong Masukan Title dan Author")
	}

	// transform dari format dto ke entity book
	bookEntity := book.ToBookEntity()

	// send data ke book usecase untuk simpan data
	if err := controller.domain.BookUsecase.SaveBook(bookEntity); err != nil {
		resp, statusCode := utils.ConstructorResponseError(fiber.StatusBadRequest, "Failed to save book")
		return ctx.Status(statusCode).JSON(resp)
	}

	return ctx.SendString("Success Added Book")
}

func (controller *BookControllerImpl) DeleteBookHtmx(ctx *fiber.Ctx) error {
	// ambil id dari params
	Pid := ctx.Params("id")
	fmt.Println(Pid)
	id, err := strconv.Atoi(Pid)
	if err != nil {
		resp, statusCode := utils.ConstructorResponseError(fiber.StatusInternalServerError, "Failed to delete book")
		return ctx.Status(statusCode).JSON(resp)
	}

	if err := controller.domain.BookUsecase.DeleteBook(id); err != nil {
		resp, statusCode := utils.ConstructorResponseError(fiber.StatusBadRequest, "Failed to delete book")
		return ctx.Status(statusCode).JSON(resp)
	}

	return ctx.SendString("")
}
