package repository

import (
	"movies-project/internal/models"

	"gorm.io/gorm"
)

type GenreRepository interface {
	GetAllGenre(g *[]models.Genre) error
	CreateGenre(g models.Genre) (*models.Genre, error)
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

func (genre *gormGenreRepository) CreateGenre(g models.Genre) (*models.Genre, error) {
	if err := genre.db.Create(&g).Error; err != nil {
		return nil, err
	}
	return &g, nil
}
