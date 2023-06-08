package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/smazydev/abcde/app/models"
	"github.com/smazydev/abcde/app/services"
)

func CreateUser(c *fiber.Ctx, containerService *services.Container) error {
	// Parse request body
	var user models.User
	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}
	user.ID = uuid.New()

	userService := containerService.GetUserService()
	// Create the user
	createdUser, err := userService.CreateUser(&user)

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

func UpdateUser(c *fiber.Ctx, containerService *services.Container) error {
	// Parse request body
	var user models.User
	err := c.BodyParser(&user)
	usrId := c.Params("id")
	log.Print("USERID", usrId)
	parsedId, err := uuid.Parse(usrId)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid userId",
		})
	}

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	// Update the user
	userService := containerService.GetUserService()
	updatedUser, err := userService.UpdateUser(&user, parsedId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update the user",
			"err":     err,
		})
	}

	return c.JSON(fiber.Map{
		"message": "User updated successfully",
		"data":    updatedUser,
	})
}

func GetUserByID(c *fiber.Ctx, containerService *services.Container) error {
	// Parse user ID from path parameter
	userID := c.Params("id")

	parsedId, err := uuid.Parse(userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Invalid User ID",
		})
	}
	// Retrieve the user by ID
	userService := containerService.GetUserService()
	user, err := userService.GetUserByID(parsedId)
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

func DeleteUser(c *fiber.Ctx, containerService *services.Container) error {
	// Parse user ID from path parameter
	userID := c.Params("id")

	// Delete the user
	userService := containerService.GetUserService()
	err := userService.DeleteUser(string(userID))
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
