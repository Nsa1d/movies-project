package services

import (
	"errors"
	"movies-project/internal/models"
	"movies-project/internal/repository"
	"strings"
	"unicode"
)

type GenreService interface {
	GetAllGenres(genres *[]models.Genre) error
	CreateGenre(req models.GenreCreateRequest) (*models.Genre, error)
}

type genreService struct {
	genreRepo repository.GenreRepository
}

func NewGenreService(genreRepo repository.GenreRepository) GenreService {
	return &genreService{genreRepo: genreRepo}
}

func (g *genreService) GetAllGenres(genres *[]models.Genre) error {
	return g.genreRepo.GetGenres(genres)
}

func (g *genreService) CreateGenre(req models.GenreCreateRequest) (*models.Genre, error) {
	var genre []models.Genre
	trimmed := strings.TrimSpace(req.Name)

	if req.Name == trimmed {
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
		Name: req.Name,
	}

	if err := g.genreRepo.CreateGenre(createGenre); err != nil {
		return nil, err
	}
	return createGenre, nil
}
