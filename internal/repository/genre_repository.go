package repository

import (
	"movies-project/internal/models"

	"gorm.io/gorm"
)

type GenreRepository interface {
	GetAllGenre(g *[]models.Genre) error
	CreateGenre(g models.Genre) error
	Exist(id uint) (bool, error)
}
type gormGenreRepository struct {
	db *gorm.DB
}

func NewGenreRepository(db *gorm.DB) GenreRepository {
	return &gormGenreRepository{db: db}
}

func (genre *gormGenreRepository) GetAllGenre(g *[]models.Genre) error {
	if err := genre.db.Find(g).Error; err != nil {
		return err
	}
	return nil
}

func (genre *gormGenreRepository) CreateGenre(g models.Genre) error {
	if err := genre.db.Create(&g).Error; err != nil {
		return nil
	}
	return nil
}

func (genre *gormGenreRepository) Exist(id uint) (bool, error) {
	var count int64
	err := genre.db.
		Model(&models.Genre{}).
		Where("id = ?", id).
		Count(&count).
		Error
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
