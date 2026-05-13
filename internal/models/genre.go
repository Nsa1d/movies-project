package models

import "gorm.io/gorm"

type Genre struct {
	gorm.Model
	Name string
}

type GenreCreateRequest struct {
	Name string `json:"name"`
}
