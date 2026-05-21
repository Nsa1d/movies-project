package services

import (
	"errors"
	"movies-project/internal/models"
	"movies-project/internal/repository"
	"strings"
	"time"
	"unicode/utf8"

	"gorm.io/gorm"
)

var ErrMovieNotFound = errors.New("фильм не найден")

type MovieService interface {
	CreateMovie(req models.MovieUpsertRequest) (*models.Movie, error)

	GetMovieByID(id uint) (*models.Movie, error)

	ListMovies(filter repository.MovieFilter) ([]models.Movie, error)

	UpdateMovie(id uint, req models.MovieUpsertRequest) (*models.Movie, error)

	DeleteMovie(id uint) error
}

type movieService struct {
	movies repository.MovieRepository
	genres repository.GenreRepository
}

func NewMovieService(movies repository.MovieRepository, genres repository.GenreRepository) MovieService {
	return &movieService{
		movies: movies,
		genres: genres,
	}
}

func (s *movieService) CreateMovie(req models.MovieUpsertRequest) (*models.Movie, error) {
	if err := s.validateMovie(req); err != nil {
		return nil, err
	}

	movie := &models.Movie{
		Title:       strings.TrimSpace(req.Title),
		Year:        req.Year,
		DurationMin: req.DurationMin,
		Rating:      0,
		RatingCount: 0,
		Country:     strings.TrimSpace(req.Country),
		Description: strings.TrimSpace(req.Description),
		GenreID:     req.GenreID,
	}

	if err := s.movies.Create(movie); err != nil {
		return nil, err
	}

	return movie, nil
}

func (s *movieService) ListMovies(filter repository.MovieFilter) ([]models.Movie, error) {
	return s.movies.List(filter)
}

func (s *movieService) GetMovieByID(id uint) (*models.Movie, error) {
	movie, err := s.movies.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrMovieNotFound
		}
		return nil, err
	}

	return movie, nil
}

func (s *movieService) UpdateMovie(id uint, req models.MovieUpsertRequest) (*models.Movie, error) {
	movie, err := s.movies.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrMovieNotFound
		}
		return nil, err
	}

	if err := s.validateMovie(req); err != nil {
		return nil, err
	}

	movie.Title = strings.TrimSpace(req.Title)
	movie.Year = req.Year
	movie.DurationMin = req.DurationMin
	movie.Rating = movie.Rating
	movie.Country = strings.TrimSpace(req.Country)
	movie.Description = strings.TrimSpace(req.Description)
	movie.GenreID = req.GenreID

	if err := s.movies.Update(movie); err != nil {
		return nil, err
	}

	return movie, nil
}

func (s *movieService) DeleteMovie(id uint) error {
	if _, err := s.movies.GetByID(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		return err
	}
	return s.movies.Delete(id)
}

func (s *movieService) validateMovie(req models.MovieUpsertRequest) error {
	found, err := s.movies.GetByTitle(strings.TrimSpace(req.Title))
	if err != nil {
		return err
	}
	if found != nil {
		return errors.New("фильм с таким названием уже существует")
	}

	if strings.TrimSpace(req.Title) == "" {
		return errors.New("поле title не должно быть пустым")
	}

	minDate, maxDate := time.Date(1888, time.January, 1, 0, 0, 0, 0, time.UTC), time.Now().UTC()
	releaseDate := time.Date(req.Year, time.January, 1, 0, 0, 0, 0, time.UTC)
	if releaseDate.Before(minDate) || releaseDate.After(maxDate) {
		return errors.New("дата выпуска должна быть между 1888-01-01 и текущей датой")
	}

	if req.DurationMin <= 0 || req.DurationMin > 420 {
		return errors.New("недопустимая длительность фильма")
	}

	if req.Rating < 1 || req.Rating > 10 {
		return errors.New("рейтинг должен быть от 1 до 10")
	}

	if strings.TrimSpace(req.Country) == "" {
		return errors.New("поле country не должно быть пустым")
	}

	trimmedDesc := strings.TrimSpace(req.Description)

	if trimmedDesc == "" {
		return errors.New("поле description не может быть пустым")
	}

	if utf8.RuneCountInString(trimmedDesc) < 10 || utf8.RuneCountInString(trimmedDesc) > 250 {
		return errors.New("описание должно содержать от 10 до 250 символов")
	}

	exists, err := s.genres.Exist(req.GenreID)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("жанра с таким ID не существует")
	}

	return nil
}
