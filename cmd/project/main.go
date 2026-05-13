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

	if err := db.AutoMigrate(&models.Genre{}, &models.Movie{}); err != nil {
		log.Fatalf("не удалось выполнить миграции: %v", err)
	}

	//создание репозитория
	movieRepo := repository.NewMovieRepository(db)
	genreRepo := repository.NewGenreRepository(db)

	movieService := services.NewMovieService(movieRepo, genreRepo)
	genreService := services.NewGenreService(genreRepo)

	router := gin.Default()
	transport.RegisterRoutes(router, movieService, genreService)
	//тут должна быть регистрация роутов

	if err := router.Run(); err != nil {
		log.Fatalf("не удалось запустить HTTP-сервер: %v", err)
	}
}
