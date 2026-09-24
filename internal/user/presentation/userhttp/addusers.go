package userhttp

import (
	"net/http"

	"recipe-plateform/internal/user/application"
	"recipe-plateform/internal/user/presentation/adapter"
	"recipe-plateform/internal/user/presentation/models"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	CreateUser application.AddUser
	GetUser    application.GetUser
	DeleteUser application.DeleteUser
	UpdateUser application.UpdateUser
	GetAllUser application.GetAllUser
	LoginUser  application.LoginUser
}

func (u UserHandler) Create(c *gin.Context) {

	var request models.CreateUserRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user := adapter.ToDomain(request)

	result, err := u.CreateUser.Execute(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, adapter.ToResponse(result))
}
