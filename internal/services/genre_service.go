package services

import (
	"errors"
	"movies-project/internal/models"
	"movies-project/internal/repository"
	"regexp"
	"strings"
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
	errValid := s.genreValidate(req)
	if errValid != nil {
		return nil, errValid
	}

	createGenre := &models.Genre{
		Name: req.Name,
	}

	if err := s.genres.CreateGenre(createGenre); err != nil {
		return nil, err
	}
	return createGenre, nil
}

func (s *genreService) genreValidate(req models.GenreCreateRequest) error {
	var genre models.Genre
	pattern := `[а-яА-Я]+$`
	genreValidate := regexp.MustCompile(pattern)
	genreTrimmed := strings.TrimSpace(req.Name)

	if genreTrimmed == "" {
		return errors.New("имя пустое")
	}

	if !genreValidate.MatchString(genreTrimmed) {
		return errors.New("название жанра должно быть только на русском")
	}

	if genreTrimmed == genre.Name {
		return errors.New("такой жанр уже есть")
	}

	return nil
}
