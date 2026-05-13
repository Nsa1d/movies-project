package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username  string
	Password  string
	Login     string
	Watchlist string
	Rewiews   string
}

type Registration struct {
	Username string
	Password string
	Login    string
}
