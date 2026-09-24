package domain

type RecipeRepo interface {
	GetRandomRecipe(number int) ([]Recipe, error)
	GetRecipebyID(id int) (*Recipe, error)
	GetRecipeDetails(id int) (*RecipeDetails, error)
	SearchByname(query string) ([]Recipe, error)
	SreachByingre(query string) ([]Recipe, error)
}
type FavoriteRepo interface {
	AddFavorite(userID uint, recipeID int) error
	RemoveFavorite(userID uint, recipeID int) error
	GetFavorites(userID uint) ([]Recipe, error)
}
