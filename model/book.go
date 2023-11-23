package model

import "time"

type Books struct {
	ID          int       `gorm:"primary_key;column:id" json:"id,omitempty"`
	Title       string    `gorm:"column:title" json:"title" binding:"required"`
	Description string    `gorm:"column:description" json:"description" binding:"required"`
	Quantity    int       `gorm:"column:quantity;check:quantity >= 0" json:"qty" binding:"required"`
	Cover       string    `gorm:"column:cover" json:"cover"`
	AuthorId    int       `gorm:"column:author_id" json:"author_id"`
	Authors     Authors   `gorm:"foreignKey:AuthorId" json:"authors"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"-"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"-"`
	DeletedAt   time.Time `gorm:"column:deleted_at" json:"-"`
}