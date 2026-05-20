package repository

import (
	"errors"
	"movies-project/internal/models"

	"gorm.io/gorm"
)

type WatchlistRepository interface {
	CreateWatchlist(watchlist *models.Watchlist) error
	DeleteWatchlist(id uint) error
	FindWatchlist(id uint) (*models.Watchlist, error)
}
type gormWatchlistRepository struct {
	db *gorm.DB
}

func NewWatchlistRepository(db *gorm.DB) WatchlistRepository {
	return &gormGenreRepository{db: db}
}

func (r *gormGenreRepository) CreateWatchlist(watchlist *models.Watchlist) error {
	if watchlist == nil {
		return nil
	}
	return r.db.Create(watchlist).Error
}

func (r *gormGenreRepository) DeleteWatchlist(id uint) error {
	return r.db.Delete(&models.Watchlist{}, "id = ?", id).Error
}

func (r *gormGenreRepository) FindWatchlist(id uint) (*models.Watchlist, error) {
	var wList models.Watchlist
	err := r.db.Where("id = ?", id).First(&wList).Error
	if err != nil {
		return nil, errors.New("список просмотренного с таким id не найден")
	}
	return &wList, nil
}
