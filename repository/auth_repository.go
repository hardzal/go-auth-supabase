package repository

import (
	"github.com/google/uuid"
	"github.com/hardzal/go-auth-supabase/models"
	"gorm.io/gorm"
)

type IAuthRepository interface {
	GetUserByEmail(email string) (*models.User, error)
	GetUserById(id string) (*models.User, error)
	CreateUser(user models.UserRegisterDTO) (*models.User, error)
}

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User

	if err := r.db.Where(&models.User{Email: email}).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *AuthRepository) GetUserById(id uuid.UUID) (*models.User, error) {
	var user models.User

	if err := r.db.Where(&models.User{ID: id}).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *AuthRepository) CreateUser(user models.User) (*models.User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
