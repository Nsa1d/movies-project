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
	userService services.UserService,
	watchlistService services.WatchlistService,
) {
	movieHandler := NewMovieHandler(movieService)
	genreHandler := NewGenreHandler(genreService)
	reviewHandler := NewReviewHandler(reviewService)
	userHandler := NewUserHandler(userService)
	watchlistHandler := NewWatchlistHandler(watchlistService)

	movieHandler.RegisterRoutes(router)
	genreHandler.RegisterRoutes(router)
	reviewHandler.RegisterRoutes(router)
	userHandler.RegisterRoutes(router)
	watchlistHandler.RegisterRoutes(router)

}
