package models

type RecipeResponse struct {
	Id             int    `json:"id"`
	Title          string `json:"title"`
	Image          string `json:"image"`
	ReadyInMinutes int    `json:"readyInMinutes"`
	Servings       int    `json:"servings"`
}
