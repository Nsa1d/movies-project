package services

import (
	"errors"
	"movies-project/internal/models"
	"movies-project/internal/repository"
)

type UserService interface {
	GetUserByID(id uint) (*models.User, error)
	CreateUser(reg models.Registration) (*models.User, error)
	Authenticate(authy models.Login) (*models.User, error)
}

type userService struct {
	user repository.UserRepository
}

func NewUserService(user repository.UserRepository) UserService {
	return &userService{user: user}
}

func (s *userService) GetUserByID(id uint) (*models.User, error) {
	return s.user.FindUserByID(id)
}

func (s *userService) CreateUser(reg models.Registration) (*models.User, error) {
	var user models.User

	if reg.Username == "" {
		return nil, errors.New("пустое имя")
	}
	if reg.Password == "" {
		return nil, errors.New("пустой пароль")
	}
	if reg.Login == "" {
		return nil, errors.New("пустой логин")
	}

	if reg.Login == user.Login {
		return nil, errors.New("login is already taken")
	}

	register := &models.User{
		Login:    reg.Login,
		Username: reg.Username,
		Password: reg.Password,
	}

	if err := s.user.CreateUserRegister(register); err != nil {
		return nil, err
	}

	return register, nil
}

func (s *userService) Authenticate(authy models.Login) (*models.User, error) {
	if authy.Login == "" {
		return nil, errors.New("логин пустой")
	}
	if authy.Password == "" {
		return nil, errors.New("пароль пустой")
	}

	login := &models.Login{
		Login:    authy.Login,
		Password: authy.Password,
	}
	return s.user.Authentication(login)
}
