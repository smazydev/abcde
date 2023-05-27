package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/smazydev/abcde/app/services"
)

func AuthMiddleware(authService services.AuthService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get the Authorization header
		authHeader := c.Get("Authorization")

		// Check if the Authorization header is present and starts with "Bearer"
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		// Extract the token
		token := authHeader[7:]

		// Validate the token
		userID, err := authService.ValidateJWT(token)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Invalid token",
			})
		}

		// Set the user ID in the request context for subsequent handlers to access
		c.Locals("userID", userID)

		// Continue to the next handler
		return c.Next()
	}
}
