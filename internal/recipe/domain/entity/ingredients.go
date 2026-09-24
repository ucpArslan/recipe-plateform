package entity

type Ingredient struct {
	ID   int    `gorm:"primaryKey" json:"id"`
	Name string `gorm:"not null;unique" json:"name"`
}
