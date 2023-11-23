package repository

import (
	"context"
	"errors"

	"github.com/adityatresnobudi/library-api/internal/model"
	"github.com/adityatresnobudi/library-api/internal/shared"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

type UserRepository interface {
	FindAll(ctx context.Context, name string) ([]model.Users, error)
	Create(ctx context.Context, user model.Users) (model.Users, error)
	FindByEmail(ctx context.Context, email string) (model.Users, error)
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) FindAll(ctx context.Context, name string) ([]model.Users, error) {
	users := []model.Users{}

	err := u.db.WithContext(ctx).
		Model(&model.Users{}).
		Where("user_name ILIKE ?", "%"+name+"%").
		Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *userRepository) Create(ctx context.Context, user model.Users) (model.Users, error) {
	err := u.db.WithContext(ctx).Create(&user).Error
	if err != nil {
		return model.Users{}, err
	}

	return user, nil
}

func (u *userRepository) FindByEmail(ctx context.Context, email string) (model.Users, error) {
	user := model.Users{}

	err := u.db.WithContext(ctx).
		Model(&model.Users{}).
		Where("email = ?", email).
		First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Users{}, shared.ErrRecordNotFound
		}
		return model.Users{}, err
	}

	return user, nil
}
