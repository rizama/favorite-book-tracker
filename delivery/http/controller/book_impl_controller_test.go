package controller

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/rizama/favorite-book-tracker/delivery/http/dto/request"
	"github.com/rizama/favorite-book-tracker/domain"
	"github.com/rizama/favorite-book-tracker/domain/entity"
	"github.com/stretchr/testify/mock"
)

// MockBookUsecase adalah sebuah mock implementation dari BookUsecase.
type MockBookUsecase struct {
	mock.Mock
}

func (m *MockBookUsecase) GetBook() ([]entity.Book, error) {
	args := m.Called()
	return args.Get(0).([]entity.Book), args.Error(1)
}

func (m *MockBookUsecase) SaveBook(book entity.Book) error {
	args := m.Called(book)
	return args.Error(0)
}

func (m *MockBookUsecase) DeleteBook(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockBookUsecase) GetBookById(id int) (entity.Book, error) {
	args := m.Called(id)
	return args.Get(0).(entity.Book), args.Error(1)
}

func (m *MockBookUsecase) UpdateBook(book request.RequestBookDTO, id int) error {
	args := m.Called(id)
	return args.Error(1)
}

func TestGetBook(t *testing.T) {
	// Buat a mock instance of BookUsecase
	validate := validator.New()
	mockUsecase := new(MockBookUsecase)

	// Buat sebuah instance dari BookController dengan mock usecase sebelumnya
	controller := NewBookController(domain.Domain{BookUsecase: mockUsecase}, validate)

	// Set ekspektasi untuk GetBook method pada mock usecase
	mockBooks := []entity.Book{{Id: 1, Title: "Test Book 1", Author: "Sam"}}
	mockUsecase.On("GetBook").Return(mockBooks, nil)

	// Buat sebuah Fiber app untuk testing
	app := fiber.New()

	// Buat sebuah request
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	// Buat Route ke handler dari app fiber
	app.Get("/", controller.GetBook)

	// Lakukan Testing http
	resp, _ := app.Test(req)

	// Lakukan sesuatu dari results:
	if resp.StatusCode == fiber.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		fmt.Println(string(body)) // => Hello, World!
	}
	// assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestSaveBook(t *testing.T) {
	validate := validator.New()
	// Buat a mock instance of BookUsecase
	mockUsecase := new(MockBookUsecase)

	// Buat sebuah instance dari BookController dengan mock usecase sebelumnya
	controller := NewBookController(domain.Domain{BookUsecase: mockUsecase}, validate)

	// Set ekspektasi untuk GetBook method pada mock usecase
	mockUsecase.On("SaveBook", mock.Anything).Return(nil)

	// Buat sebuah Fiber app untuk testing
	app := fiber.New()

	// Buat sebuah request
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	// Buat Route ke handler dari app fiber
	app.Get("/", controller.GetBook)

	// Lakukan Testing http
	resp, _ := app.Test(req)

	// Lakukan sesuatu dari results:
	app.Post("/", controller.SaveBook)
	if resp.StatusCode == fiber.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		fmt.Println(string(body)) // => Hello, World!
	}
	// assert.Equal(t, http.StatusOK, resp.StatusCode)
}
