package usecase_test

import (
	"context"
	"encoding/json"
	"strings"
	"testing"

	"github.com/adityatresnobudi/library-api/internal/dto"
	"github.com/adityatresnobudi/library-api/internal/mocks"
	"github.com/adityatresnobudi/library-api/internal/model"
	"github.com/adityatresnobudi/library-api/internal/shared"
	"github.com/adityatresnobudi/library-api/internal/usecase"
	"github.com/stretchr/testify/assert"
)

func createModelBook() model.Books {
	return model.Books{
		ID:          1,
		Title:       "test_book",
		Description: "test_book desc",
		Quantity:    2,
		Cover:       "soft",
		AuthorId:    1,
		Authors:     model.Authors{ID: 1, Name: "test"},
	}
}

func createBookDTO() dto.BooksDTO {
	return dto.BooksDTO{
		ID:          1,
		Title:       "test_book",
		Description: "test_book desc",
		Quantity:    2,
		Cover:       "soft",
		AuthorId:    1,
		Authors:     dto.AuthorsDTO{Name: "test"},
	}
}

func createBookPayload() dto.BookPayload {
	return dto.BookPayload{
		Title:       "test_book",
		Description: "test_book desc",
		Quantity:    2,
		Cover:       "soft",
		AuthorId:    1,
	}
}

func createBookResponse() dto.BookResponse {
	return dto.BookResponse{
		Title:       "test_book",
		Description: "test_book desc",
		Quantity:    2,
		Cover:       "soft",
	}
}

func MakeRequestBody(dto interface{}) *strings.Reader {
	payload, _ := json.Marshal(dto)
	return strings.NewReader(string(payload))
}

func TestBookUsecase_GetBooks(t *testing.T) {
	t.Run("should return dto books when all book fetched", func(t *testing.T) {
		// Given
		mockBookRepository := new(mocks.BookRepository)
		bu := usecase.NewBookUsecase(mockBookRepository)

		books := []model.Books{
			createModelBook(),
		}

		bookDto := []dto.BooksDTO{
			createBookDTO(),
		}

		ctx := context.Background()

		// When
		mockBookRepository.On("FindAllBooks", ctx, "").Return(books, nil)
		output, err := bu.GetBooks(ctx, "")

		// Then
		assert.Equal(t, bookDto, output)
		assert.Nil(t, err)
	})

	t.Run("should return dto books when book with related title fetched", func(t *testing.T) {
		// Given
		mockBookRepository := new(mocks.BookRepository)
		bu := usecase.NewBookUsecase(mockBookRepository)

		books := []model.Books{
			createModelBook(),
		}

		bookDto := []dto.BooksDTO{
			createBookDTO(),
		}

		ctx := context.Background()

		// When
		mockBookRepository.On("FindAllBooks", ctx, "test_book").Return(books, nil)
		output, err := bu.GetBooks(ctx, "test_book")

		// Then
		assert.Equal(t, bookDto, output)
		assert.Nil(t, err)
	})

	t.Run("should return ErrGettingBooks when book fetched failed", func(t *testing.T) {
		// Given
		mockBookRepository := new(mocks.BookRepository)
		bu := usecase.NewBookUsecase(mockBookRepository)
		c := context.Background()

		// When
		mockBookRepository.On("FindAllBooks", c, "").Return(nil, shared.ErrGettingBooks)
		output, err := bu.GetBooks(c, "")

		// Then
		assert.Nil(t, output)
		assert.ErrorIs(t, shared.ErrGettingBooks, err)
	})
}

func TestBookUsecase_AddBooks(t *testing.T) {
	t.Run("should return book response when book is successfully added", func(t *testing.T) {
		// Given
		mockBookRepository := new(mocks.BookRepository)
		bu := usecase.NewBookUsecase(mockBookRepository)

		book := createModelBook()
		book.ID = 0
		book.Authors.ID = 0
		book.Authors.Name = ""

		books := []model.Books{}

		bookPayload := createBookPayload()
		bookResponse := createBookResponse()

		ctx := context.Background()

		// When
		mockBookRepository.On("FindAllBooks", ctx, book.Title).Return(books, nil)
		mockBookRepository.On("AddOne", ctx, book).Return(book, nil)
		output, err := bu.AddBooks(ctx, bookPayload)

		// Then
		assert.Equal(t, bookResponse, output)
		assert.Nil(t, err)
	})

	t.Run("should return error duplicate book when new book is already added before", func(t *testing.T) {
		// Given
		mockBookRepository := new(mocks.BookRepository)
		bu := usecase.NewBookUsecase(mockBookRepository)

		book := createModelBook()
		book.ID = 0
		book.Authors.ID = 0
		book.Authors.Name = ""

		books := []model.Books{
			book,
		}

		bookPayload := createBookPayload()

		ctx := context.Background()

		// When
		mockBookRepository.On("FindAllBooks", ctx, book.Title).Return(books, nil)
		mockBookRepository.On("AddOne", ctx, book).Return(book, nil)
		output, err := bu.AddBooks(ctx, bookPayload)

		// Then
		assert.Equal(t, dto.BookResponse{}, output)
		assert.ErrorIs(t, shared.ErrDuplicateBook, err)
	})

	t.Run("should return error adding book when add book failed", func(t *testing.T) {
		// Given
		mockBookRepository := new(mocks.BookRepository)
		bu := usecase.NewBookUsecase(mockBookRepository)

		book := createModelBook()
		book.ID = 0
		book.Authors.ID = 0
		book.Authors.Name = ""

		books := []model.Books{}

		bookPayload := createBookPayload()

		ctx := context.Background()

		// When
		mockBookRepository.On("FindAllBooks", ctx, book.Title).Return(books, nil)
		mockBookRepository.On("AddOne", ctx, book).Return(book, shared.ErrAddingBooks)
		output, err := bu.AddBooks(ctx, bookPayload)

		// Then
		assert.Equal(t, dto.BookResponse{}, output)
		assert.ErrorIs(t, shared.ErrAddingBooks, err)
	})
}
