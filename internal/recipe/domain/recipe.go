package domain

type Recipe struct {
	Id             int    `gorm:"primaryKey"`
	Title          string `gorm:"not null"`
	Image          string
	ReadyInMinutes int
	Servings       int
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
