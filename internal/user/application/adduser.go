package application

import (
	"errors"
	"net/mail"
	"recipe-plateform/internal/user/domain"

	"golang.org/x/crypto/bcrypt"
)

type AddUser struct {
	UserRepo domain.UserRepo
}

func (c AddUser) Execute(user domain.User) (*domain.User, error) {

	_, err := mail.ParseAddress(user.Email)
	if err != nil {
		return nil, errors.New("invalid email format")
	}

	// Business Logic
	if user.Username == "" {
		return nil, errors.New("username is required")
	}
	if user.Email == "" {
		return nil, errors.New("Email is required")
	}
	if user.Password == "" {
		return nil, errors.New("The password is required")
	}

	hassedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(user.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return nil, errors.New("The password is invalid")
	}
	user.Password = string(hassedPassword)

	return c.UserRepo.Create(&user)
}
