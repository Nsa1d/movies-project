package transport

import (
	"movies-project/internal/services"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	router *gin.Engine,
	userService services.UserService,
	movieService services.MovieService,
	genreService services.GenreService,
) {
	userHandler := NewUserHandler(userService)
	movieHandler := NewMovieHandler(movieService)
	genreHandler := NewGenreHandler(genreService)

	userHandler.RegisterRoutes(router)
	movieHandler.RegisterRoutes(router)
	genreHandler.RegisterRoutes(router)
}
