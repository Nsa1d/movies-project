package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username  string
	Password  string
	Login     string `gorm:"unique"`
	Watchlist string `gorm:"unique"`
	Rewiews   string
}

type Registration struct {
	Username string
	Password string
	Login    string `gorm:"unique"`
}
