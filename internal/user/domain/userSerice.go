package domain

type UserService interface {
	Creat(user User) (*User, error)
	getUserByID(userID int) (*User, error)
	Delete(userID int) error
	Update(user UserUpdate) (*User, error)
	GetAllUser() ([]User, error)
}
