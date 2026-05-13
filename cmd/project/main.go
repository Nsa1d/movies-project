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

	if err := db.AutoMigrate(&models.Movie{}); err != nil {
		log.Fatalf("не удалось выполнить миграции: %v", err)
	}

	movieRepo := repository.NewMovieRepository(db)
	genreRepo := &repository.MockGenreRepository{}

	movieService := services.NewMovieService(movieRepo, genreRepo)

	router := gin.Default()

	transport.RegisterRoutes(router, movieService)

	if err := router.Run(); err != nil {
		log.Fatalf("не удалось запустить HTTP-сервер: %v", err)
	}
}
