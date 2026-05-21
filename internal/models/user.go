package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique" json:"username" binding:"required"`
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Registration struct {
	Username string `json:"username" binding:"required"`
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Login struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}
