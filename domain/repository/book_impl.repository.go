package repository

import (
	"log"

	"github.com/rizama/favorite-book-tracker/domain/entity"
	"gorm.io/gorm"
)

type bookRepositoryImpl struct {
	database *gorm.DB
}

// provider
func NewBookRepository(database *gorm.DB) BookRepository {
	return &bookRepositoryImpl{
		database: database,
	}
}

func (bookRepo *bookRepositoryImpl) FindBook(book any, condition ...any) error {
	// lakukan pencarian data ke database
	// setelah dapat, simpan hasilnya ke memory address book yang didapat dari parameter
	// secara langsung itu akan menyimpan ke variable book yang ada di usecase
	// sehingga fungsi ini tidak mengembalikan nilai
	result := bookRepo.database.Order("id desc").Find(book, condition...)

	// jika error, return
	if result.Error != nil {
		log.Printf("error fetching book:: %v", result.Error)
		return result.Error
	}

	return nil
}

func (bookRepo *bookRepositoryImpl) FindOneBook(book any, id int) error {
	result := bookRepo.database.Where("id = ?", id).First(book)

	// jika error, return
	if result.Error != nil {
		log.Printf("error fetching book:: %v", result.Error)
		return result.Error
	}

	return nil
}

func (bookRepo *bookRepositoryImpl) Create(book any) error {
	// store data
	result := bookRepo.database.Create(book)

	if result.Error != nil {
		log.Printf("error creating book:: %v", result.Error)
		return result.Error
	}

	return nil
}

func (bookRepo *bookRepositoryImpl) Delete(book any, id int) error {
	// delete data
	result := bookRepo.database.Where("id = ?", id).Delete(book)

	if result.Error != nil {
		log.Printf("error deleting book:: %v", result.Error)
		return result.Error
	}

	return nil
}

func (bookRepo *bookRepositoryImpl) Update(book *entity.Book, id int) error {
	// update data
	updateData := entity.Book{
		Title:     book.Title,
		Author:    book.Author,
		Rating:    book.Rating,
		UpdatedAt: book.UpdatedAt,
	}

	result := bookRepo.database.Model(&entity.Book{}).Where("id = ?", id).Updates(updateData)

	if result.Error != nil {
		log.Printf("error updating book:: %v", result.Error)
		return result.Error
	}

	return nil
}
