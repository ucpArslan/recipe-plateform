package application

import (
	"errors"

	"recipe-plateform/internal/user/domain"

	"golang.org/x/crypto/bcrypt"
)

type LoginUser struct {
	UserRepo     domain.UserRepo
	TokenService TokenService
}

func (l LoginUser) Execute(email string, password string) (string, error) {

	users, err := l.UserRepo.GetAllUser()
	if err != nil {
		return "", err
	}

	for _, user := range users {

		if user.Email != email {
			continue
		}

		err := bcrypt.CompareHashAndPassword(
			[]byte(user.Password),
			[]byte(password),
		)

		if err != nil {
			return "", errors.New("invalid email or password")
		}

		token, err := l.TokenService.GenerateToken(&user)
		if err != nil {
			return "", err
		}

		return token, nil
	}

	return "", errors.New("invalid email or password")
}
