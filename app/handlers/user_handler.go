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

func UpdateUser(c *fiber.Ctx, repo repositories.UserRepository) error {
	// Parse request body
	var user models.User
	err := c.BodyParser(&user)
	usrId := c.Params("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	// Update the user
	err = repo.Update(&user, usrId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update the user",
			"err":     err,
		})
	}

	return c.JSON(fiber.Map{
		"message": "User updated successfully",
		"data":    user,
	})
}

func GetUserByID(c *fiber.Ctx, repo repositories.UserRepository) error {
	// Parse user ID from path parameter
	userID := c.Params("id")

	// Retrieve the user by ID
	user, err := repo.GetByID(string(userID))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	return c.JSON(fiber.Map{
		"message": "User retrieved successfully",
		"data":    user,
	})
}

func DeleteUser(c *fiber.Ctx, repo repositories.UserRepository) error {
	// Parse user ID from path parameter
	userID := c.Params("id")

	// Delete the user
	err := repo.Delete(string(userID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete the user",
			"err":     err,
		})
	}

	return c.JSON(fiber.Map{
		"message": "User deleted successfully",
	})
}
