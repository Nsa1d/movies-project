package transport

import (
	"movies-project/internal/services"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	router *gin.Engine,
	movieService services.MovieService,
	genreService services.GenreService,
	userService services.UserService,
	watchlistService services.WatchlistService,
) {
	movieHandler := NewMovieHandler(movieService)
	genreHandler := NewGenreHandler(genreService)
	userHandler := NewUserHandler(userService)
	watchlistHandler := NewWatchlistHandler(watchlistService)

	movieHandler.RegisterRoutes(router)
	genreHandler.RegisterRoutes(router)
	userHandler.RegisterRoutes(router)
	watchlistHandler.RegisterRoutes(router)
}
