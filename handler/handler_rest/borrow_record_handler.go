package handlerrest

import (
	"log"
	"net/http"
	"strconv"

	"github.com/adityatresnobudi/library-api/dto"
	"github.com/adityatresnobudi/library-api/shared"
	"github.com/gin-gonic/gin"
)

const TimeNull = "0001-01-01 00:00:00"

func (h *Handler) BorrowBook(c *gin.Context) {
	ctx := c.Request.Context()
	userId := c.GetInt("id")

	record := dto.BorrowRecordsDTO{}
	if err := c.ShouldBindJSON(&record); err != nil {
		log.Println(err)
		c.Error(shared.ErrInvalidRequestBody)
		return
	}

	borrowRecord, err := h.BorrowRecordUsecase.BorrowBook(ctx, record, userId)
	if err != nil {
		log.Println(err)
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dto.JsonResponse{Data: borrowRecord})
}

func (h *Handler) ReturnBook(c *gin.Context) {
	ctx := c.Request.Context()
	userId := c.GetInt("id")

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.Error(shared.ErrIdNotFound)
		return
	}

	brUpdate, err := h.BorrowRecordUsecase.GetBorrowRecordByID(ctx, id)
	if err != nil {
		log.Println(err)
		c.Error(err)
		return
	}

	if brUpdate.ReturningDate != TimeNull {
		c.Error(shared.ErrAlreadyReturned)
		return
	}

	brUpdate.ID = id
	output, err := h.BorrowRecordUsecase.ReturnBook(ctx, brUpdate, userId)
	if err != nil {
		log.Println(err)
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dto.JsonResponse{Data: output})
}
