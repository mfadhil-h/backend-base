package service

import (
	"backend-base/internal/model"
	"backend-base/internal/repository"
)

type UserService struct {
	Repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) GetUsers() ([]model.User, error) {
	return s.Repo.GetAll()
}

func (s *UserService) CreateUser(user *model.User) error {
	return s.Repo.Create(user)
}
