package application

import (
	"errors"
	"recipe-plateform/internal/user/domain"
)

type GetUser struct {
	UserRepo domain.UserRepo
}

func (g GetUser) Execute(userID int) (*domain.User, error) {

	// Business Logic
	if userID <= 0 {
		return nil, errors.New("invalid user id")
	}

	return g.UserRepo.GetUserByID(userID)
}
