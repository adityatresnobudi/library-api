package handlerrest

import "github.com/adityatresnobudi/library-api/internal/usecase"

type Handler struct {
	BookUsecase         usecase.BookUsecase
	UserUsecase         usecase.UserUsecase
	BorrowRecordUsecase usecase.BorrowRecordUsecase
}

func NewHandler(BookUsecase usecase.BookUsecase, UserUsecase usecase.UserUsecase, BorrowRecordUsecase usecase.BorrowRecordUsecase) *Handler {
	return &Handler{
		BookUsecase:         BookUsecase,
		UserUsecase:         UserUsecase,
		BorrowRecordUsecase: BorrowRecordUsecase,
	}
}
