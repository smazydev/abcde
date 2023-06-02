package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/smazydev/abcde/app/handlers"
	"github.com/smazydev/abcde/app/repositories"
	"gorm.io/gorm"
)

func SetupEmployeeRoutes(app *fiber.App, db *gorm.DB, userRepo repositories.UserRepository) {
	app.Post("api/employees", func(c *fiber.Ctx) error {
		return handlers.CreateUser(c, userRepo)
	})

	// Update a user
	app.Put("api/employees/:id", func(c *fiber.Ctx) error {
		return handlers.UpdateUser(c, userRepo)
	})

	// Get a user by ID
	app.Get("api/users/:id", func(c *fiber.Ctx) error {
		return handlers.GetUserByID(c, userRepo)
	})

	// Delete a user
	app.Delete("api/users/:id", func(c *fiber.Ctx) error {
		return handlers.DeleteUser(c, userRepo)
	})
}
