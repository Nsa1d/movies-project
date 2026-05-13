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

	if err := db.AutoMigrate(&models.Genre{}); err != nil {
		log.Fatalf("не удалось выполнить миграции: %v", err)
	}

	//создание репозитория
	genre := repository.NewGenreRepository(db)
	genreService := services.NewGenreService(genre)

	router := gin.Default()
	transport.RegisterRoutes(router, genreService)
	//тут должна быть регистрация роутов

	if err := router.Run(); err != nil {
		log.Fatalf("не удалось запустить HTTP-сервер: %v", err)
	}
}
