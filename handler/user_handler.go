package handler

import (
	"log"
	"net/http"

	"github.com/adityatresnobudi/library-api/dto"
	"github.com/adityatresnobudi/library-api/shared"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetUsers(c *gin.Context) {
	ctx := c.Request.Context()
	name := c.Query("name")
	users, err := h.UserUsecase.GetUsers(ctx, name)
	if err != nil {
		log.Println(err)
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dto.JsonResponse{Data: users})
}

func (h *Handler) CreateUser(c *gin.Context) {
	ctx := c.Request.Context()
	newUser := dto.UserPayload{}
	if err := c.ShouldBindJSON(&newUser); err != nil {
		log.Println(err)
		c.Error(shared.ErrInvalidRequestBody)
		return
	}

	users, err := h.UserUsecase.CreateUsers(ctx, newUser)
	if err != nil {
		log.Println(err)
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, dto.JsonResponse{Data: users})
}

func (h *Handler) LoginUser(c *gin.Context) {
	ctx := c.Request.Context()
	req := dto.LoginRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println(err)
		c.Error(shared.ErrInvalidRequestBody)
		return
	}

	response, err := h.UserUsecase.LoginUser(ctx, req)
	if err != nil {
		log.Println(err)
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response)
}
