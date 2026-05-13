package repository

import (
	"movies-project/internal/models"

	"gorm.io/gorm"
)

type GenreRepository interface {
	GetGenres(genre *[]models.Genre) error
	CreateGenre(genre *models.Genre) error
	Exist(id uint) (bool, error)
}
type gormGenreRepository struct {
	db *gorm.DB
}

func NewGenreRepository(db *gorm.DB) GenreRepository {
	return &gormGenreRepository{db: db}
}

func (g *gormGenreRepository) GetGenres(genre *[]models.Genre) error {
	if err := g.db.Find(genre).Error; err != nil {
		return err
	}
	return nil
}

func (g *gormGenreRepository) CreateGenre(genre *models.Genre) error {
	if genre == nil {
		return nil
	}
	return g.db.Create(genre).Error
}

func (g *gormGenreRepository) Exist(id uint) (bool, error) {
	var count int64
	err := g.db.
		Model(&models.Genre{}).
		Where("id = ?", id).
		Count(&count).
		Error
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
