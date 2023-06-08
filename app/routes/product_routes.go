package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/smazydev/abcde/app/handlers"
	"github.com/smazydev/abcde/app/services"
	"gorm.io/gorm"
)

func SetupProductRoutes(app *fiber.App, db *gorm.DB, containerService *services.Container) {
	app.Post("api/products", func(c *fiber.Ctx) error {
		return handlers.CreateProduct(c, containerService)
	})

	// Update a Product
	app.Put("api/products/:id", func(c *fiber.Ctx) error {
		return handlers.UpdateProduct(c, containerService)
	})

	// Get a user by ID
	app.Get("api/products/:id", func(c *fiber.Ctx) error {
		return handlers.GetProductByID(c, containerService)
	})

	// Delete a user
	app.Delete("api/products/:id", func(c *fiber.Ctx) error {
		return handlers.DeleteProduct(c, containerService)
	})
}
