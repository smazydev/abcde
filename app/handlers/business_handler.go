package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/smazydev/abcde/app/models"
	"github.com/smazydev/abcde/app/services"
)

func CreateBusiness(c *fiber.Ctx, containerService services.Container) error {
	// Parse request body
	var business models.Business
	userId := c.Locals("userID")
	uuidValue, err := uuid.Parse(userId.(string))
	log.Print(uuidValue, "UUID VALUE")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid user id",
		})
	}
	err = c.BodyParser(&business)
	businessId := uuid.New()
	business.OwnerID = uuidValue
	business.ID = businessId
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	// Create the business
	businessService := containerService.GetBusinessService()
	createdBusiness, err := businessService.Create(&business)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create the business",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Business created successfully",
		"data":    createdBusiness,
	})
}

func GetAllBusinessesForUser(c *fiber.Ctx, containerService *services.Container) error {
	// Parse request body
	userId := c.Locals("userID")
	uuidValue := userId.(string)
	// Get all businesses by owner Id or user Id
	businessService := containerService.GetBusinessService()
	businessesOwnedByUser, err := businessService.GetBusinessesByOwnerID(uuidValue)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "There are no businesses available",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Businesses fetched successfully",
		"data":    businessesOwnedByUser,
	})
}
