package handler

import (
	"context"
	"log"

	"github.com/adityatresnobudi/library-api/internal/dto"
	"github.com/adityatresnobudi/library-api/proto/pb"
	"github.com/adityatresnobudi/library-api/internal/usecase"
)

const TimeNull = "0001-01-01 00:00:00"

type BorrowServer struct {
	pb.UnimplementedBorrowServer
	borrowRecordUsecase usecase.BorrowRecordUsecase
}

func NewBorrowHandler(borrowRecordUsecase usecase.BorrowRecordUsecase) *BorrowServer {
	return &BorrowServer{
		borrowRecordUsecase: borrowRecordUsecase,
	}
}

func (b *BorrowServer) Borrow(ctx context.Context, req *pb.BorrowRequest) (*pb.BorrowResponse, error) {
	userId := ctx.Value("id")

	record := dto.BorrowRecordsDTO{
		UserId:        int(req.UserId),
		BookId:        int(req.BookId),
		BorrowingDate: req.BorrowingDate,
	}

	borrowRecord, err := b.borrowRecordUsecase.BorrowBook(ctx, record, userId.(int))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &pb.BorrowResponse{UserId: int64(borrowRecord.UserId), BookId: int64(borrowRecord.BookId), BorrowingDate: borrowRecord.BorrowingDate}, nil
}

func (b *BorrowServer) Return(ctx context.Context, req *pb.ReturnRequest) (*pb.ReturnResponse, error) {
	userId := ctx.Value("id")

	brUpdate, err := b.borrowRecordUsecase.GetBorrowRecordByID(ctx, int(req.BorrowRecordId))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if brUpdate.ReturningDate != TimeNull {
		return nil, err
	}

	brUpdate.ID = int(req.BorrowRecordId)
	output, err := b.borrowRecordUsecase.ReturnBook(ctx, brUpdate, userId.(int))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &pb.ReturnResponse{UserId: int64(output.UserId), BookId: int64(output.BookId), ReturningDate: output.ReturningDate}, nil
}
