package repository

import (
	"errors"
	"movies-project/internal/models"

	"gorm.io/gorm"
)

type ReviewRepository interface {
	Create(review *models.Review) error

	Get() ([]models.Review, error)

}

type gormReviewRepository struct {
	db *gorm.DB
}

func NewReviewRepository(db *gorm.DB) ReviewRepository {
	return &gormReviewRepository{db: db}
}

func (r *gormReviewRepository) Create(review *models.Review) error {
	if review == nil {
		return nil
	}

	return r.db.Create(review).Error
}

func (r *gormReviewRepository) Get() ([]models.Review, error) {
	var reviews []models.Review

	err := r.db.Find(&reviews).Error
	if err != nil {
			return nil, errors.New("нет отзывов к этому фильму")
	}

	return reviews, nil
}