package adapter

import (
	"recipe-plateform/internal/user/domain"
	"recipe-plateform/internal/user/presentation/models"
)

func ToDomain(request models.CreateUserRequest) domain.User {
	return domain.User{
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
	}
}

func ToResponse(user *domain.User) models.UserResponse {
	return models.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}
}
func ToUserUpdate(request models.UpdateUserRequest) domain.UserUpdate {
	return domain.UserUpdate{
		Username: request.Username,
		Email:    request.Email,
	}
}

func ToResponses(users []domain.User) []models.UserResponse {

	response := make([]models.UserResponse, 0)

	for _, user := range users {

		response = append(response, models.UserResponse{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		})
	}

	return response
}
