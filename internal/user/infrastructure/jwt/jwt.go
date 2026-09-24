package jwt

import (
	"errors"
	"time"

	"recipe-plateform/internal/user/domain"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService struct {
	SecretKey string
}

type Claims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

func (j JWTService) GenerateToken(user *domain.User) (string, error) {

	claims := Claims{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	return token.SignedString([]byte(j.SecretKey))
}

func (j JWTService) ValidateToken(tokenString string) (*domain.Claims, error) {

	token, err := jwt.ParseWithClaims(
		tokenString,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}

			return []byte(j.SecretKey), nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return &domain.Claims{
		UserID: claims.UserID,
	}, nil
}
