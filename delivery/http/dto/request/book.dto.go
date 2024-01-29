package request

import (
	"time"

	"github.com/rizama/favorite-book-tracker/domain/entity"
)

type RequestBookDTO struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating int8   `json:"rating"`
}

func (this RequestBookDTO) ToBookEntity() entity.Book {
	return entity.Book{
		Title:     this.Title,
		Author:    this.Author,
		Rating:    this.Rating,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil,
	}
}
