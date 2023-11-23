package repository

import (
	"context"
	"time"

	"github.com/adityatresnobudi/library-api/internal/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type borrowRecordRepository struct {
	db *gorm.DB
}

type BorrowRecordRepository interface {
	FindBorrowRecordByID(ctx context.Context, id int) (model.BorrowRecords, error)
	CreateBorrowRecord(ctx context.Context, borrow model.BorrowRecords) (model.BorrowRecords, error)
	UpdateBorrowRecord(ctx context.Context, borrow model.BorrowRecords) (model.BorrowRecords, error)
}

func NewBorrowRecordRepository(db *gorm.DB) BorrowRecordRepository {
	return &borrowRecordRepository{
		db: db,
	}
}

func (rr *borrowRecordRepository) FindBorrowRecordByID(ctx context.Context, id int) (model.BorrowRecords, error) {
	borrowRecord := model.BorrowRecords{}

	
	err := rr.db.WithContext(ctx).
	Model(&model.BorrowRecords{}).
	Where("id = ?", id).
	First(&borrowRecord).Error
	if err != nil {
		return model.BorrowRecords{}, err
	}
	return borrowRecord, nil
}

func (rr *borrowRecordRepository) CreateBorrowRecord(ctx context.Context, borrowRecord model.BorrowRecords) (model.BorrowRecords, error) {
	err := rr.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		book := model.Books{}
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&book, borrowRecord.BookId).Error; err != nil {
			return err
		}
		if err := tx.Model(&book).Update("quantity", gorm.Expr("quantity - 1")).Error; err != nil {
			return err
		}
		if err := tx.Create(&borrowRecord).Error; err != nil {
			return err
		}
		return nil
	})
	return borrowRecord, err
}

func (rr *borrowRecordRepository) UpdateBorrowRecord(ctx context.Context, borrowRecord model.BorrowRecords) (model.BorrowRecords, error) {
	// rr.db.WithContext(ctx).Exec("SELECT pg_sleep(6)")
	err := rr.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		book := model.Books{}
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&book, borrowRecord.BookId).Error; err != nil {
			return err
		}
		if err := tx.Model(&borrowRecord).Update("returning_date", time.Now()).Error; err != nil {
			return err
		}
		if err := tx.Model(&book).Update("quantity", gorm.Expr("quantity + 1")).Error; err != nil {
			return err
		}
		return nil
	})
	return borrowRecord, err
}
