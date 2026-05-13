package services

import (
	"movies-project/internal/models"
	"movies-project/internal/repository"
)

type GenreService interface {
	GetAllGenre(g *[]models.Genre) *[]models.Genre
	CreateGenre(req models.GenreCreateRequest) (*models.Genre, error)
}

type genreService struct {
	genres repository.GenreRepository
}

func NewGenreService(genres repository.GenreRepository) GenreService {
	return &genreService{genres}
}

func (genre *genreService) GetAllGenre(g *[]models.Genre) *[]models.Genre {
	return genre.GetAllGenre(g)
}

func (g *genreService) CreateGenre(req models.GenreCreateRequest) (*models.Genre, error) {
	createGenre := &models.Genre{
		Name: req.Name,
	}
	if err := g.genres.CreateGenre(*createGenre); err != nil {
		return nil, err
	}
	return createGenre, nil
}
