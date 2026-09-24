package httprecipe

import (
	"net/http"
	"strconv"

	"recipe-plateform/internal/recipe/application"
	adapter "recipe-plateform/internal/recipe/presentation/adaptor"

	"github.com/gin-gonic/gin"
)

type FavoriteHandler struct {
	AddFavorite    application.AddFavorite
	RemoveFavorite application.RemoveFavorite
	GetFavorites   application.GetFavorites
}

func (h FavoriteHandler) Add(c *gin.Context) {

	userIDValue, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "user not authenticated",
		})
		return
	}

	userID := userIDValue.(int)

	recipeID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid recipe id",
		})
		return
	}

	err = h.AddFavorite.Execute(uint(userID), recipeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "recipe added to favorites",
	})
}

func (h FavoriteHandler) Remove(c *gin.Context) {

	userIDValue, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "user not authenticated",
		})
		return
	}

	userID := userIDValue.(int)

	recipeID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid recipe id",
		})
		return
	}

	err = h.RemoveFavorite.Execute(uint(userID), recipeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "recipe removed from favorites",
	})
}

func (h FavoriteHandler) Get(c *gin.Context) {

	userIDValue, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "user not authenticated",
		})
		return
	}

	userID := userIDValue.(int)

	recipes, err := h.GetFavorites.Execute(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, adapter.ToResponseList(recipes))
}
