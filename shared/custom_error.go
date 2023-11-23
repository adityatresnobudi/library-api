package shared

import (
	"fmt"
	"net/http"
)

var (
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
