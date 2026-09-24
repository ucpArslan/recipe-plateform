package application

import "recipe-plateform/internal/recipe/domain"

type GetFavorites struct {
	FavoriteRepo domain.FavoriteRepo
}

func (g GetFavorites) Execute(userID uint) ([]domain.Recipe, error) {
	return g.FavoriteRepo.GetFavorites(userID)
}
