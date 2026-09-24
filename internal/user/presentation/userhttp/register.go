package userhttp

import (
	"net/http"

	"recipe-plateform/internal/user/application"
	"recipe-plateform/internal/user/domain"

	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type RegisterHandler struct {
	RegisterUser application.RegisterUser
}

func (h RegisterHandler) Register(c *gin.Context) {

	var request RegisterRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user := domain.User{
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
	}

	createdUser, err := h.RegisterUser.Execute(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":       createdUser.ID,
		"username": createdUser.Username,
		"email":    createdUser.Email,
	})
}
