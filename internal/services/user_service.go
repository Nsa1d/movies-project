package services

import (
	"errors"
	"movies-project/internal/models"
	"movies-project/internal/repository"
	"regexp"
	"strings"
	"unicode/utf8"
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

func (s *userService) CreateUser(req models.Registration) (*models.User, error) {
	errValid := s.validateRegister(req)
	if errValid != nil {
		return nil, errValid
	}
	register := &models.User{
		Login:    req.Login,
		Username: req.Username,
		Password: req.Password,
	}

	if err := s.user.CreateUserRegister(register); err != nil {
		return nil, err
	}

	return register, nil
}

func (s *userService) Authenticate(authy models.Login) (*models.User, error) {
	errValid := s.validateAuth(authy)
	if errValid != nil {
		return nil, errValid
	}
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

	pattern := `^[a-zA-Z0-9_]+$`
	loginValidation := regexp.MustCompile(pattern)
	if !loginValidation.MatchString(req.Login) {
		return errors.New("при создании логина используйте только латиницу, числа и нижнее подчеркивание")
	}

	if len(loginTrimspace) == 0 {
		return errors.New("логин обязателен")
	}
	if len(passTrimSpace) == 0 {
		return errors.New("пароль обязателен")
	}

	return nil
}

func (s *userService) validateRegister(req models.Registration) error {
	var user models.User
	login := strings.TrimSpace(req.Login)
	password := strings.TrimSpace(req.Password)
	userName := strings.TrimSpace(req.Username)

	pattern := `^[a-zA-Z0-9_]+$`
	loginValidation := regexp.MustCompile(pattern)

	const (
		minLogin    = 4
		minPassword = 8
	)

	if len(login) == 0 {
		return errors.New("пустой логин")
	}
	if len(password) == 0 {
		return errors.New("пустой пароль")
	}
	if len(userName) == 0 {
		return errors.New("пустое имя")
	}

	if req.Login == user.Login {
		return errors.New("такой логин уже существует")
	}
	if utf8.RuneCountInString(login) < minLogin {
		return errors.New("логин должен быть длиннее 4 символов")
	}
	if utf8.RuneCountInString(password) < minPassword {
		return errors.New("пароль должен быть длиннее 8 символов")
	}

	if !loginValidation.MatchString(login) {
		return errors.New("при создании логина используйте только латиницу, числа и нижнее подчеркивание")
	}

	return nil
}
