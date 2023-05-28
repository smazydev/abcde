package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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
	user.ID = uuid.New()

	// Create the user
	createdUser, err := repo.Create(&user)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create the user",
			"err":     err,
		})
	}

	return c.JSON(fiber.Map{
		"message": "User created successfully",
		"data":    createdUser,
	})
}
