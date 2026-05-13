package repository

import (
	"fmt"
	"movies-project/internal/models"
)

type GenreRepository interface {
	// Exists проверяет существование жанра по ID
	Exists(id uint) (bool, error)
}

type MockGenreRepository struct{}

func (m *MockGenreRepository) Exists(id uint) (bool, error) {
	return id >= 1 && id <= 10, nil
}

func (m *MockGenreRepository) GetByID(id uint) (*models.Genre, error) {
	if id >= 1 && id <= 10 {
		return &models.Genre{ID: id, Name: "Жанр"}, nil
	}
	return nil, fmt.Errorf("жанр не найден")
}
