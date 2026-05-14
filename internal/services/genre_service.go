package services

import (
	"errors"
	"movies-project/internal/models"
	"movies-project/internal/repository"
	"strings"
	"unicode"
)

type GenreService interface {
	GetAllGenres() ([]models.Genre, error)
	CreateGenre(req models.GenreCreateRequest) (*models.Genre, error)
}

type genreService struct {
	genres repository.GenreRepository
}

func NewGenreService(genres repository.GenreRepository) GenreService {
	return &genreService{genres: genres}
}

func (s *genreService) GetAllGenres() ([]models.Genre, error) {
	return s.genres.GetGenres()
}

func (s *genreService) CreateGenre(req models.GenreCreateRequest) (*models.Genre, error) {
	var genre []models.Genre
	trimmed := strings.TrimSpace(req.Name)

	if trimmed == "" {
		return nil, errors.New("name is required")
	}

	for _, g := range genre {
		if g.Name == req.Name {
			return nil, errors.New("name is duplicated")
		}
	}

	for _, ch := range trimmed {
		if unicode.IsDigit(ch) {
			return nil, errors.New("name is invalid")
		}
	}

	createGenre := &models.Genre{
		Name: trimmed,
	}

	if err := s.genres.CreateGenre(createGenre); err != nil {
		return nil, err
	}
	return createGenre, nil
}
