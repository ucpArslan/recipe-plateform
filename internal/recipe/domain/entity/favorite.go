package entity

type Favorite struct {
	ID       uint `gorm:"primaryKey"`
	UserID   uint
	RecipeID int
}
