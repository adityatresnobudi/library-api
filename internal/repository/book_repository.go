package repository

import (
	"context"

	"github.com/adityatresnobudi/library-api/internal/model"
	"gorm.io/gorm"
)

type bookRepository struct {
	db *gorm.DB
}

type BookRepository interface {
	FindAllBooks(ctx context.Context, title string) ([]model.Books, error)
	AddOne(ctx context.Context, book model.Books) (model.Books, error)
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{
		db: db,
	}
}

func (b *bookRepository) FindAllBooks(ctx context.Context, title string) ([]model.Books, error) {
	books := []model.Books{}

	err := b.db.WithContext(ctx).
		Model(&model.Books{}).
		Joins("Authors").
		Where("title ILIKE ?", "%"+title+"%").
		Find(&books).Error
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (b *bookRepository) AddOne(ctx context.Context, book model.Books) (model.Books, error) {
	err := b.db.WithContext(ctx).Create(&book).Error
	if err != nil {
		return model.Books{}, err
	}

	return book, nil
}
