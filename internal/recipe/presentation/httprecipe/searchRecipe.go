package httprecipe

import (
	"net/http"

	"recipe-plateform/internal/recipe/application"
	adapter "recipe-plateform/internal/recipe/presentation/adaptor"

	"github.com/gin-gonic/gin"
)

type SearchRecipeHandler struct {
	SearchRecipe       application.SearchRecipe
	SearchByIngredient application.SearchByIngredient
}

func (r SearchRecipeHandler) Search(c *gin.Context) {
	query := c.Query("query")

	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "query is required",
		})
		return
	}

	recipes, err := r.SearchRecipe.Execute(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, adapter.ToResponseList(recipes))
}

func (r SearchRecipeHandler) SearchIngredient(c *gin.Context) {
	query := c.Query("query")

	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "query is required",
		})
		return
	}

	recipes, err := r.SearchByIngredient.Execute(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, adapter.ToResponseList(recipes))
}
