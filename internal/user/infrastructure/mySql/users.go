package mySql

import (
	"errors"

	"recipe-plateform/internal/user/domain"
	"recipe-plateform/internal/user/domain/entity"

	"gorm.io/gorm"
)

type UserImpl struct {
	DB *gorm.DB
}

// create the user and assign it a ID
func (u UserImpl) Create(user *domain.User) (*domain.User, error) {

	userEntity := entity.User{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}

	if err := u.DB.Create(&userEntity).Error; err != nil {
		return nil, err
	}

	user.ID = int(userEntity.ID)

	return user, nil
}

// Get the user by its ID
func (u UserImpl) GetUserByID(userID int) (*domain.User, error) {

	if userID == 0 {
		return nil, errors.New("invalid user id")
	}

	var userEntity entity.User

	if err := u.DB.First(&userEntity, userID).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}

		return nil, err
	}

	user := domain.User{
		ID:       int(userEntity.ID),
		Username: userEntity.Username,
		Email:    userEntity.Email,
		Password: userEntity.Password,
	}

	return &user, nil
}

// Delete operation
func (u UserImpl) Delete(userID int) error {
	var userEntity entity.User
	if err := u.DB.First(&userEntity, userID).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}

		return err

	}
	if err := u.DB.Delete(&userEntity).Error; err != nil {
		return err
	}
	return nil
}

func (u UserImpl) Update(userID int, update domain.UserUpdate) (*domain.User, error) {

	var userEntity entity.User

	// Check if user exists
	if err := u.DB.First(&userEntity, userID).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}

		return nil, err
	}

	updates := make(map[string]interface{})

	if update.Username != nil {
		updates["username"] = *update.Username
	}

	if update.Email != nil {
		updates["email"] = *update.Email
	}

	if len(updates) > 0 {
		if err := u.DB.Model(&userEntity).Updates(updates).Error; err != nil {
			return nil, err
		}
	}

	// Fetch updated record
	if err := u.DB.First(&userEntity, userID).Error; err != nil {
		return nil, err
	}

	return &domain.User{
		ID:       int(userEntity.ID),
		Username: userEntity.Username,
		Email:    userEntity.Email,
		Password: userEntity.Password,
	}, nil
}

func (u UserImpl) GetAllUser() ([]domain.User, error) {

	var userEntities []entity.User

	if err := u.DB.Find(&userEntities).Error; err != nil {
		return nil, err
	}

	users := make([]domain.User, 0)

	for _, userEntity := range userEntities {

		users = append(users, domain.User{
			ID:       int(userEntity.ID),
			Username: userEntity.Username,
			Email:    userEntity.Email,
			Password: userEntity.Password,
		})

	}

	return users, nil
}
