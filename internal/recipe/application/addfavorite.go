package application

import "recipe-plateform/internal/recipe/domain"

type AddFavorite struct {
	FavoriteRepo domain.FavoriteRepo
	RecipeRepo   domain.RecipeRepo
}

func (a AddFavorite) Execute(userID uint, recipeID int) error {

	_, err := a.RecipeRepo.GetRecipebyID(recipeID)
	if err != nil {
		return err
	}

	return a.FavoriteRepo.AddFavorite(userID, recipeID)
}
