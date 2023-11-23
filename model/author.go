package model

import "time"

type Authors struct {
	ID        int       `gorm:"primary_key;column:id" json:"id,omitempty"`
	Name      string    `gorm:"column:author_name" json:"name"`
	CreatedAt time.Time `gorm:"column:created_at" json:"-"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"-"`
	DeletedAt time.Time `gorm:"column:deleted_at" json:"-"`
}