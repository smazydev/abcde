package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/smazydev/abcde/app/handlers"
	"github.com/smazydev/abcde/app/services"
	"gorm.io/gorm"
)

func SetupEmployeeRoutes(app *fiber.App, db *gorm.DB, containerService *services.Container) {
	app.Post("api/employees", func(c *fiber.Ctx) error {
		return handlers.CreateUser(c, containerService)
	})

	// Update a user
	app.Put("api/employees/:id", func(c *fiber.Ctx) error {
		return handlers.UpdateUser(c, containerService)
	})

	// Get a user by ID
	app.Get("api/users/:id", func(c *fiber.Ctx) error {
		return handlers.GetUserByID(c, containerService)
	})

	// Delete a user
	app.Delete("api/users/:id", func(c *fiber.Ctx) error {
		return handlers.DeleteUser(c, containerService)
	})
}
