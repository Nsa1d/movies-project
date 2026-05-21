package transport

import (
	"movies-project/internal/services"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	router *gin.Engine,
	movieService services.MovieService,
	genreService services.GenreService,
	reviewService services.ReviewService,
) {
	movieHandler := NewMovieHandler(movieService)
	genreHandler := NewGenreHandler(genreService)
	reviewHandler := NewReviewHandler(reviewService)

	movieHandler.RegisterRoutes(router)
	genreHandler.RegisterRoutes(router)
	reviewHandler.RegisterRoutes(router)
}
