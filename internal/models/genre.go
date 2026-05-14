package models

type Genre struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `json:"name" binding:"required"`
}
type GenreCreateRequest struct {
	Name string `json:"name"`
}
