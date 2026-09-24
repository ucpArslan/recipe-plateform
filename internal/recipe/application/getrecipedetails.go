package application

import "recipe-plateform/internal/recipe/domain"

type GetRecipeDetails struct {
	RecipeRepo domain.RecipeRepo
}

func (g GetRecipeDetails) Execute(id int) (*domain.RecipeDetails, error) {
	return g.RecipeRepo.GetRecipeDetails(id)
}
