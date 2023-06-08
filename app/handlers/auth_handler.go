package handlers

import (
	"errors"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/smazydev/abcde/app/models"
	"github.com/smazydev/abcde/app/services"
	"github.com/smazydev/abcde/app/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Login(container *services.Container) fiber.Handler {
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
		userService := container.GetUserService()

		fetchedUser, err := userService.GetUserByEmail(user.Email)
		log.Default().Print(fetchedUser)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Interal Server Error",
			})
		}

		err = bcrypt.CompareHashAndPassword([]byte(fetchedUser.Password), []byte(user.Password))
		if err != nil {
			// Passwords don't match
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"message": "Email or password incorrect",
			})
		}

		// Passwords match

		log.Print(fetchedUser.ID) // Generate JWT
		token, err := utils.GenerateJWT(fetchedUser.ID.String())
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to generate JWT",
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"token": token,
		})
	}
}

func Signup(container *services.Container) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var user models.User
		err := c.BodyParser(&user)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid request body",
			})
		}
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Something went wrong",
				"err":     err,
			})
		}

		user.ID = uuid.New()
		user.Password = string(hashedPassword)

		userService := container.GetUserService()
		result, err := userService.CreateUser(&user)
		log.Default().Print(err)

		log.Default().Print(errors.Is(err, gorm.ErrDuplicatedKey))
		if err != nil {
			if strings.Contains(string(err.Error()), "duplicate key value violates unique constraint") {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "Failed to create the user, email already exists",
				})
			}
		}

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to create user",
			})
		}

		return c.JSON(fiber.Map{
			"message": "User created successfully",
			"user":    result,
		})
	}
}
