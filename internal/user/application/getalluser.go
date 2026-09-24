package application

import "recipe-plateform/internal/user/domain"

type GetAllUser struct {
	UserRepo domain.UserRepo
}

func (g GetAllUser) Execute() ([]domain.User, error) {

	users, err := g.UserRepo.GetAllUser()
	if err != nil {
		return nil, err
	}

	return users, nil
}
