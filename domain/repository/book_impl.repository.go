package repository

import (
	"fmt"
	"log"

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
		log.Println(fmt.Sprintf("error fetching book:: %v", result.Error))
		return result.Error
	}

	return nil
}

func (bookRepo bookRepositoryImpl) Create(value any) error {
	// store data
	result := bookRepo.database.Create(value)

	if result.Error != nil {
		log.Println(fmt.Sprintf("error creating book:: %v", result.Error))
		return result.Error
	}

	return nil
}
