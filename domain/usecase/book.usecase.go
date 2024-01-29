package usecase

import (
	"github.com/rizama/favorite-book-tracker/domain/entity"
)

type BookUsecase interface {
	GetBook() ([]entity.Book, error)
	SaveBook(action entity.Book) error
}
