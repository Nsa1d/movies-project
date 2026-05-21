package services

import (
	"errors"
	"movies-project/internal/models"
	"movies-project/internal/repository"
)

type WatchlistService interface {
	PostWatchlist(req models.WatchlistRequest) (*models.Watchlist, error)
	DeleteWatchlist(id uint) error
	GetWatchlist(id uint) (*models.Watchlist, error)
}

type watchlistService struct {
	movie     repository.MovieRepository
	user      repository.UserRepository
	watchlist repository.WatchlistRepository
}

func NewWatchlistService(watchlist repository.WatchlistRepository, movie repository.MovieRepository, user repository.UserRepository) WatchlistService {
	return &watchlistService{
		watchlist: watchlist,
		user:      user,
		movie:     movie,
	}
}

func (s *watchlistService) PostWatchlist(req models.WatchlistRequest) (*models.Watchlist, error) {
	if errValid := s.validatePost(req); errValid != nil {
		return nil, errValid
	}
	wList := &models.Watchlist{
		UserID:  req.UserID,
		MovieID: req.MovieID,
	}
	if err := s.watchlist.CreateWatchlist(wList); err != nil {
		return nil, err
	}

	return wList, nil
}

func (s *watchlistService) DeleteWatchlist(id uint) error {
	if _, err := s.watchlist.FindWatchlist(id); err != nil {
		return errors.New("неверный id")
	}
	return s.watchlist.DeleteWatchlist(id)
}

func (s *watchlistService) GetWatchlist(id uint) (*models.Watchlist, error) {
	if _, err := s.user.GetByID(id); err != nil {
		return nil, errors.New("неверный id")
	}
	return s.watchlist.FindWatchlist(id)
}

func (s *watchlistService) validatePost(req models.WatchlistRequest) error {
	if _, err := s.movie.GetByID(req.MovieID); err != nil {
		return errors.New("фильма с таким id не существует")
	}
	if _, err := s.user.GetByID(req.UserID); err != nil {
		return errors.New("пользователя с таким id не существует")
	}
	return nil
}
