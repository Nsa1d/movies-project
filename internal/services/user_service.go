package services

import (
	"errors"
	"movies-project/internal/models"
	"movies-project/internal/repository"
	"strings"
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
	return s.user.GetByID(id)
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
	//errValid := s.validateAuth(authy)
	//if errValid != nil {
	//	return nil, errValid
	//}
	user, err := s.user.GetByLogin(authy.Login)
	if err != nil {
		return nil, err
	}
	if user.Password != authy.Password {
		return nil, errors.New("неверный пароль")
	}

	return user, nil
}

func (s *userService) validateAuth(req models.Login) error {
	loginTrimspace := strings.TrimSpace(req.Login)
	passTrimSpace := strings.TrimSpace(req.Password)

	if len(loginTrimspace) > 0 {
		return errors.New("логин обязателен")
	}

	if len(passTrimSpace) > 0 {
		return errors.New("пароль обязателен")
	}

	return nil
}
