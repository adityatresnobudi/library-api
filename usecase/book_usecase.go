package usecase

import (
	"context"

	"github.com/adityatresnobudi/library-api/dto"
	"github.com/adityatresnobudi/library-api/model"
	"github.com/adityatresnobudi/library-api/repository"
	"github.com/adityatresnobudi/library-api/shared"
)

type bookUsecase struct {
	bookRepo repository.BookRepository
}

type BookUsecase interface {
	GetBooks(ctx context.Context, title string) ([]dto.BooksDTO, error)
	AddBooks(ctx context.Context, book dto.BookPayload) (dto.BookResponse, error)
}

func NewBookUsecase(bookRepo repository.BookRepository) BookUsecase {
	return &bookUsecase{
		bookRepo: bookRepo,
	}
}

func (bu *bookUsecase) GetBooks(ctx context.Context, title string) ([]dto.BooksDTO, error) {
	books := []dto.BooksDTO{}
	book := dto.BooksDTO{}
	uc, err := bu.bookRepo.FindAllBooks(ctx, title)
	if err != nil {
		return nil, shared.ErrGettingBooks
	}

	for _, b := range uc {
		book.ID = b.ID
		book.Title = b.Title
		book.Description = b.Description
		book.Quantity = b.Quantity
		book.Cover = b.Cover
		book.AuthorId = b.AuthorId
		book.Authors.Name = b.Authors.Name
		books = append(books, book)
	}

	return books, nil
}

func (bu *bookUsecase) AddBooks(ctx context.Context, book dto.BookPayload) (dto.BookResponse, error) {
	newBook := model.Books{}
	response := dto.BookResponse{}
	length, _ := bu.bookRepo.FindAllBooks(ctx, book.Title)
	if len(length) != 0 {
		return dto.BookResponse{}, shared.ErrDuplicateBook
	}

	newBook.Title = book.Title
	newBook.Description = book.Description
	newBook.Quantity = book.Quantity
	newBook.Cover = book.Cover
	newBook.AuthorId = book.AuthorId

	uc, err := bu.bookRepo.AddOne(ctx, newBook)
	if err != nil {
		return dto.BookResponse{}, shared.ErrAddingBooks
	}

	response.ID = uc.ID
	response.Title = uc.Title
	response.Description = uc.Description
	response.Quantity = uc.Quantity
	response.Cover = uc.Cover

	return response, nil
}
