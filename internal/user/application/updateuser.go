package application

import (
	"errors"

	"recipe-plateform/internal/user/domain"
)

type UpdateUser struct {
	UserRepo domain.UserRepo
}

func (u UpdateUser) Execute(userID int, update domain.UserUpdate) (*domain.User, error) {

	if userID <= 0 {
		return nil, errors.New("invalid user id")
	}

	if update.Username == nil && update.Email == nil {
		return nil, errors.New("at least one field is required")
	}

	return u.UserRepo.Update(userID, update)
}
