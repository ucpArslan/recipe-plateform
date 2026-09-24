package adapter

import (
	"recipe-plateform/internal/recipe/domain"
	"recipe-plateform/internal/recipe/presentation/models"
)

func ToResponse(recipe domain.Recipe) models.RecipeResponse {

	return models.RecipeResponse{
		Id:             recipe.Id,
		Title:          recipe.Title,
		Image:          recipe.Image,
		ReadyInMinutes: recipe.ReadyInMinutes,
		Servings:       recipe.Servings,
	}
}

func ToResponseList(recipes []domain.Recipe) []models.RecipeResponse {

	response := make([]models.RecipeResponse, 0, len(recipes))

	for _, recipe := range recipes {

		response = append(response, ToResponse(recipe))

	}

	return response
}
