package application

import "recipe-plateform/internal/recipe/domain"

type RemoveFavorite struct {
	FavoriteRepo domain.FavoriteRepo
}

func (r RemoveFavorite) Execute(userID uint, recipeID int) error {
	return r.FavoriteRepo.RemoveFavorite(userID, recipeID)
}
