package usecase

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/adityatresnobudi/library-api/dto"
	"github.com/adityatresnobudi/library-api/model"
	"github.com/adityatresnobudi/library-api/repository"
	"github.com/adityatresnobudi/library-api/shared"
)

type borrowRecordUsecase struct {
	borrowRecordRepo repository.BorrowRecordRepository
}

type BorrowRecordUsecase interface {
	GetBorrowRecordByID(ctx context.Context, id int) (dto.BorrowRecordsDTO, error)
	BorrowBook(ctx context.Context, record dto.BorrowRecordsDTO, requestId int) (dto.BorrowRecordsDTO, error)
	ReturnBook(ctx context.Context, record dto.BorrowRecordsDTO, requestId int) (dto.BorrowRecordsDTO, error)
}

func NewBorrowRecordUsecase(borrowRecordRepo repository.BorrowRecordRepository) BorrowRecordUsecase {
	return &borrowRecordUsecase{
		borrowRecordRepo: borrowRecordRepo,
	}
}

func (ru *borrowRecordUsecase) GetBorrowRecordByID(ctx context.Context, id int) (dto.BorrowRecordsDTO, error) {
	borrowRecord := dto.BorrowRecordsDTO{}
	uc, err := ru.borrowRecordRepo.FindBorrowRecordByID(ctx, id)
	if err != nil {
		return dto.BorrowRecordsDTO{}, shared.ErrGettingBorrowRecords
	}

	borrowRecord.ID = uc.ID
	borrowRecord.UserId = uc.UserId
	borrowRecord.BookId = uc.BookId
	borrowRecord.BorrowingDate = TimeToStrConv(uc.BorrowingDate)
	borrowRecord.ReturningDate = TimeToStrConv(uc.ReturningDate.Time)

	return borrowRecord, nil
}

func (ru *borrowRecordUsecase) BorrowBook(ctx context.Context, record dto.BorrowRecordsDTO, requestId int) (dto.BorrowRecordsDTO, error) {
	if record.UserId != requestId {
		log.Println(record.UserId)
		log.Println(requestId)
		return dto.BorrowRecordsDTO{}, shared.ErrUnauthorized
	}

	borrowRecord := model.BorrowRecords{
		ID:            record.ID,
		UserId:        record.UserId,
		BookId:        record.BookId,
		BorrowingDate: StrToTimeConv(record.BorrowingDate),
		ReturningDate: sql.NullTime{},
	}

	uc, err := ru.borrowRecordRepo.CreateBorrowRecord(ctx, borrowRecord)
	if err != nil {
		return dto.BorrowRecordsDTO{}, shared.ErrCreateBorrowRecord
	}

	dto := dto.BorrowRecordsDTO{
		ID:            uc.ID,
		UserId:        uc.UserId,
		BookId:        uc.BookId,
		BorrowingDate: TimeToStrConv(uc.BorrowingDate),
		ReturningDate: "",
	}

	return dto, nil
}

func (ru *borrowRecordUsecase) ReturnBook(ctx context.Context, record dto.BorrowRecordsDTO, requestId int) (dto.BorrowRecordsDTO, error) {
	if record.UserId != requestId {
		return dto.BorrowRecordsDTO{}, shared.ErrUnauthorized
	}

	borrowRecord := model.BorrowRecords{
		ID:            record.ID,
		UserId:        record.UserId,
		BookId:        record.BookId,
		BorrowingDate: StrToTimeConv(record.BorrowingDate),
		ReturningDate: sql.NullTime{},
	}

	uc, err := ru.borrowRecordRepo.UpdateBorrowRecord(ctx, borrowRecord)
	if err != nil {
		return dto.BorrowRecordsDTO{}, shared.ErrUpdateBorrowRecord
	}

	dto := dto.BorrowRecordsDTO{
		ID:            uc.ID,
		UserId:        uc.UserId,
		BookId:        uc.BookId,
		BorrowingDate: TimeToStrConv(uc.BorrowingDate),
		ReturningDate: TimeToStrConv(uc.ReturningDate.Time),
	}

	return dto, nil
}

func TimeToStrConv(dateTime time.Time) string {
	return dateTime.Format("2006-01-02 15:04:05")
}

func StrToTimeConv(dateString string) time.Time {
	layoutFormat := "2006-01-02 15:04:05"
	date, _ := time.Parse(layoutFormat, dateString)
	return date
}
