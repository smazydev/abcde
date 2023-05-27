package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func ProtectedRoute() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := c.Locals("userID").(string)
		return c.JSON(fiber.Map{
			"message": "Protected route accessed",
			"userID":  userID,
		})
	}
}
