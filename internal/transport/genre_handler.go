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
func (g *GenreHandler) RegisterRoutes(r *gin.Engine) {
	genre := r.Group("/genre")
	{
		genre.GET("/genre")
		genre.POST("/genre")
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

	c.JSON(http.StatusCreated, genre)
}
