package services

import (
	"errors"
	"movies-project/internal/models"
	"movies-project/internal/repository"
)

type UserService interface {
	GetUserByID(id uint) (*models.User, error)
	CreatUser(reg *models.Registration) (*models.User, error)
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

func (s *userService) CreatUser(reg models.Registration) (*models.User, error) {
	var user models.User
	if reg.Login == user.Login {
		return errors.New("login is already taken")
	}

	register := &models.User{
		Login:    reg.Login,
		Name:     reg.Name,
		Password: reg.Password,
	}

	if err := s.user.CreateUserRegister(register); err != nil {
		return nil, err
	}

	return &register, nil
}
