package auth

import (
	fiber "github.com/gofiber/fiber/v2"
	"github.com/hardzal/go-auth-supabase/models"
	"github.com/hardzal/go-auth-supabase/service"
)

type AuthHandler struct {
	AuthService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{AuthService: authService}
}

func (h *AuthHandler) RegisterHandler(ctx *fiber.Ctx) error {
	var userRegisterDTO models.UserRegisterDTO

	if err := ctx.BodyParser(&userRegisterDTO); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Register failed",
			"errors":  err,
		})
	}
	user, err := h.AuthService.RegisterUser(userRegisterDTO)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Register failed",
			"errors":  err,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Register success",
		"data":    user,
	})
}

func (h *AuthHandler) LoginHandler(ctx *fiber.Ctx) error {
	var userLoginDTO models.UserLoginDTO

	if err := ctx.BodyParser(&userLoginDTO); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Login failed",
			"errors":  err,
		})
	}

	token, err := h.AuthService.LoginUser(userLoginDTO)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Login failed",
			"errors":  err,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&models.AuthResponse{
		Message: "Login success",
		Token:   token,
	})
}
