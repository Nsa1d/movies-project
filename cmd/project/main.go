package main

import (
	"log"
	"movies-project/internal/config"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.SetUpDatabaseConnection()

	if err := db.AutoMigrate(); err != nil {
		log.Fatalf("не удалось выполнить миграции: %v", err)
	}

	//создание репозитория

	router := gin.Default()

	//тут должна быть регистрация роутов

	if err := router.Run(); err != nil {
		log.Fatalf("не удалось запустить HTTP-сервер: %v", err)
	}
}
