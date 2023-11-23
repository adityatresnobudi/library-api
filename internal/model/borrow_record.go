package model

import (
	"database/sql"
	"time"
)

type BorrowRecords struct {
	ID            int          `gorm:"primary_key;column:id"`
	UserId        int          `gorm:"column:user_id"`
	BookId        int          `gorm:"column:book_id"`
	BorrowingDate time.Time    `gorm:"column:borrowing_date"`
	ReturningDate sql.NullTime `gorm:"column:returning_date"`
	CreatedAt     time.Time    `gorm:"column:created_at" json:"-"`
	UpdatedAt     time.Time    `gorm:"column:updated_at" json:"-"`
	DeletedAt     time.Time    `gorm:"column:deleted_at" json:"-"`
}
