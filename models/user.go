package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID  `json:"id"`
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type UserProfile struct {
	ID        int8       `json:"id"`
	FullName  string     `json:"full_name"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type UserLoginDTO struct {
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"password"`
}

type UserRegisterDTO struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
}

type AuthResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}
