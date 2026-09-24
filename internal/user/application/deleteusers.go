package application

import (
	"errors"
	"recipe-plateform/internal/user/domain"
)

type DeleteUser struct {
	UserRepo domain.UserRepo
}

func (d DeleteUser) Execute(userID int) error {
	if userID <= 0 {
		return errors.New("The in-valid id")
	}
	return d.UserRepo.Delete(userID)
}
