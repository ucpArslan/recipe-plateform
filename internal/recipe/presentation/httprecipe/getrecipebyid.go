package httprecipe

import (
	"net/http"
	"strconv"

	"recipe-plateform/internal/recipe/application"
	adapter "recipe-plateform/internal/recipe/presentation/adaptor"

	"github.com/gin-gonic/gin"
)

type GetRecipeByIDHandler struct {
	GetRecipeByID application.GetRecipeByID
}

func (r GetRecipeByIDHandler) GetByID(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid recipe id",
		})
		return
	}

	recipe, err := r.GetRecipeByID.Execute(id)
	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, adapter.ToResponse(*recipe))
}
