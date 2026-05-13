package models

import (
	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model
	Title       string  `json:"title"`
	Year        int     `json:"year"`
	DurationMin int     `json:"duration_min"`
	Rating      float64 `json:"rating"`
	Country     string  `json:"country"`
	Description string  `json:"description"`
	GenreID     uint    `json:"genre_id"`
}

type MovieUpsertRequest struct {
	Title       string  `json:"title"`
	Year        int     `json:"year"`
	DurationMin int     `json:"duration_min"`
	Rating      float64 `json:"rating"`
	Country     string  `json:"country"`
	Description string  `json:"description"`
	GenreID     uint    `json:"genre_id"`
}

