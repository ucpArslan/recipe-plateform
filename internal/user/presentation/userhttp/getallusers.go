package userhttp

import (
	"net/http"
	"recipe-plateform/internal/user/presentation/adapter"

	"github.com/gin-gonic/gin"
)

func (u UserHandler) GetAll(c *gin.Context) {

	users, err := u.GetAllUser.Execute()

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, adapter.ToResponses(users))
}
