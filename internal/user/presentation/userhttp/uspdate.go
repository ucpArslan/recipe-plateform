package userhttp

import (
	"net/http"
	"strconv"

	"recipe-plateform/internal/user/presentation/adapter"
	"recipe-plateform/internal/user/presentation/models"

	"github.com/gin-gonic/gin"
)

func (u UserHandler) Update(c *gin.Context) {

	id := c.Param("id")

	userID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid user id",
		})
		return
	}

	var request models.UpdateUserRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	update := adapter.ToUserUpdate(request)

	result, err := u.UpdateUser.Execute(userID, update)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, adapter.ToResponse(result))
}
