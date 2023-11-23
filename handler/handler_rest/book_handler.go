package handlerrest

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/adityatresnobudi/library-api/dto"
	"github.com/adityatresnobudi/library-api/shared"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetBooks(c *gin.Context) {
	ctx := c.Request.Context()
	title := c.Query("title")
	books, err := h.BookUsecase.GetBooks(ctx, title)
	if err != nil {
		log.Println(err)
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dto.JsonResponse{Data: books})
}

func (h *Handler) AddBooks(c *gin.Context) {
	ctx := c.Request.Context()
	newBook := dto.BookPayload{}
	if err := c.ShouldBindJSON(&newBook); err != nil {
		log.Println(err)
		c.Error(shared.ErrInvalidRequestBody)
		return
	}

	newBook.Title = strings.TrimSpace(newBook.Title)

	books, err := h.BookUsecase.AddBooks(ctx, newBook)
	if err != nil {
		log.Println(err)
		c.Error(err)
		return
	}
	message := fmt.Sprintf("successfully add new book with id %d", books.ID)
	c.JSON(http.StatusCreated, dto.JsonResponse{Message: message, Data: books})
}
