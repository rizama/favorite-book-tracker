package usecase

import (
	"testing"
	"time"

	"github.com/rizama/favorite-book-tracker/domain/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockBookRepository struct {
	mock.Mock
}

func (m *MockBookRepository) FindBook(book any, condition ...any) error {
	args := m.Called(book, condition)
	return args.Error(0)
}

func (m *MockBookRepository) Create(value any) error {
	args := m.Called(value)
	return args.Error(0)
}

func TestGetBook(t *testing.T) {
	// Buat sebuah mock implementation dari repository.BookRepository
	mockRepo := new(MockBookRepository)

	// Buat sebuah book usecase dengan mock repository diatas
	bookUsecase := NewBookUsecase(mockRepo)

	// Mock data yang akan dikembalikan ketika FIndBook dipanggil
	mockBooks := []entity.Book{
		{
			Title:     "Sam2 Book",
			Author:    "Sam2 Author",
			Rating:    5,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	var dest []entity.Book

	// Set expectations untuk FindBook method pada mock repository
	mockRepo.On("FindBook", &dest, mock.Anything).Return(nil).Run(func(args mock.Arguments) {
		books := args.Get(0).(*[]entity.Book)

		// Copy mockBooks ke slice yang diberikan yaitu variable dest, dest disini pada real casenya adalah book
		*books = append(*books, mockBooks...)
	})

	// Panggil GetBook method
	books, err := bookUsecase.GetBook()

	// Assert method tersebut tidak menimbulkan error
	assert.NoError(t, err)

	// Assert FindBook method telah terpanggil dengan argument ya tepat
	mockRepo.AssertExpectations(t)

	// Tambahkan asserts lainnya
	assert.NoError(t, err, "Expected no error from GetBook")
	assert.NotNil(t, books)
	assert.Len(t, books, 1)
}

func TestSaveBook(t *testing.T) {
	// Buat sebuah mock implementation of repository.BookRepository
	mockRepo := new(MockBookRepository)

	// Buat sebuah book usecase with the mock repository
	bookUsecase := NewBookUsecase(mockRepo)

	// Mock data to be returned when FindBook is called
	mockBook := entity.Book{
		Title:     "New Book",
		Author:    "Sammm Author",
		Rating:    5,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Buat Create method mock
	mockRepo.On("Create", &mockBook).Return(nil)

	// Panggil SaveBook method
	err := bookUsecase.SaveBook(mockBook)

	// Lakukan assert
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
