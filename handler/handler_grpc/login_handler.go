package handler

import (
	"context"
	"log"
	"time"

	"github.com/adityatresnobudi/library-api/dto"
	"github.com/adityatresnobudi/library-api/proto/pb"
	"github.com/adityatresnobudi/library-api/shared"
	"github.com/adityatresnobudi/library-api/usecase"
)

type AuthServer struct {
	pb.UnimplementedAuthServer
	userUsecase usecase.UserUsecase
}

func NewLoginHandler(userUsecase usecase.UserUsecase) *AuthServer {
	return &AuthServer{
		userUsecase: userUsecase,
	}
}

func (a *AuthServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	time.Sleep(5 * time.Second)
	LoginRequest := dto.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	}

	response, err := a.userUsecase.LoginUser(ctx, LoginRequest)
	if err != nil {
		log.Println(err)
		return nil, shared.ErrFailedLogin
	}

	if response.AccessToken == "" {
		msg := "login failed"
		return &pb.LoginResponse{Message: msg}, nil
	}
	msg := "login success"
	return &pb.LoginResponse{Message: msg, Token: response.AccessToken}, nil
}
