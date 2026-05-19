package transport

import (
	"movies-project/internal/models"
	"movies-project/internal/services"
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service services.UserService
}

func NewUserHandler(service services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) RegisterRoutes(r *gin.Engine) {
	user := r.Group("/user")
	{
		user.GET("/:id", h.GetUser)
		user.POST("/register", h.Register)
		user.POST("/login", h.Login)
	}
}

func (h *UserHandler) Register(c *gin.Context) {
	var reg models.Registration

	if err := c.ShouldBindJSON(&reg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := h.service.CreateUser(reg)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "вы успешно зарегались"})
}

func (h *UserHandler) GetUser(c *gin.Context) {
	idNet := c.Param("id")
	id, err := strconv.Atoi(idNet)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	user, err := h.service.GetUserByID(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (h *UserHandler) Login(c *gin.Context) {
	var authy models.Login
	if err := c.ShouldBindJSON(&authy); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := h.service.Authenticate(authy)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "вы успешно вошли"})
}
