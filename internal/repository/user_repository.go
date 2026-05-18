package repository

import (
	"errors"
	"movies-project/internal/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUserRegister(register *models.User) error
	FindUserByID(id uint) (*models.User, error)
	Authentication(authy *models.Login) (*models.User, error)
	ExistUser(id uint) (bool, error)
}
type gormUserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &gormUserRepository{db: db}
}

func (r *gormUserRepository) CreateUserRegister(register *models.User) error {
	if register == nil {
		return errors.New(`регистр содержит нил`)
	}
	return r.db.Create(register).Error
}

func (r *gormUserRepository) FindUserByID(id uint) (*models.User, error) {
	var user models.User
	if id == 0 {
		return &user, nil
	}
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, errors.New("пользователь не найден")
	}
	return &user, nil
}

func (r *gormUserRepository) Authentication(authy *models.Login) (*models.User, error) {
	var user models.User
	query := r.db.Model(&models.User{})
	if &authy.Login != nil {
		query = query.Where("login = ?", authy.Login)
	}
	if &authy.Password != nil {
		query = query.Where("password = ?", authy.Password)
	}

	if err := query.First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (g *gormUserRepository) ExistUser(id uint) (bool, error) {
	var count int64
	err := g.db.
		Model(&models.User{}).
		Where("id = ?", id).
		Count(&count).
		Error
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
