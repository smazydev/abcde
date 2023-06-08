package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/smazydev/abcde/app/handlers"
	"github.com/smazydev/abcde/app/services"
)

func SetupAuthRoutes(app *fiber.App, containerService *services.Container) {
	app.Post("/api/login", handlers.Login(containerService))
	app.Post("/api/signup", handlers.Signup(containerService))
}
