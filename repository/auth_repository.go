package repository

import (
	"github.com/Mokarama/assignment-6-spotsync-api/database"
	"github.com/Mokarama/assignment-6-spotsync-api/models"

	"gorm.io/gorm"
)

type AuthRepository struct {
	DB *gorm.DB
}

func NewAuthRepository() *AuthRepository {
	return &AuthRepository{
		DB: database.DB,
	}
}

func (r *AuthRepository) CreateUser(user *models.User) error {
	return r.DB.Create(user).Error
}

func (r *AuthRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User

	err := r.DB.Where("email = ?", email).First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}
