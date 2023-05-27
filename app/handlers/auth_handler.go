package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/smazydev/abcde/app/models"
	"github.com/smazydev/abcde/app/repositories"
	"github.com/smazydev/abcde/app/utils"
)

func Login(repo repositories.UserRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Parse request body
		var user models.User
		err := c.BodyParser(&user)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid request body",
			})
		}

		// TODO: Perform user authentication, validate credentials, etc.
		// Assume a successful login for demonstration purposes

		// Generate JWT
		token, err := utils.GenerateJWT(user.ID.String())
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to generate JWT",
			})
		}

		return c.JSON(fiber.Map{
			"token": token,
		})
	}
}
