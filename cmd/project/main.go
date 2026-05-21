package main

import (
	"log"
	"movies-project/internal/config"
	"movies-project/internal/models"
	"movies-project/internal/repository"
	"movies-project/internal/services"
	"movies-project/internal/transport"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.SetUpDatabaseConnection()

	if err := db.AutoMigrate(&models.Genre{}, &models.Movie{}, &models.Review{}, &models.User{}, &models.Watchlist{}); err != nil {
		log.Fatalf("не удалось выполнить миграции: %v", err)
	}

	movieRepo := repository.NewMovieRepository(db)
	genreRepo := repository.NewGenreRepository(db)
	reviewRepo := repository.NewReviewRepository(db)
	userRepo := repository.NewUserRepository(db)
	watchlistRepo := repository.NewWatchlistRepository(db)

	movieService := services.NewMovieService(movieRepo, genreRepo)
	genreService := services.NewGenreService(genreRepo)
	reviewService := services.NewReviewService(reviewRepo, movieRepo, userRepo)
	userService := services.NewUserService(userRepo)
	watchlistService := services.NewWatchlistService(watchlistRepo, movieRepo, userRepo)

	router := gin.Default()
	transport.RegisterRoutes(router, movieService, genreService, reviewService, userService, watchlistService)

	if err := router.Run(); err != nil {
		log.Fatalf("не удалось запустить HTTP-сервер: %v", err)
	}
}
