package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/smazydev/abcde/app/globals"
	"github.com/smazydev/abcde/app/handlers"
	"github.com/smazydev/abcde/app/middleware"
	"github.com/smazydev/abcde/app/repositories"
	"gorm.io/gorm"
)

func SetupBusinessRoutes(app *fiber.App, db *gorm.DB, businessRepo repositories.BusinessRepository) {
	businessHandler := func(c *fiber.Ctx) error {
		return handlers.CreateBusiness(c, businessRepo)
	}

	app.Post("/api/business", middleware.AuthMiddleware(globals.AuthService), businessHandler)
}
