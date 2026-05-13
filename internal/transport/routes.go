package transport

import (
	"movies-project/internal/services"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	router *gin.Engine,
	movieService services.MovieService,
) {
	movieHandler := NewMovieHandler(movieService)

	movieHandler.RegisterRoutes(router)
}