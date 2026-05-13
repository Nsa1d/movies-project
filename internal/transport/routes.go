package transport

import (
	"movies-project/internal/services"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	router *gin.Engine,
	movieService services.MovieService,
	genreService services.GenreService,
) {
	movieHandler := NewMovieHandler(movieService)
	genreHandler := NewGenreHandler(genreService)

	movieHandler.RegisterRoutes(router)
	genreHandler.RegisterRoutes(router)
}
