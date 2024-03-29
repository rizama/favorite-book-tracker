package controller

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/rizama/favorite-book-tracker/delivery/http/dto/request"
	"github.com/rizama/favorite-book-tracker/domain"
	"github.com/rizama/favorite-book-tracker/shared/utils"
)

type BookControllerImpl struct {
	domain   domain.Domain // domain disini mengandung usecase + repository
	Validate *validator.Validate
}

// provider atau constructor
func NewBookController(domain domain.Domain, validate *validator.Validate) BookController {
	return &BookControllerImpl{
		domain:   domain,
		Validate: validate,
	}
}

func (v BookControllerImpl) CheckValidation(data interface{}) []utils.ErrorResponseMsg {
	validationErrors := []utils.ErrorResponseMsg{}

	errs := v.Validate.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			var elem utils.ErrorResponseMsg

			elem.FailedField = err.Field()
			elem.Tag = err.Tag()
			elem.Value = err.Value()
			elem.Error = true

			validationErrors = append(validationErrors, elem)
		}
	}

	return validationErrors
}

// Method BookController

// Server Rendering
func (controller *BookControllerImpl) GetBook(ctx *fiber.Ctx) error {
	// panggil book usecase untuk mengambil data
	book, err := controller.domain.BookUsecase.GetBook()

	if err != nil {
		resp, statuCode := utils.ResponseError(fiber.StatusBadRequest, "Failed ti fetch book")
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
		resp, statusCode := utils.ResponseError(fiber.StatusBadRequest, "Invalid request body")
		return ctx.Status(statusCode).JSON(resp)
	}

	// transform dari format dto ke entity book
	bookEntity := book.ToBookEntity()

	// send data ke book usecase untuk simpan data
	if err := controller.domain.BookUsecase.SaveBook(bookEntity); err != nil {
		resp, statusCode := utils.ResponseError(fiber.StatusBadRequest, "Failed to save book")
		return ctx.Status(statusCode).JSON(resp)
	}

	return ctx.Redirect("/")
}

// HTMX
func (controller *BookControllerImpl) GetBookHtmx(ctx *fiber.Ctx) error {
	book, err := controller.domain.BookUsecase.GetBook()

	if err != nil {
		resp, statuCode := utils.ResponseError(fiber.StatusBadRequest, "Failed to fetch book")
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
	var book request.RequestBookDTO

	if err := ctx.BodyParser(&book); err != nil {
		resp, statusCode := utils.ResponseError(fiber.StatusBadRequest, "Invalid request body")
		return ctx.Status(statusCode).JSON(resp)
	}

	if book.Title == "" || book.Author == "" {
		return ctx.SendString("Tolong Masukan Title dan Author")
	}

	bookEntity := book.ToBookEntity()

	if err := controller.domain.BookUsecase.SaveBook(bookEntity); err != nil {
		resp, statusCode := utils.ResponseError(fiber.StatusBadRequest, "Failed to save book")
		return ctx.Status(statusCode).JSON(resp)
	}

	return ctx.SendString("Success Added Book")
}

func (controller *BookControllerImpl) DeleteBookHtmx(ctx *fiber.Ctx) error {
	Pid := ctx.Params("id")

	id, err := strconv.Atoi(Pid)
	if err != nil {
		resp, statusCode := utils.ResponseError(fiber.StatusInternalServerError, "Failed to delete book")
		return ctx.Status(statusCode).JSON(resp)
	}

	if err := controller.domain.BookUsecase.DeleteBook(id); err != nil {
		resp, statusCode := utils.ResponseError(fiber.StatusBadRequest, "Failed to delete book")
		return ctx.Status(statusCode).JSON(resp)
	}

	return ctx.SendString("")
}

// API
func (controller *BookControllerImpl) GetBookApi(ctx *fiber.Ctx) error {
	book, err := controller.domain.BookUsecase.GetBook()

	if err != nil {
		resp, statuCode := utils.ResponseError(fiber.StatusBadRequest, "Failed to fetch book")
		return ctx.Status(statuCode).JSON(resp)
	}

	resp := utils.ResponseSuccess(fiber.StatusOK, "success", book)

	return ctx.JSON(resp)
}

func (controller *BookControllerImpl) GetBookByIdApi(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	idParsed, err := strconv.Atoi(id)
	if err != nil {
		resp, statuCode := utils.ResponseError(fiber.StatusBadRequest, err.Error())
		return ctx.Status(statuCode).JSON(resp)
	}

	book, err2 := controller.domain.BookUsecase.GetBookById(idParsed)
	if err2 != nil {
		resp, statuCode := utils.ResponseError(fiber.StatusBadRequest, err2.Error())
		return ctx.Status(statuCode).JSON(resp)
	}

	resp := utils.ResponseSuccess(fiber.StatusOK, "success", book)

	return ctx.JSON(resp)
}

func (controller *BookControllerImpl) SaveBookApi(ctx *fiber.Ctx) error {
	var book request.RequestBookDTO

	if err := ctx.BodyParser(&book); err != nil {
		resp, statusCode := utils.ResponseError(fiber.StatusBadRequest, "Invalid request body")
		return ctx.Status(statusCode).JSON(resp)
	}

	// Validation
	if errs := controller.CheckValidation(book); len(errs) > 0 && errs[0].Error {
		errMsgs := utils.ErrValidationMsg(errs)
		resp, statusCode := utils.ResponseError(fiber.StatusBadRequest, strings.Join(errMsgs, " | "))
		return ctx.Status(statusCode).JSON(resp)
	}

	bookEntity := book.ToBookEntity()

	if err := controller.domain.BookUsecase.SaveBook(bookEntity); err != nil {
		resp, statusCode := utils.ResponseError(fiber.StatusBadRequest, "Failed to save book")
		return ctx.Status(statusCode).JSON(resp)
	}

	resp := utils.ResponseSuccess(fiber.StatusCreated, "success", book)
	return ctx.Status(fiber.StatusCreated).JSON(resp)
}

func (controller *BookControllerImpl) UpdateBookApi(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	idParsed, err := strconv.Atoi(id)
	if err != nil {
		resp, statuCode := utils.ResponseError(fiber.StatusBadRequest, err.Error())
		return ctx.Status(statuCode).JSON(resp)
	}

	var book request.RequestBookDTO

	if err := ctx.BodyParser(&book); err != nil {
		resp, statusCode := utils.ResponseError(fiber.StatusBadRequest, "Invalid request body")
		return ctx.Status(statusCode).JSON(resp)
	}

	// Validation
	if errs := controller.CheckValidation(book); len(errs) > 0 && errs[0].Error {
		errMsgs := utils.ErrValidationMsg(errs)
		resp, statusCode := utils.ResponseError(fiber.StatusBadRequest, strings.Join(errMsgs, " | "))
		return ctx.Status(statusCode).JSON(resp)
	}

	err2 := controller.domain.BookUsecase.UpdateBook(book, idParsed)
	if err2 != nil {
		resp, statusCode := utils.ResponseError(fiber.StatusBadRequest, err2.Error())
		return ctx.Status(statusCode).JSON(resp)
	}

	resp := utils.ResponseSuccess(fiber.StatusOK, "success", book)
	return ctx.Status(fiber.StatusCreated).JSON(resp)
}
