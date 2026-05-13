package models

import "gorm.io/gorm"

type Genre struct {
	gorm.Model
	Name string `gorm:"unique"`
}

type GenreCreateRequest struct {
	Name string `gorm:"unique"`
}
