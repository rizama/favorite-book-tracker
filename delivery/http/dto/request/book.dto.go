package request

import (
	"time"

	"github.com/rizama/favorite-book-tracker/domain/entity"
)

type RequestBookDTO struct {
	Title  string `validate:"required,min=5,max=20"`
	Author string `validate:"required,min=5,max=20"`
	Rating int8   `validate:"required,gte=1,lte=5"`
}

func (book RequestBookDTO) ToBookEntity() entity.Book {
	return entity.Book{
		Title:     book.Title,
		Author:    book.Author,
		Rating:    book.Rating,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil,
	}
}
