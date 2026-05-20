package transport

import (
	"movies-project/internal/models"
	"movies-project/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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
	err := c.ShouldBindJSON(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = h.service.PostWatchlist(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "список просмотренного успешно создан"})
}

func (h *WatchlistHandler) Delete(c *gin.Context) {
	index := c.Param("id")
	id, err := strconv.Atoi(index)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = h.service.DeleteWatchlist(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "список просмотренного удален"})
}

func (h *WatchlistHandler) Get(c *gin.Context) {
	index := c.Param("id")
	id, err := strconv.Atoi(index)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	wList, err := h.service.GetWatchlist(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": wList})
}
