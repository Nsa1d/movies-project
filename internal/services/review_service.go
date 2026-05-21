package services

import (
	"errors"
	"movies-project/internal/models"
	"movies-project/internal/repository"
)

type ReviewService interface {
	CreateReview(req models.ReviewCreateRequest) (*models.Review, error)

	GetAllReview() ([]models.Review, error)
}

type reviewService struct {
	reviews repository.ReviewRepository
	movie   repository.MovieRepository
	user    repository.UserRepository
}

func NewReviewService(
	reviews repository.ReviewRepository,
	movies repository.MovieRepository,
	users repository.UserRepository,
) ReviewService {
	return &reviewService{
		reviews: reviews,
		movie:   movies,
		user:    users,
	}
}

func (s *reviewService) CreateReview(req models.ReviewCreateRequest) (*models.Review, error) {
	if err := s.validateReview(req); err != nil {
		return nil, err
	}

	review := &models.Review{
		MovieID: req.MovieID,
		UserID:  req.UserID,
		Rating:  req.Rating,
		Review:  req.Review,
	}

	if err := s.reviews.Create(review); err != nil {
		return nil, err
	}

	return review, nil
}

func (s *reviewService) GetAllReview() ([]models.Review, error) {
	return s.reviews.Get()
}

func (s *reviewService) validateReview(req models.ReviewCreateRequest) error {
	found, err := s.movie.GetByID(req.MovieID)
	if err != nil {
		return err
	}
	if found == nil {
		return errors.New("нет фильма с таким ID")
	}

	user, err := s.user.GetByID(req.UserID)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("пользователь с таким ID не существует")
	}

	if req.Rating < 1 || req.Rating > 10 {
		return errors.New("рейтинг должен быть от 1 до 10")
	}

	if err := s.AddRating(req.MovieID, float64(req.Rating)); err != nil {
		return err
	}

	return nil
}

func (s *reviewService) AddRating(movieID uint, newRating float64) error {
	movie, err := s.movie.GetByID(movieID)
	if err != nil {
		return err
	}

	total := movie.Rating * float64(movie.RatingCount)
	movie.RatingCount++
	movie.Rating = (total + float64(newRating)) / float64(movie.RatingCount)

	return s.movie.Update(movie)
}
