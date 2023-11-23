package model

import "time"

type Users struct {
	ID int `gorm:"primary_key;column:id"`
	Name string `gorm:"column:user_name"`
	Email string `gorm:"column:email"`
	Phone string `gorm:"column:phone"`
	Password string `gorm:"column:user_password"`
	CreatedAt time.Time `gorm:"column:created_at" json:"-"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"-"`
	DeletedAt time.Time `gorm:"column:deleted_at" json:"-"`
}