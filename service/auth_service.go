package service

import (
	"errors"

	"github.com/hardzal/go-auth-supabase/models"
	"github.com/hardzal/go-auth-supabase/repository"
	jwt "github.com/hardzal/go-auth-supabase/utils/jwt"
	password "github.com/hardzal/go-auth-supabase/utils/password"
)

type IAuthRepository interface {
	LoginUser(user models.UserLoginDTO) (string, error)
	RegisterUser(user models.UserRegisterDTO) (models.User, error)
}

type AuthService struct {
	repo *repository.AuthRepository
}

func NewAuthService(repo *repository.AuthRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) RegisterUser(user models.UserRegisterDTO) (*models.User, error) {
	// check exists email
	checkUser, _ := s.repo.GetUserByEmail(user.Email)

	if checkUser != nil {
		return nil, errors.New("email already registered")
	}

	hashedPassword, err := password.Generate(user.Password)

	if err != nil {
		return nil, err
	}

	userCreated := models.User{
		Username: user.Username,
		Email:    user.Email,
		Password: hashedPassword,
	}

	return s.repo.CreateUser(userCreated)
}

func (s *AuthService) LoginUser(user models.UserLoginDTO) (string, error) {
	userLogin, err := s.repo.GetUserByEmail(user.Email)
	if err != nil {
		return "", errors.New("user not found")
	}

	if err = password.Verify(userLogin.Password, user.Password); err != nil {
		return "", errors.New("email or password wrong")
	}

	token := jwt.Generate(&jwt.TokenPayload{
		ID:       (userLogin.ID).String(),
		Username: userLogin.Username,
	})

	return token, nil
}
