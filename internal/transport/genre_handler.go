package transport

import (
	"movies-project/internal/models"
	"movies-project/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GenreHandler struct {
	genre services.GenreService
}

func NewGenreHandler(genre services.GenreService) *GenreHandler {
	return &GenreHandler{genre: genre}
}
func (h *GenreHandler) RegisterRoutes(r *gin.Engine) {
	genre := r.Group("/genre")
	{
		genre.GET("", h.GetGenres)
		genre.POST("", h.PostGenre)
	}
}

func (g *GenreHandler) PostGenre(c *gin.Context) {
	var req models.GenreCreateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	genre, err := g.genre.CreateGenre(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, genre)
}

func (h *GenreHandler) GetGenres(c *gin.Context) {
	var genres []models.Genre

	err := h.genre.GetAllGenres(&genres)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": genres})
}
