package shared

import (
	"fmt"
	"net/http"
)

var (
	ErrGettingBooks         = NewCustomError(http.StatusInternalServerError, "error getting all books")
	ErrAddingBooks          = NewCustomError(http.StatusInternalServerError, "error adding books")
	ErrGettingUsers         = NewCustomError(http.StatusInternalServerError, "error getting users")
	ErrCreateUsers          = NewCustomError(http.StatusInternalServerError, "error creating users")
	ErrGettingBorrowRecords = NewCustomError(http.StatusInternalServerError, "error getting borrow records")
	ErrCreateBorrowRecord   = NewCustomError(http.StatusInternalServerError, "error creating borrow record")
	ErrUpdateBorrowRecord   = NewCustomError(http.StatusInternalServerError, "error creating borrow record")
	ErrInvalidRequestBody   = NewCustomError(http.StatusBadRequest, "invalid request body")
	ErrDuplicateBook        = NewCustomError(http.StatusBadRequest, "book already exist")
	ErrAlreadyReturned      = NewCustomError(http.StatusBadRequest, "book already returned")
	ErrIdNotFound           = NewCustomError(http.StatusNotFound, "error id not found")
	ErrUserDoesntExist      = NewCustomError(http.StatusBadRequest, "invalid email or password")
	ErrFailedLogin          = NewCustomError(http.StatusInternalServerError, "error failed login")
	ErrInvalidPassword      = NewCustomError(http.StatusBadRequest, "invalid email or password")
	ErrInvalidAuthHeader    = NewCustomError(http.StatusBadRequest, "error invalid auth header")
	ErrInvalidToken         = NewCustomError(http.StatusBadRequest, "error invalid token")
	ErrUnauthorized         = NewCustomError(http.StatusBadRequest, "error unauthorized")
	ErrRecordNotFound       = NewCustomError(http.StatusBadRequest, "record not found")
)

type CustomError struct {
	StatusCode int    `json:"code"`
	Message    string `json:"message"`
}

type ErrorDTO struct {
	Message string `json:"message"`
}

func NewCustomError(statuscode int, message string) *CustomError {
	return &CustomError{
		StatusCode: statuscode,
		Message:    message,
	}
}

func (ce *CustomError) Error() string {
	return fmt.Sprintf("Error %d: %s", ce.StatusCode, ce.Message)
}

func (ce *CustomError) ToErrorDTO() *ErrorDTO {
	return &ErrorDTO{
		Message: ce.Message,
	}
}
