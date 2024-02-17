package usecase

import (
	"github.com/rizama/favorite-book-tracker/delivery/http/dto/request"
	"github.com/rizama/favorite-book-tracker/domain/entity"
)

type BookUsecase interface {
	GetBook() ([]entity.Book, error)
	GetBookById(id int) (entity.Book, error)
	SaveBook(action entity.Book) error
	DeleteBook(id int) error
	UpdateBook(action request.RequestBookDTO, id int) error
}
