package repository

import (
	"github.com/google/uuid"
	"github.com/hardzal/go-auth-supabase/models"
	"gorm.io/gorm"
)

type IAuthRepository interface {
	GetUserByEmail(email string) (*models.UserModel, error)
	GetUserById(id string) (*models.UserModel, error)
	CreateUser(user models.UserRegisterDTO) (*models.UserModel, error)
}

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) GetUserByEmail(email string) (*models.UserModel, error) {
	var user models.UserModel

	if err := r.db.Where(&models.UserModel{Email: email}).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *AuthRepository) GetUserById(id uuid.UUID) (*models.UserModel, error) {
	var user models.UserModel

	if err := r.db.Where(&models.UserModel{ID: id}).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *AuthRepository) CreateUser(user models.UserModel) (*models.UserModel, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
