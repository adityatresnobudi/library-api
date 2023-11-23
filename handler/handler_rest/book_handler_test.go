package handlerrest_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/adityatresnobudi/library-api/dto"
	handler "github.com/adityatresnobudi/library-api/handler/handler_rest"
	"github.com/adityatresnobudi/library-api/mocks"
	"github.com/adityatresnobudi/library-api/router"
	"github.com/adityatresnobudi/library-api/shared"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func createBook() dto.BooksDTO {
	return dto.BooksDTO{
		ID:          1,
		Title:       "test_book",
		Description: "test_book desc",
		Quantity:    2,
		Cover:       "soft",
		AuthorId:    1,
		Authors:     dto.AuthorsDTO{ID: 1, Name: "test"},
	}
}

func MakeRequestBody(dto interface{}) *strings.Reader {
	payload, _ := json.Marshal(dto)
	return strings.NewReader(string(payload))
}

func TestBookHandler_GetBooks(t *testing.T) {
	t.Run("should return status code 200 when book fetched", func(t *testing.T) {
		// 1. setup router
		mockBookUsecase := new(mocks.BookUsecase)
		mockUserUsecase := new(mocks.UserUsecase)
		mockBorrowRecordUsecase := new(mocks.BorrowRecordUsecase)
		h := handler.NewHandler(mockBookUsecase, mockUserUsecase, mockBorrowRecordUsecase)
		books := []dto.BooksDTO{
			createBook(),
		}

		// 2. make request
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		req, _ := http.NewRequest("GET", "/books", nil)
		c.Request = req

		mockBookUsecase.On("GetBooks", c.Request.Context(), "").Return(books, nil)
		expectedResp, _ := json.Marshal(dto.JsonResponse{Data: books})
		h.GetBooks(c)

		// 3. assert
		assert.Equal(t, http.StatusOK, w.Code)
		str := strings.Trim(w.Body.String(), "\n")
		assert.Equal(t, string(expectedResp), str)
	})

	t.Run("should return status code 500 when books fetched failed", func(t *testing.T) {
		// 1. setup router
		mockBookUsecase := new(mocks.BookUsecase)
		mockUserUsecase := new(mocks.UserUsecase)
		mockBorrowRecordUsecase := new(mocks.BorrowRecordUsecase)
		h := handler.NewHandler(mockBookUsecase, mockUserUsecase, mockBorrowRecordUsecase)
		router := router.NewRouter(h)
		mockBookUsecase.On("GetBooks", mock.Anything, "").Return(nil, errors.New("internal server error"))
		expectedResp, _ := json.Marshal(dto.JsonResponse{Message: "internal server error"})

		// 2. make request
		rec := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/books", nil)
		router.ServeHTTP(rec, req)

		// 3. assert
		assert.Equal(t, 500, rec.Code)
		str := strings.Trim(rec.Body.String(), "\n")
		assert.Equal(t, string(expectedResp), str)
	})

	t.Run("should return status code 200 when book with same title is fetched", func(t *testing.T) {
		// 1. setup router
		mockBookUsecase := new(mocks.BookUsecase)
		mockUserUsecase := new(mocks.UserUsecase)
		mockBorrowRecordUsecase := new(mocks.BorrowRecordUsecase)
		h := handler.NewHandler(mockBookUsecase, mockUserUsecase, mockBorrowRecordUsecase)
		router := router.NewRouter(h)
		books := []dto.BooksDTO{
			createBook(),
		}
		mockBookUsecase.On("GetBooks", mock.Anything, "test_book").Return(books, nil)
		expectedResp, _ := json.Marshal(dto.JsonResponse{Data: books})

		// 2. make request
		rec := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/books?title=test_book", nil)
		router.ServeHTTP(rec, req)

		// 3. assert
		assert.Equal(t, 200, rec.Code)
		str := strings.Trim(rec.Body.String(), "\n")
		assert.Equal(t, string(expectedResp), str)
	})
}

func TestBookHandler_AddBooks(t *testing.T) {
	t.Run("should return status code 201 when book is added", func(t *testing.T) {
		// 1. setup router
		mockBookUsecase := new(mocks.BookUsecase)
		mockUserUsecase := new(mocks.UserUsecase)
		mockBorrowRecordUsecase := new(mocks.BorrowRecordUsecase)
		h := handler.NewHandler(mockBookUsecase, mockUserUsecase, mockBorrowRecordUsecase)
		router := router.NewRouter(h)
		newBooks := dto.BookPayload{
			Title:       "test_book",
			Description: "test_book desc",
			Quantity:    2,
			Cover:       "soft",
			AuthorId:    1,
		}
		bookResponse := dto.BookResponse{
			ID:          1,
			Title:       "test_book",
			Description: "test_book desc",
			Quantity:    2,
			Cover:       "soft",
		}
		mockBookUsecase.On("AddBooks", mock.Anything, newBooks).Return(bookResponse, nil)
		message := fmt.Sprintf("successfully add new book with id %d", bookResponse.ID)
		expectedResp, _ := json.Marshal(dto.JsonResponse{Message: message, Data: bookResponse})

		// 2. make request
		rec := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/books", MakeRequestBody(newBooks))
		router.ServeHTTP(rec, req)

		// 3. assert
		assert.Equal(t, rec.Code, 201)
		str := strings.Trim(rec.Body.String(), "\n")
		assert.Equal(t, str, string(expectedResp))
	})

	t.Run("should return status code 400 when book is added without title, description, quantity, or authorId", func(t *testing.T) {
		// 1. setup router
		mockBookUsecase := new(mocks.BookUsecase)
		mockUserUsecase := new(mocks.UserUsecase)
		mockBorrowRecordUsecase := new(mocks.BorrowRecordUsecase)
		h := handler.NewHandler(mockBookUsecase, mockUserUsecase, mockBorrowRecordUsecase)
		router := router.NewRouter(h)
		newBooks := dto.BookPayload{
			Title:       "test_book",
			Description: "test_book desc",
			Quantity:    2,
			Cover:       "soft",
		}
		mockBookUsecase.On("AddBooks", mock.Anything, newBooks).Return(dto.BookResponse{}, errors.New("internal server error"))
		expectedResp, _ := json.Marshal(dto.JsonResponse{Message: "invalid request body"})

		// 2. make request
		rec := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/books", MakeRequestBody(newBooks))
		router.ServeHTTP(rec, req)

		// 3. assert
		assert.Equal(t, rec.Code, 400)
		str := strings.Trim(rec.Body.String(), "\n")
		assert.Equal(t, str, string(expectedResp))
	})

	t.Run("should return status code 400 when add book that already added", func(t *testing.T) {
		// 1. setup router
		mockBookUsecase := new(mocks.BookUsecase)
		mockUserUsecase := new(mocks.UserUsecase)
		mockBorrowRecordUsecase := new(mocks.BorrowRecordUsecase)
		h := handler.NewHandler(mockBookUsecase, mockUserUsecase, mockBorrowRecordUsecase)
		router := router.NewRouter(h)
		newBooks := dto.BookPayload{
			Title:       "test_book",
			Description: "test_book desc",
			Quantity:    2,
			Cover:       "soft",
			AuthorId:    1,
		}
		mockBookUsecase.On("AddBooks", mock.Anything, newBooks).Return(dto.BookResponse{}, shared.ErrDuplicateBook)
		expectedResp, _ := json.Marshal(dto.JsonResponse{Message: shared.ErrDuplicateBook.Message})

		// 2. make request
		rec := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/books", MakeRequestBody(newBooks))
		router.ServeHTTP(rec, req)

		// 3. assert
		assert.Equal(t, rec.Code, 400)
		str := strings.Trim(rec.Body.String(), "\n")
		assert.Equal(t, str, string(expectedResp))
	})

	t.Run("should return status code 500 when add book failed", func(t *testing.T) {
		// 1. setup router
		mockBookUsecase := new(mocks.BookUsecase)
		mockUserUsecase := new(mocks.UserUsecase)
		mockBorrowRecordUsecase := new(mocks.BorrowRecordUsecase)
		h := handler.NewHandler(mockBookUsecase, mockUserUsecase, mockBorrowRecordUsecase)
		router := router.NewRouter(h)
		newBooks := dto.BookPayload{
			Title:       "test_book",
			Description: "test_book desc",
			Quantity:    2,
			Cover:       "soft",
			AuthorId:    1,
		}
		mockBookUsecase.On("AddBooks", mock.Anything, newBooks).Return(dto.BookResponse{}, shared.ErrAddingBooks)
		expectedResp, _ := json.Marshal(dto.JsonResponse{Message: shared.ErrAddingBooks.Message})

		// 2. make request
		rec := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/books", MakeRequestBody(newBooks))
		router.ServeHTTP(rec, req)

		// 3. assert
		assert.Equal(t, rec.Code, 500)
		str := strings.Trim(rec.Body.String(), "\n")
		assert.Equal(t, str, string(expectedResp))
	})
}
