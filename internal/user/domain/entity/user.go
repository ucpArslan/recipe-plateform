package entity

type User struct {
	ID       int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	Username string `gorm:"type:varchar(255);not null;unique" json:"username"`
	Email    string `gorm:"type:varchar(255);unique" json:"email"`
	Password string `gorm:"type:varchar(255);" json:"password"`
}
type UserUpdate struct {
	Username *string
	Email    *string
}
