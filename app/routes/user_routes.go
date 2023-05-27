package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/smazydev/abcde/app/handlers"
	"github.com/smazydev/abcde/app/repositories"
	"gorm.io/gorm"
)

func SetupUserRoutes(app *fiber.App, db *gorm.DB, userRepo repositories.UserRepository) {
	userHandler := func(c *fiber.Ctx) error {
		return handlers.CreateUser(c, userRepo)
	}

	app.Post("/api/user", userHandler)
}
