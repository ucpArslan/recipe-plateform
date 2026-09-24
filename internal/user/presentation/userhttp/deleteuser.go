package userhttp

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h UserHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "some error occur",
		})
	}

	err = h.DeleteUser.Execute(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "user delete successfully",
		})
	}

}
