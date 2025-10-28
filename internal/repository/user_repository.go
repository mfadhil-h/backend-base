package repository

import (
	"backend-base/internal/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) GetAll() ([]model.User, error) {
	var users []model.User
	err := r.DB.Find(&users).Error
	return users, err
}

func (r *UserRepository) Create(user *model.User) error {
	return r.DB.Create(user).Error
}
