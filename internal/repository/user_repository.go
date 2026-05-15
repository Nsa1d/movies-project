package repository

import (
	"errors"
	"movies-project/internal/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUserRegister(register *models.User) error
	FindUserByID(id uint) (*models.User, error)
	//LoginWhithPassword(login string, password string) error
}
type gormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) UserRepository {
	return &gormUserRepository{db: db}
}

func (r *gormUserRepository) CreateUserRegister(register *models.User) error {
	if register == nil {
		return errors.New(`Регистр содержит нил`)
	}
	return r.db.Create(register).Error
}

func (r *gormUserRepository) FindUserByID(id uint) (*models.User, error) {
	var user models.User
	if id == 0 {
		return &user, nil
	}
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, errors.New("Пользователь не найден")
	}
	return &user, nil
}
