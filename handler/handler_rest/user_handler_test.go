package handlerrest_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/adityatresnobudi/library-api/dto"
	handler "github.com/adityatresnobudi/library-api/handler/handler_rest"
	"github.com/adityatresnobudi/library-api/mocks"
	"github.com/adityatresnobudi/library-api/router"
	"github.com/adityatresnobudi/library-api/shared"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func createUser() dto.UserPayload {
	return dto.UserPayload{
		ID:       1,
		Name:     "aditbuddy",
		Email:    "aditbuddy@abc.com",
		Phone:    "081290882428",
		Password: "relativiteam",
	}
}

func createUserResponse() dto.UserResponse {
	return dto.UserResponse{
		ID:    1,
		Name:  "aditbuddy",
		Email: "aditbuddy@abc.com",
		Phone: "081290882428",
	}
}

func createLoginPayload() dto.LoginRequest {
	return dto.LoginRequest{
		Email:    "aditbuddy@abc.com",
		Password: "relativiteam",
	}
}

func createLoginResponse() dto.LoginResponse {
	return dto.LoginResponse{
		AccessToken: "aaabbbccc",
	}
}

func TestUserHandler_GetUsers(t *testing.T) {
	t.Run("should return status code 200 when users fetched", func(t *testing.T) {
		// 1. setup router
		mockBookUsecase := new(mocks.BookUsecase)
		mockUserUsecase := new(mocks.UserUsecase)
		mockBorrowRecordUsecase := new(mocks.BorrowRecordUsecase)
		h := handler.NewHandler(mockBookUsecase, mockUserUsecase, mockBorrowRecordUsecase)
		router := router.NewRouter(h)
		users := []dto.UserResponse{
			createUserResponse(),
		}
		mockUserUsecase.On("GetUsers", mock.Anything, "").Return(users, nil)
		expectedResp, _ := json.Marshal(dto.JsonResponse{Data: users})

		// 2. make request
		rec := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/users", nil)
		router.ServeHTTP(rec, req)

		// 3. assert
		assert.Equal(t, rec.Code, 200)
		str := strings.Trim(rec.Body.String(), "\n")
		assert.Equal(t, str, string(expectedResp))
	})

	t.Run("should return status code 500 when users fetched failed", func(t *testing.T) {
		// 1. setup router
		mockBookUsecase := new(mocks.BookUsecase)
		mockUserUsecase := new(mocks.UserUsecase)
		mockBorrowRecordUsecase := new(mocks.BorrowRecordUsecase)
		h := handler.NewHandler(mockBookUsecase, mockUserUsecase, mockBorrowRecordUsecase)
		router := router.NewRouter(h)
		mockUserUsecase.On("GetUsers", mock.Anything, "").Return(nil, errors.New("internal server error"))
		expectedResp, _ := json.Marshal(dto.JsonResponse{Message: "internal server error"})

		// 2. make request
		rec := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/users", nil)
		router.ServeHTTP(rec, req)

		// 3. assert
		assert.Equal(t, rec.Code, 500)
		str := strings.Trim(rec.Body.String(), "\n")
		assert.Equal(t, str, string(expectedResp))
	})

	t.Run("should return status code 200 when username exist", func(t *testing.T) {
		// 1. setup router
		mockBookUsecase := new(mocks.BookUsecase)
		mockUserUsecase := new(mocks.UserUsecase)
		mockBorrowRecordUsecase := new(mocks.BorrowRecordUsecase)
		h := handler.NewHandler(mockBookUsecase, mockUserUsecase, mockBorrowRecordUsecase)
		router := router.NewRouter(h)
		users := []dto.UserResponse{
			createUserResponse(),
		}
		mockUserUsecase.On("GetUsers", mock.Anything, "aditbuddy").Return(users, nil)
		expectedResp, _ := json.Marshal(dto.JsonResponse{Data: users})

		// 2. make request
		rec := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/users?name=aditbuddy", nil)
		router.ServeHTTP(rec, req)

		// 3. assert
		assert.Equal(t, rec.Code, 200)
		str := strings.Trim(rec.Body.String(), "\n")
		assert.Equal(t, str, string(expectedResp))
	})

	t.Run("should return status code 200 and empty string when username did not exist", func(t *testing.T) {
		// 1. setup router
		mockBookUsecase := new(mocks.BookUsecase)
		mockUserUsecase := new(mocks.UserUsecase)
		mockBorrowRecordUsecase := new(mocks.BorrowRecordUsecase)
		h := handler.NewHandler(mockBookUsecase, mockUserUsecase, mockBorrowRecordUsecase)
		router := router.NewRouter(h)
		users := []dto.UserResponse{
			createUserResponse(),
		}
		mockUserUsecase.On("GetUsers", mock.Anything, "natnat").Return(users, nil)
		expectedResp, _ := json.Marshal(dto.JsonResponse{Data: users})

		// 2. make request
		rec := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/users?name=natnat", nil)
		router.ServeHTTP(rec, req)

		// 3. assert
		assert.Equal(t, rec.Code, 200)
		str := strings.Trim(rec.Body.String(), "\n")
		assert.Equal(t, str, string(expectedResp))
	})
}

func TestUserHandler_CreateUsers(t *testing.T) {
	t.Run("should return status code 201 when create user successful", func(t *testing.T) {
		// 1. setup router
		mockBookUsecase := new(mocks.BookUsecase)
		mockUserUsecase := new(mocks.UserUsecase)
		mockBorrowRecordUsecase := new(mocks.BorrowRecordUsecase)
		h := handler.NewHandler(mockBookUsecase, mockUserUsecase, mockBorrowRecordUsecase)
		router := router.NewRouter(h)
		userPayload := createUser()
		userResponse := dto.UserResponse{
			ID:    userPayload.ID,
			Name:  userPayload.Name,
			Email: userPayload.Email,
			Phone: userPayload.Phone,
		}

		mockUserUsecase.On("CreateUsers", mock.Anything, userPayload).Return(userResponse, nil)
		expectedResp, _ := json.Marshal(dto.JsonResponse{Data: userResponse})

		// 2. make request
		rec := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/users", MakeRequestBody(userPayload))
		router.ServeHTTP(rec, req)

		// 3. assert
		assert.Equal(t, rec.Code, 201)
		str := strings.Trim(rec.Body.String(), "\n")
		assert.Equal(t, str, string(expectedResp))
	})

	t.Run("should return status code 400 when create with invalid user request", func(t *testing.T) {
		// 1. setup router
		mockBookUsecase := new(mocks.BookUsecase)
		mockUserUsecase := new(mocks.UserUsecase)
		mockBorrowRecordUsecase := new(mocks.BorrowRecordUsecase)
		h := handler.NewHandler(mockBookUsecase, mockUserUsecase, mockBorrowRecordUsecase)
		router := router.NewRouter(h)
		userPayload := dto.UserPayload{
			ID:    1,
			Name:  "aditbuddy",
			Email: "aditbuddy@abc.com",
			Phone: "081290882428",
		}

		mockUserUsecase.On("CreateUsers", mock.Anything, userPayload).Return(dto.UserResponse{}, shared.ErrInvalidRequestBody)
		expectedResp, _ := json.Marshal(dto.JsonResponse{Message: shared.ErrInvalidRequestBody.Message})

		// 2. make request
		rec := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/users", MakeRequestBody(userPayload))
		router.ServeHTTP(rec, req)

		// 3. assert
		assert.Equal(t, rec.Code, 400)
		str := strings.Trim(rec.Body.String(), "\n")
		assert.Equal(t, str, string(expectedResp))
	})

	t.Run("should return status code 500 when create user failed", func(t *testing.T) {
		// 1. setup router
		mockBookUsecase := new(mocks.BookUsecase)
		mockUserUsecase := new(mocks.UserUsecase)
		mockBorrowRecordUsecase := new(mocks.BorrowRecordUsecase)
		h := handler.NewHandler(mockBookUsecase, mockUserUsecase, mockBorrowRecordUsecase)
		router := router.NewRouter(h)
		userPayload := createUser()
		mockUserUsecase.On("CreateUsers", mock.Anything, userPayload).Return(dto.UserResponse{}, shared.ErrCreateUsers)
		expectedResp, _ := json.Marshal(dto.JsonResponse{Message: shared.ErrCreateUsers.Message})

		// 2. make request
		rec := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/users", MakeRequestBody(userPayload))
		router.ServeHTTP(rec, req)

		// 3. assert
		assert.Equal(t, rec.Code, 500)
		str := strings.Trim(rec.Body.String(), "\n")
		assert.Equal(t, str, string(expectedResp))
	})
}

func TestUserHandler_LoginUsers(t *testing.T) {
	t.Run("should return status code 200 when login successful", func(t *testing.T) {
		// 1. setup router
		mockBookUsecase := new(mocks.BookUsecase)
		mockUserUsecase := new(mocks.UserUsecase)
		mockBorrowRecordUsecase := new(mocks.BorrowRecordUsecase)
		h := handler.NewHandler(mockBookUsecase, mockUserUsecase, mockBorrowRecordUsecase)
		router := router.NewRouter(h)
		loginPayload := createLoginPayload()
		loginResponse := createLoginResponse()

		mockUserUsecase.On("LoginUser", mock.Anything, loginPayload).Return(loginResponse, nil)
		expectedResp, _ := json.Marshal(loginResponse)

		// 2. make request
		rec := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/users/login", MakeRequestBody(loginPayload))
		router.ServeHTTP(rec, req)

		// 3. assert
		assert.Equal(t, rec.Code, 200)
		str := strings.Trim(rec.Body.String(), "\n")
		assert.Equal(t, str, string(expectedResp))
	})

	t.Run("should return status code 400 when create with invalid user request", func(t *testing.T) {
		// 1. setup router
		mockBookUsecase := new(mocks.BookUsecase)
		mockUserUsecase := new(mocks.UserUsecase)
		mockBorrowRecordUsecase := new(mocks.BorrowRecordUsecase)
		h := handler.NewHandler(mockBookUsecase, mockUserUsecase, mockBorrowRecordUsecase)
		router := router.NewRouter(h)
		loginPayload := dto.LoginRequest{
			Email: "aditbuddy@abc.com",
		}

		mockUserUsecase.On("LoginUser", mock.Anything, loginPayload).Return(dto.LoginResponse{}, shared.ErrInvalidRequestBody)
		expectedResp, _ := json.Marshal(dto.JsonResponse{Message: shared.ErrInvalidRequestBody.Message})

		// 2. make request
		rec := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/users/login", MakeRequestBody(loginPayload))
		router.ServeHTTP(rec, req)

		// 3. assert
		assert.Equal(t, rec.Code, 400)
		str := strings.Trim(rec.Body.String(), "\n")
		assert.Equal(t, str, string(expectedResp))
	})

	t.Run("should return status code 400 when user doesn't exist", func(t *testing.T) {
		// 1. setup router
		mockBookUsecase := new(mocks.BookUsecase)
		mockUserUsecase := new(mocks.UserUsecase)
		mockBorrowRecordUsecase := new(mocks.BorrowRecordUsecase)
		h := handler.NewHandler(mockBookUsecase, mockUserUsecase, mockBorrowRecordUsecase)
		router := router.NewRouter(h)
		loginPayload := createLoginPayload()

		mockUserUsecase.On("LoginUser", mock.Anything, loginPayload).Return(dto.LoginResponse{}, shared.ErrUserDoesntExist)
		expectedResp, _ := json.Marshal(dto.JsonResponse{Message: shared.ErrUserDoesntExist.Message})

		// 2. make request
		rec := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/users/login", MakeRequestBody(loginPayload))
		router.ServeHTTP(rec, req)

		// 3. assert
		assert.Equal(t, rec.Code, 400)
		str := strings.Trim(rec.Body.String(), "\n")
		assert.Equal(t, str, string(expectedResp))
	})

	t.Run("should return status code 500 when login failed", func(t *testing.T) {
		// 1. setup router
		mockBookUsecase := new(mocks.BookUsecase)
		mockUserUsecase := new(mocks.UserUsecase)
		mockBorrowRecordUsecase := new(mocks.BorrowRecordUsecase)
		h := handler.NewHandler(mockBookUsecase, mockUserUsecase, mockBorrowRecordUsecase)
		router := router.NewRouter(h)
		loginPayload := createLoginPayload()

		mockUserUsecase.On("LoginUser", mock.Anything, loginPayload).Return(dto.LoginResponse{}, shared.ErrFailedLogin)
		expectedResp, _ := json.Marshal(dto.JsonResponse{Message: shared.ErrFailedLogin.Message})

		// 2. make request
		rec := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/users/login", MakeRequestBody(loginPayload))
		router.ServeHTTP(rec, req)

		// 3. assert
		assert.Equal(t, rec.Code, 500)
		str := strings.Trim(rec.Body.String(), "\n")
		assert.Equal(t, str, string(expectedResp))
	})

	t.Run("should return status code 400 when login with wrong password", func(t *testing.T) {
		// 1. setup router
		mockBookUsecase := new(mocks.BookUsecase)
		mockUserUsecase := new(mocks.UserUsecase)
		mockBorrowRecordUsecase := new(mocks.BorrowRecordUsecase)
		h := handler.NewHandler(mockBookUsecase, mockUserUsecase, mockBorrowRecordUsecase)
		router := router.NewRouter(h)
		loginPayload := createLoginPayload()

		mockUserUsecase.On("LoginUser", mock.Anything, loginPayload).Return(dto.LoginResponse{}, shared.ErrInvalidPassword)
		expectedResp, _ := json.Marshal(dto.JsonResponse{Message: shared.ErrInvalidPassword.Message})

		// 2. make request
		rec := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/users/login", MakeRequestBody(loginPayload))
		router.ServeHTTP(rec, req)

		// 3. assert
		assert.Equal(t, rec.Code, 400)
		str := strings.Trim(rec.Body.String(), "\n")
		assert.Equal(t, str, string(expectedResp))
	})
}
