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
	// Create a mock implementation of repository.BookRepository
	mockRepo := new(MockBookRepository)

	// Create a book usecase with the mock repository
	bookUsecase := NewBookUsecase(mockRepo)

	// Mock data to be returned when FindBook is called
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

	// Set expectations for the FindBook method in the mock repository
	mockRepo.On("FindBook", &dest, mock.Anything).Return(nil).Run(func(args mock.Arguments) {
		books := args.Get(0).(*[]entity.Book)

		// Copy mockBooks into the provided slice (dereference it)
		*books = append(*books, mockBooks...)
	})

	// Call GetBook method
	books, err := bookUsecase.GetBook()

	// Assert that the method did not return an error
	assert.NoError(t, err)

	// Assert that the FindBook method was called with the correct arguments
	mockRepo.AssertExpectations(t)

	// Add more assertions based on your specific use case
	// For example, assert the content of the 'books' slice.
	// Remember to customize these assertions based on your actual implementation.
	assert.NoError(t, err, "Expected no error from GetBook")
	assert.NotNil(t, books)
	assert.Len(t, books, 1)
}

func TestSaveBook(t *testing.T) {
	// Create a mock implementation of repository.BookRepository
	mockRepo := new(MockBookRepository)

	// Create a book usecase with the mock repository
	bookUsecase := NewBookUsecase(mockRepo)

	// Mock data to be returned when FindBook is called
	mockBook := entity.Book{
		Title:     "New Book",
		Author:    "Sammm Author",
		Rating:    5,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Set expectations for the Create method in the mock repository
	mockRepo.On("Create", &mockBook).Return(nil) // You can customize the return value or error based on your scenario

	// Call the SaveBook method
	err := bookUsecase.SaveBook(mockBook)

	// Assert that the method did not return an error
	assert.NoError(t, err)

	// Assert that the Create method was called with the correct arguments
	mockRepo.AssertExpectations(t)
}
