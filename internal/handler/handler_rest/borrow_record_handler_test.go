package handlerrest_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/adityatresnobudi/library-api/internal/dto"
	handler "github.com/adityatresnobudi/library-api/internal/handler/handler_rest"
	"github.com/adityatresnobudi/library-api/internal/mocks"
	"github.com/adityatresnobudi/library-api/internal/router"
	"github.com/adityatresnobudi/library-api/internal/shared"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func createBorrowRecord() dto.BorrowRecordsDTO {
	return dto.BorrowRecordsDTO{
		ID:            1,
		UserId:        2,
		BookId:        3,
		BorrowingDate: "2023-01-01 10:00:00",
		ReturningDate: "0001-01-01 00:00:00",
	}
}

func TestBorrowRecordHandler_BorrowBook(t *testing.T) {
	t.Setenv("ENV_MODE", "testing")
	t.Run("should return status code 200 when successfully borrowed book", func(t *testing.T) {
		// 1. setup router
		mockBookUsecase := new(mocks.BookUsecase)
		mockUserUsecase := new(mocks.UserUsecase)
		mockBorrowRecordUsecase := new(mocks.BorrowRecordUsecase)
		h := handler.NewHandler(mockBookUsecase, mockUserUsecase, mockBorrowRecordUsecase)
		record := createBorrowRecord()

		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Set("id", 1)
		req, _ := http.NewRequest("POST", "/records/borrow", MakeRequestBody(record))
		c.Request = req

		mockBorrowRecordUsecase.On("BorrowBook", c.Request.Context(), record, 1).Return(record, nil)
		expectedResp, _ := json.Marshal(dto.JsonResponse{Data: record})

		// 2. make request
		h.BorrowBook(c)

		// 3. assert
		assert.Equal(t, rec.Code, 200)
		str := strings.Trim(rec.Body.String(), "\n")
		assert.Equal(t, str, string(expectedResp))
	})

	t.Run("should return status code 400 when invalid request body", func(t *testing.T) {
		// 1. setup router
		mockBookUsecase := new(mocks.BookUsecase)
		mockUserUsecase := new(mocks.UserUsecase)
		mockBorrowRecordUsecase := new(mocks.BorrowRecordUsecase)
		h := handler.NewHandler(mockBookUsecase, mockUserUsecase, mockBorrowRecordUsecase)
		router := router.NewRouter(h)

		record := dto.BorrowRecordsDTO{
			BookId:        1,
			BorrowingDate: "2023-01-01 10:00:00",
		}

		mockBorrowRecordUsecase.On("BorrowBook", mock.Anything, record, 1).Return(dto.BorrowRecordsDTO{}, shared.ErrInvalidRequestBody)
		expectedResp, _ := json.Marshal(dto.JsonResponse{Message: shared.ErrInvalidRequestBody.Message})

		// 2. make request
		rec := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/records/borrow", MakeRequestBody(record))
		router.ServeHTTP(rec, req)

		// 3. assert
		assert.Equal(t, rec.Code, 400)
		str := strings.Trim(rec.Body.String(), "\n")
		assert.Equal(t, str, string(expectedResp))
	})

	t.Run("should return status code 500 when borrow books failed", func(t *testing.T) {
		// 1. setup router
		mockBookUsecase := new(mocks.BookUsecase)
		mockUserUsecase := new(mocks.UserUsecase)
		mockBorrowRecordUsecase := new(mocks.BorrowRecordUsecase)
		h := handler.NewHandler(mockBookUsecase, mockUserUsecase, mockBorrowRecordUsecase)
		record := createBorrowRecord()

		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Set("id", 1)
		req, _ := http.NewRequest("POST", "/records/borrow", MakeRequestBody(record))
		c.Request = req

		mockBorrowRecordUsecase.On("BorrowBook", c.Request.Context(), record, 1).Return(dto.BorrowRecordsDTO{}, shared.ErrCreateBorrowRecord)

		// 2. make request
		h.BorrowBook(c)

		// 3. assert
		assert.Equal(t, c.Errors[0].Err, shared.ErrCreateBorrowRecord)
	})
}

func TestBorrowRecordHandler_ReturnBook(t *testing.T) {
	t.Setenv("ENV_MODE", "testing")
	t.Run("should return status code 200 when successfully return book", func(t *testing.T) {
		// 1. setup router
		mockBookUsecase := new(mocks.BookUsecase)
		mockUserUsecase := new(mocks.UserUsecase)
		mockBorrowRecordUsecase := new(mocks.BorrowRecordUsecase)
		h := handler.NewHandler(mockBookUsecase, mockUserUsecase, mockBorrowRecordUsecase)
		borrowRecordID := 1
		record := createBorrowRecord()

		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Set("id", 1)
		c.AddParam("id", "1")
		req, _ := http.NewRequest("PUT", "/records/return/1", nil)
		c.Request = req

		mockBorrowRecordUsecase.On("GetBorrowRecordByID", c.Request.Context(), borrowRecordID).Return(record, nil)
		mockBorrowRecordUsecase.On("ReturnBook", c.Request.Context(), record, 1).Return(record, nil)
		expectedResp, _ := json.Marshal(dto.JsonResponse{Data: record})

		// 2. make request
		h.ReturnBook(c)

		// 3. assert
		assert.Equal(t, rec.Code, 200)
		str := strings.Trim(rec.Body.String(), "\n")
		assert.Equal(t, str, string(expectedResp))
	})

	t.Run("should return status code 404 when borrowRecordID not found", func(t *testing.T) {
		// 1. setup router
		mockBookUsecase := new(mocks.BookUsecase)
		mockUserUsecase := new(mocks.UserUsecase)
		mockBorrowRecordUsecase := new(mocks.BorrowRecordUsecase)
		h := handler.NewHandler(mockBookUsecase, mockUserUsecase, mockBorrowRecordUsecase)

		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.AddParam("id", "a")

		// 2. make request
		req, _ := http.NewRequest("PUT", "/records/return/a", nil)
		c.Request = req
		h.ReturnBook(c)

		// 3. assert
		assert.Equal(t, c.Errors[0].Err, shared.ErrIdNotFound)
	})

	t.Run("should return status code 500 when borrow record fetched failed", func(t *testing.T) {
		// 1. setup router
		mockBookUsecase := new(mocks.BookUsecase)
		mockUserUsecase := new(mocks.UserUsecase)
		mockBorrowRecordUsecase := new(mocks.BorrowRecordUsecase)
		h := handler.NewHandler(mockBookUsecase, mockUserUsecase, mockBorrowRecordUsecase)
		borrowRecordid := 1

		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.AddParam("id", "1")
		req, _ := http.NewRequest("PUT", "/records/return/1", nil)
		c.Request = req

		mockBorrowRecordUsecase.On("GetBorrowRecordByID", c.Request.Context(), borrowRecordid).Return(dto.BorrowRecordsDTO{}, shared.ErrGettingBorrowRecords)

		// 2. make request
		h.ReturnBook(c)

		// 3. assert
		assert.Equal(t, c.Errors[0].Err, shared.ErrGettingBorrowRecords)
	})

	t.Run("should return status code 400 when book already returned", func(t *testing.T) {
		// 1. setup router
		mockBookUsecase := new(mocks.BookUsecase)
		mockUserUsecase := new(mocks.UserUsecase)
		mockBorrowRecordUsecase := new(mocks.BorrowRecordUsecase)
		h := handler.NewHandler(mockBookUsecase, mockUserUsecase, mockBorrowRecordUsecase)
		borrowRecordID := 1
		record := createBorrowRecord()
		record.ReturningDate = "2023-01-01 12:00:00"

		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.AddParam("id", "1")
		req, _ := http.NewRequest("PUT", "/records/return/1", nil)
		c.Request = req

		mockBorrowRecordUsecase.On("GetBorrowRecordByID", c.Request.Context(), borrowRecordID).Return(record, nil)

		// 2. make request
		h.ReturnBook(c)

		// 3. assert
		assert.Equal(t, c.Errors[0].Err, shared.ErrAlreadyReturned)
	})

	t.Run("should return status code 500 when return book failed", func(t *testing.T) {
		// 1. setup router
		mockBookUsecase := new(mocks.BookUsecase)
		mockUserUsecase := new(mocks.UserUsecase)
		mockBorrowRecordUsecase := new(mocks.BorrowRecordUsecase)
		h := handler.NewHandler(mockBookUsecase, mockUserUsecase, mockBorrowRecordUsecase)
		borrowRecordId := 1
		record := createBorrowRecord()

		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Set("id", 1)
		c.AddParam("id", "1")
		req, _ := http.NewRequest("PUT", "/records/return/1", nil)
		c.Request = req

		mockBorrowRecordUsecase.On("GetBorrowRecordByID", c.Request.Context(), borrowRecordId).Return(record, nil)
		mockBorrowRecordUsecase.On("ReturnBook", c.Request.Context(), record, 1).Return(dto.BorrowRecordsDTO{}, shared.ErrUpdateBorrowRecord)

		// 2. make request
		h.ReturnBook(c)

		// 3. assert
		assert.Equal(t, c.Errors[0].Err, shared.ErrUpdateBorrowRecord)
	})
}
