package usecase

import (
	"github.com/rizama/favorite-book-tracker/domain/entity"
	"github.com/rizama/favorite-book-tracker/domain/repository"
)

type bookUsecaseImpl struct {
	bookRepository repository.BookRepository
}

// constructor / provider
func NewBookUsecase(bookRepository repository.BookRepository) BookUsecase {
	return &bookUsecaseImpl{
		bookRepository: bookRepository,
	}
}

// *** methods bookUsecaseImpl
func (bookUsecase *bookUsecaseImpl) GetBook() ([]entity.Book, error) {
	// siapkan variable slice of book
	var book []entity.Book

	// kirim memori address book untuk penyimpanan data
	// nanti addres variable book akan di isi value nya langsung di usecase
	// sehingga di repository tidak perlu ada return
	if err := bookUsecase.bookRepository.FindBook(&book); err != nil {
		return nil, err
	}

	// return slice of book
	return book, nil
}

func (bookUsecase *bookUsecaseImpl) SaveBook(data entity.Book) error {

	// panggil repository untuk menyimpan data
	if err := bookUsecase.bookRepository.Create(&data); err != nil {
		return err
	}

	return nil
}

func (bookUsecase *bookUsecaseImpl) DeleteBook(id int) error {
	// siapkan variable book
	var book entity.Book

	// panggil repository untuk menghapus data
	if err := bookUsecase.bookRepository.Delete(book, id); err != nil {
		return err
	}

	return nil
}
