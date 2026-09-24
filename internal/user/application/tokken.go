package application

import "recipe-plateform/internal/user/domain"

type TokenService interface {
	GenerateToken(user *domain.User) (string, error)
	ValidateToken(token string) (*domain.Claims, error)
}
