package application

import "recipe-plateform/internal/recipe/domain"

type GetRandomRecipe struct {
	RecipeRepo domain.RecipeRepo
}

func (g GetRandomRecipe) Execute(number int) ([]domain.Recipe, error) {

	return g.RecipeRepo.GetRandomRecipe(number)

}
