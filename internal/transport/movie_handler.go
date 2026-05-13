package transport

import (
	"errors"
	"movies-project/internal/models"
	"movies-project/internal/repository"
	"movies-project/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MovieHandler struct {
	service services.MovieService
}

func NewMovieHandler(service services.MovieService) *MovieHandler {
	return &MovieHandler{service: service}
}

func (h *MovieHandler) RegisterRoutes(r *gin.Engine) {
	movies := r.Group("/movies")
	{
		movies.GET("", h.ListMovies)
		movies.POST("", h.Create)
		movies.GET("/:id", h.GetByID)
		movies.PUT("/:id", h.Update)
		movies.DELETE("/:id", h.Delete)
	}
}

func (h *MovieHandler) Create(c *gin.Context) {
	var req models.MovieUpsertRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	movie, err := h.service.CreateMovie(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, movie)
}

func (h *MovieHandler) ListMovies(c *gin.Context) {
	var filter repository.MovieFilter

	if genreIDStr := c.Query("genre_id"); genreIDStr != "" {
		genreID, err := strconv.ParseUint(genreIDStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "genre_id должен быть числом"})
			return
		}
		id := uint(genreID)
		filter.GenreID = &id
	}

	if year := c.Query("year"); year != "" {
		yearInt, err := strconv.Atoi(year)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "year должен быть числом"})
			return
		}
		filter.Year = &yearInt
	}

	movies, err := h.service.ListMovies(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, movies)
}

func (h *MovieHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID должен быть числом"})
		return
	}

	movie, err := h.service.GetMovieByID(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, movie)
}

func (h *MovieHandler) Update(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "некорректный идентификатор"})
		return
	}

	var req models.MovieUpsertRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	movie, err := h.service.UpdateMovie(uint(id), req)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, movie)
}

func (h MovieHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID должен быть числом"})
		return
	}

	if err := h.service.DeleteMovie(uint(id)); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}
