package models

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	Name   string `json:"name"`
	Author string `json:"author"`
}
