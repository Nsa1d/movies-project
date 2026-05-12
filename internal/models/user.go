package models

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	Name string `json:"name"`
}