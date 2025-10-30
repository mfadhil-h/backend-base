package service

import (
	"backend-base/internal/model"
	"backend-base/internal/repository"
	"backend-base/internal/util"
	"errors"
)

type AuthService struct {
	UserRepo *repository.UserRepository
}

func NewAuthService(repo *repository.UserRepository) *AuthService {
	return &AuthService{UserRepo: repo}
}

func (s *AuthService) RegisterUser(name, email, password string) (*model.User, error) {
	existing, _ := s.UserRepo.FindByEmail(email)
	if existing != nil {
		return nil, errors.New("email already registered")
	}

	hashed, err := util.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		Name:     name,
		Email:    email,
		Password: hashed,
	}

	if err := s.UserRepo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *AuthService) Login(email, password string) (string, error) {
	user, err := s.UserRepo.FindByEmail(email)
	if err != nil || user == nil {
		return "", errors.New("invalid email or password")
	}

	if !util.CheckPasswordHash(password, user.Password) {
		return "", errors.New("invalid email or password")
	}

	token, err := util.GenerateJWT(user.ID, user.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}
