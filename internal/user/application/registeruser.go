package application

import (
	"errors"
	"net/mail"

	"recipe-plateform/internal/user/domain"

	"golang.org/x/crypto/bcrypt"
)

type RegisterUser struct {
	UserRepo domain.UserRepo
}

func (r RegisterUser) Execute(user domain.User) (*domain.User, error) {

	_, err := mail.ParseAddress(user.Email)
	if err != nil {
		return nil, errors.New("invalid email format")
	}

	if user.Username == "" {
		return nil, errors.New("username is required")
	}

	if user.Email == "" {
		return nil, errors.New("email is required")
	}

	if user.Password == "" {
		return nil, errors.New("password is required")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(user.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return nil, errors.New("password hashing failed")
	}

	user.Password = string(hashedPassword)

	return r.UserRepo.Create(&user)
}
