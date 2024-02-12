package repository

type BookRepository interface {
	FindBook(data any, condition ...any) error
	Create(value any) error
	Delete(value any, id int) error
}
