package repository

import (
	"errors"
	"movies-project/internal/models"

	"gorm.io/gorm"
)

type MovieFilter struct {
	GenreID *uint
	Year    *int
}

type MovieRepository interface {
	Create(movie *models.Movie) error

	GetByID(id uint) (*models.Movie, error)

	List(filter MovieFilter) ([]models.Movie, error)

	Update(movie *models.Movie) error

	Delete(id uint) error

	//сделал фильтр, чтобы не смогли добавить уже существующий фильм
	GetByTitle(title string) (*models.Movie, error)
}

type gormMovieRepository struct {
	db *gorm.DB
}

func NewMovieRepository(db *gorm.DB) MovieRepository {
	return &gormMovieRepository{db: db}
}

func (r *gormMovieRepository) Create(movie *models.Movie) error {
	if movie == nil {
		return nil
	}

	return r.db.Create(movie).Error
}

func (r *gormMovieRepository) GetByID(id uint) (*models.Movie, error) {
	var movie models.Movie

	if err := r.db.First(&movie, id).Error; err != nil {
		return nil, err
	}

	return &movie, nil
}

func (r *gormMovieRepository) List(filter MovieFilter) ([]models.Movie, error) {
	var movies []models.Movie

	query := r.db.Model(&models.Movie{})

	if filter.GenreID != nil {
		query = query.Where("genre_id = ?", *filter.GenreID)
	}

	if filter.Year != nil {
		query = query.Where("year = ?", *filter.Year)
	}

	if err := query.Find(&movies).Error; err != nil {
		return nil, err
	}

	return movies, nil
}

func (r *gormMovieRepository) Update(movie *models.Movie) error {
	if movie == nil {
		return nil
	}

	return r.db.Save(movie).Error
}

func (r *gormMovieRepository) Delete(id uint) error {
	return r.db.Delete(&models.Movie{}, id).Error
}

func (r *gormMovieRepository) GetByTitle(title string) (*models.Movie, error) {
	var movie models.Movie
	err := r.db.Where("title = ?", title).First(&movie).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &movie, err
}
