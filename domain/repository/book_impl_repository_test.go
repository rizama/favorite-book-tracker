package repository

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/rizama/favorite-book-tracker/domain/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	// Setup database testing
	dsn := "host=localhost user=sam password='sam' dbname=favorite_books port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
}

func TestFindBook(t *testing.T) {
	// buat sebuah sample book for testing
	book := entity.Book{
		Title:     "Sam Book",
		Author:    "Sam Author",
		Rating:    5,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Masukan sample book ke database test terlebih dahulu
	if err := db.Create(&book).Error; err != nil {
		t.Fatal(err)
	}

	// buat sebuah book repository
	bookRepo := NewBookRepository(db)

	// Panggil FindBook method
	var foundBook entity.Book
	err2 := bookRepo.FindBook(&foundBook, "title = ?", "Test Book")

	// Gunakan assertions untuk mengecek pencarian buku berhasil atau tidak
	assert.NoError(t, err2, "Expected no error from FindBook")
	assert.Equal(t, "Test Book", foundBook.Title, "Expected book title to be 'Test Book'")
}

func TestCreateBook(t *testing.T) {
	// buat sebuah book repository
	bookRepo := NewBookRepository(db)

	// buat sebuah sample book untuk testing
	book := entity.Book{
		Title:     "Rizky Book",
		Author:    "Sam Author",
		Rating:    5,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Panggil Create method
	err1 := bookRepo.Create(&book)
	fmt.Println(err1, "*err1")
	// Use assertions to check if the book was created successfully
	assert.NoError(t, err1, "Expected no error from Create")

	// Gunakan assertions untuk mengecek penyimpanan buku berhasil atau tidak
	var foundBook entity.Book
	err := db.First(&foundBook, "title = ?", "Test Book").Error
	assert.NoError(t, err, "Expected no error when querying the database")
	assert.Equal(t, "Test Book", foundBook.Title, "Expected book title to be 'Test Book'")
}
