package application

import "recipe-plateform/internal/recipe/domain"

type GetRecipeByID struct {
	RecipeRepo domain.RecipeRepo
}

func (g GetRecipeByID) Execute(id int) (*domain.Recipe, error) {

	return g.RecipeRepo.GetRecipebyID(id)

}
