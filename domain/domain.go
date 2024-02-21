package domain

import (
	"github.com/rizama/favorite-book-tracker/domain/repository"
	"github.com/rizama/favorite-book-tracker/domain/usecase"
	"github.com/rizama/favorite-book-tracker/infrastructure"
)

type Domain struct {
	BookUsecase usecase.BookUsecase
}

func ConstructDomain() Domain {
	postgresConn := infrastructure.NewPostgresConnection()

	bookRepository := repository.NewBookRepository(postgresConn)
	bookUsecase := usecase.NewBookUsecase(bookRepository)

	return Domain{
		BookUsecase: bookUsecase,
	}
}
