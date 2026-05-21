package transport

import (
	"errors"
	"movies-project/internal/models"
	"movies-project/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type WatchlistHandler struct {
	service services.WatchlistService
}

func NewWatchlistHandler(service services.WatchlistService) *WatchlistHandler {
	return &WatchlistHandler{
		service: service,
	}
}

func (h *WatchlistHandler) RegisterRoutes(r *gin.Engine) {
	wList := r.Group("/watchlist")
	{
		wList.POST("", h.Post)
		wList.DELETE("/:id", h.Delete)
	}

	r.GET("/users/:id/watchlist", h.Get)
}

func (h *WatchlistHandler) Post(c *gin.Context) {
	var req models.WatchlistRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := h.service.PostWatchlist(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusCreated, gin.H{"message": "список просмотренного успешно создан"})
}

func (h *WatchlistHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id должен быть числом"})
		return
	}
	err = h.service.DeleteWatchlist(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "запись с таким id не найдена"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (h *WatchlistHandler) Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id должен быть числом"})
		return
	}

	wList, err := h.service.GetWatchlist(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "запись с таким id не найдена"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, wList)
}
