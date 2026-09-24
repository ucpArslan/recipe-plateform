package httprecipe

import (
	"net/http"
	"strconv"

	"recipe-plateform/internal/recipe/application"

	"github.com/gin-gonic/gin"
)

type GetRecipeDetailsHandler struct {
	GetRecipeDetails application.GetRecipeDetails
}

func (r GetRecipeDetailsHandler) GetDetails(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid recipe id",
		})
		return
	}

	recipe, err := r.GetRecipeDetails.Execute(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, recipe)
}
