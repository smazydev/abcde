package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/smazydev/abcde/app/handlers"
	"github.com/smazydev/abcde/app/middleware"
	"github.com/smazydev/abcde/app/services"
	"gorm.io/gorm"
)

func SetupBusinessRoutes(app *fiber.App, db *gorm.DB, containerService *services.Container) {

	app.Post("/api/businesses", middleware.AuthMiddleware(containerService.GetAuthService()), func(c *fiber.Ctx) error {
		return handlers.CreateBusiness(c, containerService)
	})

	app.Get("/api/businesses", middleware.AuthMiddleware(containerService.GetAuthService()), func(c *fiber.Ctx) error {
		return handlers.GetAllBusinessesForUser(c, containerService)
	})

}
