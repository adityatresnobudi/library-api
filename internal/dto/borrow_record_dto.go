package dto

type BorrowRecordsDTO struct {
	ID            int    `json:"id"`
	UserId        int    `json:"user_id" binding:"required"`
	BookId        int    `json:"book_id" binding:"required"`
	BorrowingDate string `json:"borrowing_date" binding:"required"`
	ReturningDate string `json:"returning_date,omitempty"`
}

type BorrowDTO struct {
	ID            int    `json:"id"`
	UserId        int    `json:"user_id"`
	BookId        int    `json:"book_id"`
	BorrowingDate string `json:"borrowing_date"`
}
