package main

import (
	"log"

	fiber "github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"message": "Welcome to Auth-Supabase!",
		})
	})

	log.Println("🚀 Server running at http://localhost:4000")
	app.Listen(":4000")
}
