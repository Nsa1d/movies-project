package repository

import (
	"errors"
	"movies-project/internal/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUserRegister(register *models.User) error
	GetByID(id uint) (*models.User, error)
	GetByLogin(login string) (*models.User, error)
	Exist(id uint) (bool, error)
}
type gormUserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &gormUserRepository{db: db}
}

func (r *gormUserRepository) CreateUserRegister(register *models.User) error {
	if register == nil {
		return nil
	}
	return r.db.Create(register).Error
}

func (r *gormUserRepository) GetByID(id uint) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, errors.New("пользователь не найден")
	}
	return &user, nil
}

func (r *gormUserRepository) GetByLogin(login string) (*models.User, error) {
	var user models.User
	err := r.db.Where("login = ?", login).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (g *gormUserRepository) Exist(id uint) (bool, error) {
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
