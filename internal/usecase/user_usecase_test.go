package usecase_test

import (
	"context"
	"testing"

	"github.com/adityatresnobudi/library-api/internal/dto"
	"github.com/adityatresnobudi/library-api/internal/mocks"
	"github.com/adityatresnobudi/library-api/internal/model"
	"github.com/adityatresnobudi/library-api/internal/shared"
	"github.com/adityatresnobudi/library-api/internal/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

func createModelUser() model.Users {
	return model.Users{
		ID:       1,
		Name:     "aditbuddy",
		Email:    "aditbuddy@abc.com",
		Phone:    "081290882428",
		Password: "relativiteam",
	}
}

func createUserPayload() dto.UserPayload {
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

func TestUserUsecase_GetUsers(t *testing.T) {
	t.Run("should return dto response when all user fetched", func(t *testing.T) {
		// Given
		mockUserRepository := new(mocks.UserRepository)
		uu := usecase.NewUserUsecase(mockUserRepository)

		users := []model.Users{
			createModelUser(),
		}

		userDTO := []dto.UserResponse{
			createUserResponse(),
		}

		ctx := context.Background()

		// When
		mockUserRepository.On("FindAll", ctx, "").Return(users, nil)
		output, err := uu.GetUsers(ctx, "")

		// Then
		assert.Equal(t, userDTO, output)
		assert.Nil(t, err)
	})

	t.Run("should return dto response when specific user fetched", func(t *testing.T) {
		// Given
		mockUserRepository := new(mocks.UserRepository)
		uu := usecase.NewUserUsecase(mockUserRepository)

		users := []model.Users{
			createModelUser(),
		}

		userDTO := []dto.UserResponse{
			createUserResponse(),
		}

		ctx := context.Background()

		// When
		mockUserRepository.On("FindAll", ctx, "aditbuddy").Return(users, nil)
		output, err := uu.GetUsers(ctx, "aditbuddy")

		// Then
		assert.Equal(t, userDTO, output)
		assert.Nil(t, err)
	})

	t.Run("should return dto response when specific user fetched", func(t *testing.T) {
		// Given
		mockUserRepository := new(mocks.UserRepository)
		uu := usecase.NewUserUsecase(mockUserRepository)

		users := []model.Users{}

		userDTO := []dto.UserResponse{}

		ctx := context.Background()

		// When
		mockUserRepository.On("FindAll", ctx, "natnat").Return(users, nil)
		output, err := uu.GetUsers(ctx, "natnat")

		// Then
		assert.Equal(t, userDTO, output)
		assert.Nil(t, err)
	})

	t.Run("should return error getting users when users fetch failed", func(t *testing.T) {
		// Given
		mockUserRepository := new(mocks.UserRepository)
		uu := usecase.NewUserUsecase(mockUserRepository)

		ctx := context.Background()

		// When
		mockUserRepository.On("FindAll", ctx, "").Return(nil, shared.ErrGettingUsers)
		output, err := uu.GetUsers(ctx, "")

		// Then
		assert.Nil(t, output)
		assert.ErrorIs(t, shared.ErrGettingUsers, err)
	})
}

func TestUserUsecase_CreateUsers(t *testing.T) {
	t.Run("should return dto user response when create user successful", func(t *testing.T) {
		// Given
		mockUserRepository := new(mocks.UserRepository)
		uu := usecase.NewUserUsecase(mockUserRepository)

		user := createModelUser()
		hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
		user.Password = string(hash)

		userPayload := createUserPayload()
		userResponse := createUserResponse()

		ctx := context.Background()

		// When
		mockUserRepository.On("Create", ctx, mock.AnythingOfType("model.Users")).Return(user, nil)
		output, err := uu.CreateUsers(ctx, userPayload)

		// Then
		assert.Equal(t, userResponse, output)
		assert.Nil(t, err)
	})

	t.Run("should return error when create user failed", func(t *testing.T) {
		// Given
		mockUserRepository := new(mocks.UserRepository)
		uu := usecase.NewUserUsecase(mockUserRepository)

		userPayload := createUserPayload()
		userResponse := dto.UserResponse{}

		ctx := context.Background()

		// When
		mockUserRepository.On("Create", ctx, mock.AnythingOfType("model.Users")).Return(model.Users{}, shared.ErrCreateUsers)
		output, err := uu.CreateUsers(ctx, userPayload)

		// Then
		assert.Equal(t, userResponse, output)
		assert.ErrorIs(t, shared.ErrCreateUsers, err)
	})
}

func TestUserUsecase_LoginUsers(t *testing.T) {
	t.Run("should return dto user response when create user successful", func(t *testing.T) {
		// Given
		mockUserRepository := new(mocks.UserRepository)
		uu := usecase.NewUserUsecase(mockUserRepository)

		user := createModelUser()
		hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
		user.Password = string(hash)

		userPayload := createUserPayload()
		userResponse := createUserResponse()

		ctx := context.Background()

		// When
		mockUserRepository.On("Create", ctx, mock.AnythingOfType("model.Users")).Return(user, nil)
		output, err := uu.CreateUsers(ctx, userPayload)

		// Then
		assert.Equal(t, userResponse, output)
		assert.Nil(t, err)
	})

	t.Run("should return error when create user failed", func(t *testing.T) {
		// Given
		mockUserRepository := new(mocks.UserRepository)
		uu := usecase.NewUserUsecase(mockUserRepository)

		userPayload := createUserPayload()
		userResponse := dto.UserResponse{}

		ctx := context.Background()

		// When
		mockUserRepository.On("Create", ctx, mock.AnythingOfType("model.Users")).Return(model.Users{}, shared.ErrCreateUsers)
		output, err := uu.CreateUsers(ctx, userPayload)

		// Then
		assert.Equal(t, userResponse, output)
		assert.ErrorIs(t, shared.ErrCreateUsers, err)
	})
}
