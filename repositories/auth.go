package repositories

import (
	"context"

	"github.com/kalush66/ticket-booking-project-v1/models"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func (r *AuthRepository) RegisterUser(ctx context.Context, registerData *models.AuthCredentials) (*models.User, error) {
    user := &models.User{
        Email:    registerData.Email,
        Password: registerData.Password,
    }

    res := r.db.Model(&models.User{}).Create(user)
    if res.Error != nil {
        return nil, res.Error
    }

    var createdUser models.User
    if err := r.db.First(&createdUser, user.ID).Error; err != nil {
        return nil, err
    }

    return &createdUser, nil
}

func (r *AuthRepository) GetUser(ctx context.Context, query interface{}, args ...interface{}) (*models.User, error) {
	user := &models.User{}

	res := r.db.Model(&models.User{}).Where(query, args).First(user)

	if res.Error != nil {
		return nil, res.Error
	}

	return user, nil
}

func NewAuthRepository(db *gorm.DB) models.AuthRepository {
	return &AuthRepository{
		db,
	}
}