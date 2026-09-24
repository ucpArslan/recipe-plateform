package domain

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type UserUpdate struct {
	Username *string
	Email    *string
}

type UserRepo interface {
	Create(user *User) (*User, error)
	GetUserByID(userID int) (*User, error)
	Delete(userID int) error
	Update(userID int, user UserUpdate) (*User, error)
	GetAllUser() ([]User, error)
}
