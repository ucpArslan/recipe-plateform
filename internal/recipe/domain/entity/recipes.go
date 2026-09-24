package entity

type Recipe struct {
	Id             int    `gorm:"primaryKey" json:"id"`
	Title          string `gorm:"not null" json:"title"`
	Image          string `json:"image"`
	ReadyInMinutes int    `json:"readyInMinutes"`
	Servings       int    `json:"servings"`

	Ingredients []Ingredient `gorm:"many2many:recipe_ingredients;" json:"ingredients"`
}
type RecipeDetails struct {
	Id             int      `json:"id"`
	Title          string   `json:"title"`
	Image          string   `json:"image"`
	ReadyInMinutes int      `json:"readyInMinutes"`
	Servings       int      `json:"servings"`
	Summary        string   `json:"summary"`
	Instructions   string   `json:"instructions"`
	Ingredients    []string `json:"ingredients"`
}
