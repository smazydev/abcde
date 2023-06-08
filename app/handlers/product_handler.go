package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/smazydev/abcde/app/models"
	"github.com/smazydev/abcde/app/services"
)

func CreateProduct(c *fiber.Ctx, containerService *services.Container) error {
	// Parse request body
	var product models.Product
	err := c.BodyParser(&product)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}
	product.ID = uuid.New()

	productService := containerService.GetProductService()
	// Create the product
	createdproduct, err := productService.CreateProduct(&product)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create the product",
			"err":     err,
		})
	}

	return c.JSON(fiber.Map{
		"message": "product created successfully",
		"data":    createdproduct,
	})
}

func UpdateProduct(c *fiber.Ctx, containerService *services.Container) error {
	// Parse request body
	var product models.Product
	err := c.BodyParser(&product)
	usrId := c.Params("id")
	log.Print("productID", usrId)
	parsedId, err := uuid.Parse(usrId)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid productId",
		})
	}

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	// Update the product
	productService := containerService.GetProductService()
	updatedproduct, err := productService.UpdateProduct(&product, parsedId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update the product",
			"err":     err,
		})
	}

	return c.JSON(fiber.Map{
		"message": "product updated successfully",
		"data":    updatedproduct,
	})
}

func GetProductByID(c *fiber.Ctx, containerService *services.Container) error {
	// Parse product ID from path parameter
	productID := c.Params("id")

	parsedId, err := uuid.Parse(productID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Invalid product ID",
		})
	}
	// Retrieve the product by ID
	productService := containerService.GetProductService()
	product, err := productService.GetProductByID(parsedId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "product not found",
		})
	}

	return c.JSON(fiber.Map{
		"message": "product retrieved successfully",
		"data":    product,
	})
}

func DeleteProduct(c *fiber.Ctx, containerService *services.Container) error {
	// Parse product ID from path parameter
	productID := c.Params("id")

	parsedId, err := uuid.Parse(productID)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Invalid productId",
		})
	}

	// Delete the product
	productService := containerService.GetProductService()
	deletedProduct, err := productService.DeleteProduct(parsedId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete the product",
			"err":     err,
		})
	}

	return c.JSON(fiber.Map{
		"message": "product deleted successfully",
		"data":    deletedProduct,
	})
}
