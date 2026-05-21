package models

import "gorm.io/gorm"

type Watchlist struct {
	gorm.Model
	UserID  uint `json:"user_id"`
	MovieID uint `json:"movie_id"`
}

type WatchlistRequest struct {
	UserID  uint `json:"user_id"`
	MovieID uint `json:"movie_id"`
}
