package service

import (
	"go-fiber-crud/app/config/logs"
	"go-fiber-crud/app/model"
	"go-fiber-crud/app/repository"
)

type authService struct {
	repository repository.AuthRepository
}
type AuthService interface {
	Login(login *model.Login) (*model.LoginResponse, error)
	Register(register *model.Register) (*model.UserResponse, error)
}

func NewAuthService(repository repository.AuthRepository) authService {
	return authService{repository: repository}
}

func (s authService) Login(login *model.Login) (*model.LoginResponse, error) {
	response, err := s.repository.Login(login)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	return response, nil
}
func (s authService) Register(register *model.Register) (*model.UserResponse, error) {
	response, err := s.repository.Register(register)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	return response, nil
}
