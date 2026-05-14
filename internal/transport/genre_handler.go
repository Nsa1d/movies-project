package transport

import (
	"movies-project/internal/models"
	"movies-project/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GenreHandler struct {
	service services.GenreService
}

func NewGenreHandler(service services.GenreService) *GenreHandler {
	return &GenreHandler{service: service}
}
func (h *GenreHandler) RegisterRoutes(r *gin.Engine) {
	genre := r.Group("/genre")
	{
		genre.GET("", h.GetGenres)
		genre.POST("", h.PostGenre)
	}
}

func (h *GenreHandler) PostGenre(c *gin.Context) {
	var req models.GenreCreateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	genre, err := h.service.CreateGenre(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, genre)
}

func (h *GenreHandler) GetGenres(c *gin.Context) {
	genres, err := h.service.GetAllGenres()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, genres)
}
