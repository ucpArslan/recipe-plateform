package httprecipe

import (
	"net/http"
	"strconv"

	"recipe-plateform/internal/recipe/application"

	adapter "recipe-plateform/internal/recipe/presentation/adaptor"

	"github.com/gin-gonic/gin"
)

type RecipeHandler struct {
	GetRandomRecipe application.GetRandomRecipe
}

func (r RecipeHandler) GetRandom(c *gin.Context) {

	number := 1

	if value := c.Query("number"); value != "" {

		n, err := strconv.Atoi(value)
		if err != nil {

			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid number",
			})
			return
		}

		number = n
	}

	recipes, err := r.GetRandomRecipe.Execute(number)
	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, adapter.ToResponseList(recipes))
}
