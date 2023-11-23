package usecase

import (
	"context"
	"errors"

	"github.com/adityatresnobudi/library-api/internal/dto"
	"github.com/adityatresnobudi/library-api/internal/model"
	"github.com/adityatresnobudi/library-api/internal/repository"
	"github.com/adityatresnobudi/library-api/internal/shared"
	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	userRepo repository.UserRepository
}

type UserUsecase interface {
	GetUsers(ctx context.Context, name string) ([]dto.UserResponse, error)
	CreateUsers(ctx context.Context, user dto.UserPayload) (dto.UserResponse, error)
	LoginUser(ctx context.Context, req dto.LoginRequest) (dto.LoginResponse, error)
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}

func (uu *userUsecase) GetUsers(ctx context.Context, name string) ([]dto.UserResponse, error) {
	users := []dto.UserResponse{}
	user := dto.UserResponse{}
	uc, err := uu.userRepo.FindAll(ctx, name)
	if err != nil {
		return nil, shared.ErrGettingUsers
	}

	for _, u := range uc {
		user.ID = u.ID
		user.Name = u.Name
		user.Email = u.Email
		user.Phone = u.Phone
		users = append(users, user)
	}

	return users, nil
}

func (uu *userUsecase) CreateUsers(ctx context.Context, user dto.UserPayload) (dto.UserResponse, error) {
	newUser := model.Users{}
	userRes := dto.UserResponse{}

	newUser.ID = user.ID
	newUser.Name = user.Name
	newUser.Email = user.Email
	newUser.Phone = user.Phone

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		return dto.UserResponse{}, shared.ErrCreateUsers
	}
	newUser.Password = string(hash)

	uc, err := uu.userRepo.Create(ctx, newUser)
	if err != nil {
		return dto.UserResponse{}, shared.ErrCreateUsers
	}

	userRes.ID = uc.ID
	userRes.Name = uc.Name
	userRes.Email = uc.Email
	userRes.Phone = uc.Phone

	return userRes, nil
}

func (uu *userUsecase) LoginUser(ctx context.Context, req dto.LoginRequest) (dto.LoginResponse, error) {
	output := dto.LoginResponse{}

	user, err := uu.userRepo.FindByEmail(ctx, req.Email)
	if err != nil || user.ID == 0 {
		if errors.Is(err, shared.ErrRecordNotFound) {
			return output, shared.ErrUserDoesntExist
		}
		return output, shared.ErrFailedLogin
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return output, shared.ErrInvalidPassword
	}

	claims := shared.JWTClaims{
		ID: user.ID,
	}

	userPayload := dto.UserPayload{
		ID:       user.ID,
		Name:     user.Email,
		Email:    user.Email,
		Phone:    user.Phone,
		Password: user.Password,
	}

	token, err := shared.AuthorizedJWT(claims, userPayload)
	if err != nil {
		return output, shared.ErrFailedLogin
	}

	output = dto.LoginResponse{
		AccessToken: token,
	}

	return output, nil
}
