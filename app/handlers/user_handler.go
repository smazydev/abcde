package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/smazydev/abcde/app/models"
	"github.com/smazydev/abcde/app/repositories"
)

func CreateUser(c *fiber.Ctx, repo repositories.UserRepository) error {
	// Parse request body
	var user models.User
	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	// Check if user with the same email already exists
	existingUser, err := repo.GetByEmail(user.Email)

	// If user with the same email exists, return an error
	if existingUser != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "User with the same email already exists",
		})
	}

	// Create the user
	err = repo.Create(&user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create the user",
		})
	}

	return c.JSON(fiber.Map{
		"message": "User created successfully",
		"data":    user,
	})
}
