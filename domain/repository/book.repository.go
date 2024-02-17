package repository

import "github.com/rizama/favorite-book-tracker/domain/entity"

type BookRepository interface {
	FindBook(data any, condition ...any) error
	FindOneBook(data any, id int) error
	Create(book any) error
	Delete(book any, id int) error
	Update(book *entity.Book, id int) error
}
