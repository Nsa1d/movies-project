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
) {
	movieHandler := NewMovieHandler(movieService)
	genreHandler := NewGenreHandler(genreService)
	userHandler := NewUserHandler(userService)

	movieHandler.RegisterRoutes(router)
	genreHandler.RegisterRoutes(router)
	userHandler.RegisterRoutes(router)
}
