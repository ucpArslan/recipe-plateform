package application

import "recipe-plateform/internal/recipe/domain"

type SearchByIngredient struct {
	RecipeRepo domain.RecipeRepo
}

func (s SearchByIngredient) Execute(query string) ([]domain.Recipe, error) {
	return s.RecipeRepo.SreachByingre(query)
}
