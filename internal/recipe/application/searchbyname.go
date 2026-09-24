package application

import "recipe-plateform/internal/recipe/domain"

type SearchRecipe struct {
	RecipeRepo domain.RecipeRepo
}

func (s SearchRecipe) Execute(query string) ([]domain.Recipe, error) {
	return s.RecipeRepo.SearchByname(query)
}
