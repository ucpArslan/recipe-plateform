package userhttp

import (
	"net/http"
	"strconv"

	"recipe-plateform/internal/user/presentation/adapter"

	"github.com/gin-gonic/gin"
)

func (u UserHandler) GetByID(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid user id",
		})
		return
	}

	user, err := u.GetUser.Execute(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, adapter.ToResponse(user))
}
