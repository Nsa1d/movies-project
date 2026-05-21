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

type ReviewHandler struct {
	service services.ReviewService
}

func NewReviewHandler(service services.ReviewService) *ReviewHandler {
	return &ReviewHandler{service: service}
}

func (h *ReviewHandler) RegisterRoutes(r *gin.Engine) {
	review := r.Group("movies")
	{
		review.POST("/:id/reviews", h.Create)
		review.GET("/:id/reviews", h.GetAll)
	}
}

func (h *ReviewHandler) Create(c *gin.Context) {
	movieID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "некорректный movie_id"})
		return
	}

	var req models.ReviewCreateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.MovieID != uint(movieID) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "movie_id в URL и в теле запроса не совпадают"})
		return
	}

	review, err := h.service.CreateReview(req)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, review)
}

func (h ReviewHandler) GetAll(c *gin.Context) {
	reviews, err := h.service.GetAllReview()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, reviews)
}