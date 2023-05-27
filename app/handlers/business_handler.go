package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/smazydev/abcde/app/models"
	"github.com/smazydev/abcde/app/repositories"
)

func CreateBusiness(c *fiber.Ctx, repo repositories.BusinessRepository) error {
	// Parse request body
	var business models.Business
	err := c.BodyParser(&business)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	// Create the business
	err = repo.Create(&business)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create the business",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Business created successfully",
		"data":    business,
	})
}
