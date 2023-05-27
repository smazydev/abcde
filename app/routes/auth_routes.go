package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/smazydev/abcde/app/handlers"
	"github.com/smazydev/abcde/app/repositories"
)

func SetupAuthRoutes(app *fiber.App, authService services.AuthService, userRepo repositories.UserRepository) {
	app.Post("/api/login", handlers.Login(userRepo))
}
