package models

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	MovieID uint   `json:"movie_id"`
	UserID  uint   `json:"user_id"`
	Rating  int    `json:"rating"`
	Review  string `json:"review"`
}

type ReviewCreateRequest struct {
	MovieID uint   `json:"movie_id"`
	UserID  uint   `json:"user_id"`
	Rating  int    `json:"rating"`
	Review  string `json:"review"`
}
