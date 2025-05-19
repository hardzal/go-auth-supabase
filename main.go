package main

import (
	"log"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/hardzal/go-auth-supabase/configs"
	"github.com/hardzal/go-auth-supabase/handler/auth"
	"github.com/hardzal/go-auth-supabase/middleware"
	"github.com/hardzal/go-auth-supabase/repository"
	"github.com/hardzal/go-auth-supabase/service"
)

func main() {
	cfg, err := configs.LoadEnv()
	if err != nil {
		log.Fatalf("failed to load env: %v", err)
	}

	gormdb, err := configs.ConnectDatabase(cfg)
	if err != nil {
		log.Fatalf("failed to connect to DB: %v", err)
	}

	db, err := gormdb.DB()
	if err != nil {
		log.Fatalf("failed to initialize DB: %v", err)
	}

	defer db.Close()

	authRepo := repository.NewAuthRepository(gormdb)
	authService := service.NewAuthService(authRepo)
	AuthHandler := auth.NewAuthHandler(*authService)

	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"message": "Welcome to Auth-Supabase!",
		})
	})

	app.Post("/v1/auth/login", AuthHandler.LoginHandler)
	app.Post("/v1/auth/register", AuthHandler.RegisterHandler)

	// check auth
	app.Get("/v1/auth/check", middleware.Auth, AuthHandler.AuthCheck)

	log.Println("🚀 Server running at http://localhost:4000")
	app.Listen(":4000")
}
